package main

import (
	"bufio"
	"net"
	"net/http"

	log "gopkg.in/inconshreveable/log15.v2"
)

type responseWriterWrapper struct {
	rw     http.ResponseWriter
	status int
}

func (r *responseWriterWrapper) Write(p []byte) (int, error) {
	return r.rw.Write(p)
}

func (r *responseWriterWrapper) Header() http.Header {
	return r.rw.Header()
}

func (r *responseWriterWrapper) WriteHeader(status int) {
	r.status = status
	r.rw.WriteHeader(status)
}

func (r *responseWriterWrapper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h := r.rw.(http.Hijacker)
	return h.Hijack()
}

func accessLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Info("received request",
			"method", req.Method,
			"path", req.URL.Path,
		)

		rww := responseWriterWrapper{rw, 200}
		handler.ServeHTTP(&rww, req)

		log.Info("processed request",
			"method", req.Method,
			"path", req.URL.Path,
			"status", rww.status,
		)
	})
}
