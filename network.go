package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

// Make a get request to get data
func makeRequest(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	auth := emailAddress + ":" + password
	// get the token from email and password
	token := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", "Basic "+token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	resp.Body.Close()
	return body, nil
}
