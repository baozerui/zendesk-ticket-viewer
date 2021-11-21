package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var subDomain string
var emailAddress string
var password string

func main() {
	var err error
	subDomain, emailAddress, password, err = getCredential()
	if err != nil {
		log.Fatalln(err)
		return
	}
	for {
		command := mainMenu()
		if command == OPTION_LIST_TICKETS {
			var r ListTicketsResponse
			url := fmt.Sprintf(LIST_TICKETS_URL, subDomain)
			body, err := makeRequest(url)
			if err != nil {
				log.Fatal(err)
				return
			}
			err = json.Unmarshal(body, &r)
			if err != nil {
				log.Fatal(err)
				return
			}
			showAllTickets(r)
		} else if command == OPTION_SHOW_TICK {
			var r ShowTicketResponse
			var id string
			fmt.Print("Input the ID of ticket: ")
			fmt.Scanln(&id)
			url := fmt.Sprintf(SHOW_TICKET_URL, subDomain, id)
			body, err := makeRequest(url)
			if err != nil {
				log.Fatal(err)
				return
			}
			err = json.Unmarshal(body, &r)
			if err != nil {
				log.Fatal(err)
				return
			}
			showDetailedTicket(r)
		} else if command == OPTION_QUIT {
			fmt.Println("Bye Bye")
			break
		} else {
			fmt.Println("Invalid input")
		}
	}
}
