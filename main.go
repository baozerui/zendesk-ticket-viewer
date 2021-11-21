package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"

	"github.com/howeyc/gopass"
)

var subDomain string
var email_address string
var password string

func main() {
	fmt.Println("Welcome to Zendesk Ticket Viewer!")
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
	for {
		fmt.Println("Menu")
		fmt.Println("* Press 1 to view all tickets")
		fmt.Println("* Press 2 to view individual ticket details")
		fmt.Println("* Press 3 to quit")
		var command string
		fmt.Scanln(&command)
		if command == "1" {
			client := &http.Client{}
			var r ListTicketsResponse
			url := fmt.Sprintf("https://%s.zendesk.com/api/v2/tickets.json", subDomain)
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
			resp.Body.Close()
			err = json.Unmarshal(body, &r)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("There are %v tickets\n", len(r.Tickets))
			fmt.Printf("There are %v pages\n", math.Ceil(float64(len(r.Tickets))/float64(PAGE_SIZE)))
			cur := 0
			for {
				for i := cur * PAGE_SIZE; i < (cur+1)*PAGE_SIZE && i < len(r.Tickets); i++ {
					createdTime := r.Tickets[i].CreatedAt
					createdTimeStr := createdTime.Format("2006-01-02 15:04:05")
					fmt.Printf("Ticket with subject '%v' requested by %v and assigned by %v on %v\n",
						r.Tickets[i].Subject, r.Tickets[i].RequesterID, r.Tickets[i].AssigneeID,
						createdTimeStr)
				}
				fmt.Println("* Press 1 to see next page")
				fmt.Println("* Press 2 to see previous page")
				fmt.Println("* Press 3 to return to main menu")
				var command string
				fmt.Scanln(&command)
				if command == "1" {
					cur++
				} else if command == "2" {
					cur--
				} else {
					break
				}
			}
		} else if command == "2" {
			client := &http.Client{}
			var r ShowTicketResponse
			var id string
			fmt.Print("Input the ID of ticket: ")
			fmt.Scanln(&id)
			url := fmt.Sprintf("https://%s.zendesk.com/api/v2/tickets/%s.json", subDomain, id)
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
			resp.Body.Close()
			err = json.Unmarshal(body, &r)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Subject: %s\n", r.SingleTickets.Subject)
			fmt.Printf("Description: %s\n", r.SingleTickets.Description)
			fmt.Printf("Tags: %v\n", r.SingleTickets.Tags)
			fmt.Printf("Requester ID: %v\n", r.SingleTickets.RequesterID)
			fmt.Printf("AssigneeID ID: %v\n", r.SingleTickets.AssigneeID)
			createdTimeStr := r.SingleTickets.CreatedAt.Format("2006-01-02 15:04:05")
			fmt.Printf("Created at: %s\n", createdTimeStr)
			updatedTimeStr := r.SingleTickets.UpdatedAt.Format("2006-01-02 15:04:05")
			fmt.Printf("Updated at: %s\n", updatedTimeStr)
		} else if command == "3" {
			fmt.Println("Bye Bye")
			break
		}
	}
}
