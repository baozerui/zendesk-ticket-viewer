package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var subDomain string
	var emailAddress string
	var password string
	var err error
	subDomain, emailAddress, password, err = getCredential()
	if err != nil {
		fmt.Println("Fail to get password")
		return
	}
	for {
		command := mainMenu(os.Stdin)
		if command == OPTION_LIST_TICKETS {
			var r ListTicketsResponse
			url := fmt.Sprintf(LIST_TICKETS_URL, subDomain)
			// get all ticket from Api
			body, err := makeRequest(url, emailAddress, password)
			if err != nil {
				fmt.Println("Fail to make request")
				return
			}
			err = json.Unmarshal(body, &r)
			if err != nil {
				fmt.Println("Invalid response")
				return
			}
			showAllTickets(os.Stdin, r)
		} else if command == OPTION_SHOW_TICK {
			var r ShowTicketResponse
			var id string
			fmt.Print("Input the ID of ticket: ")
			fmt.Scanln(&id)
			url := fmt.Sprintf(SHOW_TICKET_URL, subDomain, id)
			// get one ticket from Api
			body, err := makeRequest(url, emailAddress, password)
			if err != nil {
				fmt.Println("Fail to make request")
				return
			}
			err = json.Unmarshal(body, &r)
			if err != nil {
				fmt.Println("Invalid response")
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
