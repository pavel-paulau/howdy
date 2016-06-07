package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	log "gopkg.in/inconshreveable/log15.v2"
)

// ReadJSON reads JSON-encoded payload from HTTP request
func ReadJSON(req *http.Request, payload interface{}) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(payload)
	if err != nil {
		log.Error("failed to parse body", "err", err)
	}
	return err
}

func badResponse(statusCode int) error {
	msg := fmt.Sprintf("bad response: %d", statusCode)
	return errors.New(msg)
}

// SendJSON sends an HTTP request with a JSON-encoded body
func SendJSON(url string, body interface{}) ([]byte, error) {
	client := &http.Client{}

	b, err := json.Marshal(body)
	if err != nil {
		log.Error("failed marshal request", "err", err)
		return nil, err
	}
	j := bytes.NewReader(b)

	req, err := http.NewRequest("POST", url, j)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, badResponse(resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
