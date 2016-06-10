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

func badResponse(statusCode int, body []byte) error {
	msg := fmt.Sprintf("bad response %d: %s", statusCode, body)
	return errors.New(msg)
}

// SendJSON sends an HTTP request with a JSON-encoded body
func SendJSON(url string, body interface{}) error {
	client := &http.Client{}

	b, err := json.Marshal(body)
	if err != nil {
		log.Error("failed marshal request", "err", err)
		return err
	}
	j := bytes.NewReader(b)

	req, err := http.NewRequest("POST", url, j)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return badResponse(resp.StatusCode, respBody)
	}

	return nil
}
