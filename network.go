package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Make a get request to get data
func makeRequest(url string, emailAddress string, password string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Fail to make request")
		return nil, err
	}
	auth := emailAddress + ":" + password
	// get the token from email and password
	token := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", "Basic "+token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Fail to do the request")
		return nil, err
	}
	// Alert user api is unavailable
	if resp.StatusCode != 200 {
		fmt.Println("Unavailable API")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fail to read response")
		return nil, err
	}
	resp.Body.Close()
	return body, nil
}
