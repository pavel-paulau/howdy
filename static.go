package main

import (
	"net/http"
	"strings"

	log "gopkg.in/inconshreveable/log15.v2"
)

const (
	baseDir = "app"
)

func staticHandler(rw http.ResponseWriter, req *http.Request) {
	data, err := Asset(baseDir + req.URL.Path)
	if err != nil {
		log.Warn("problem with static resources", "msg", err)
		return
	}

	if strings.HasSuffix(req.URL.Path, ".css") {
		rw.Header().Set("Content-Type", "text/css")
	} else if strings.HasSuffix(req.URL.Path, ".js") {
		rw.Header().Set("Content-Type", "application/javascript")
	}

	rw.Write(data)
}
