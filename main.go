package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/howeyc/gopass"
)

var subDomain string
var email_address string
var password string

func main() {
	fmt.Printf("Input your subdomain: ")
	fmt.Scanln(&subDomain)
	fmt.Printf("Input your email address: ")
	fmt.Scanln(&email_address)
	fmt.Printf("Input the password: ")
	passwordByte, err := gopass.GetPasswdMasked()
	password = string(passwordByte)
	if err != nil {
		log.Fatalln(err)
		return
	}
	client := &http.Client{}
	var r TicketResponse
	url := fmt.Sprintf("https://%v.zendesk.com/api/v2/", subDomain)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	auth := email_address + ":" + password
	token := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", "Basic "+token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}

}
