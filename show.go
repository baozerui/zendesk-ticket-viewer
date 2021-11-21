package main

import (
	"fmt"
	"math"
)

func showAllTickets(r ListTicketsResponse) {
	total_page := int(math.Ceil(float64(len(r.Tickets)) / float64(PAGE_SIZE)))
	fmt.Printf("\n\nThere are total %v tickets\n", len(r.Tickets))
	fmt.Printf("There are %v pages\n", total_page)
	cur := 0
	for {
		fmt.Printf("\nThis is page %v\n", cur+1)
		for i := cur * PAGE_SIZE; i < (cur+1)*PAGE_SIZE && i < len(r.Tickets); i++ {
			createdTime := r.Tickets[i].CreatedAt
			createdTimeStr := createdTime.Format("2006-01-02 15:04:05")
			fmt.Printf("%v\nTicket with subject '%v' requested by %v and assigned by %v on %v\n",
				r.Tickets[i].ID, r.Tickets[i].Subject, r.Tickets[i].RequesterID,
				r.Tickets[i].AssigneeID, createdTimeStr)
		}
		if cur != total_page-1 {
			fmt.Println("\n\n* Press 1 to see next page")
		}
		if cur != 0 {
			fmt.Println("* Press 2 to see previous page")
		}
		fmt.Println("* Press 3 to return to main menu")
		var command string
		var isExit bool
		for {
			fmt.Scanln(&command)
			if command == OPTION_NEXT_PAGE && cur != total_page-1 {
				cur++
				break
			} else if command == OPTION_PRE_PAGE && cur != 0 {
				cur--
				break
			} else if command == OPTION_RETURN {
				isExit = true
				break
			} else {
				fmt.Println("Invalid input")
			}
		}
		if isExit {
			break
		}
	}
}

func showDetailedTicket(r ShowTicketResponse) {
	if r.SingleTickets.Subject == "" && r.SingleTickets.Description == "" {
		fmt.Println("Can not find this ticket")
		return
	}
	fmt.Printf("\n\nSubject: %s\n", r.SingleTickets.Subject)
	fmt.Printf("Description: %s\n", r.SingleTickets.Description)
	fmt.Printf("Tags: %v\n", r.SingleTickets.Tags)
	fmt.Printf("Requester ID: %v\n", r.SingleTickets.RequesterID)
	fmt.Printf("AssigneeID ID: %v\n", r.SingleTickets.AssigneeID)
	createdTimeStr := r.SingleTickets.CreatedAt.Format("2006-01-02 15:04:05")
	fmt.Printf("Created at: %s\n", createdTimeStr)
	updatedTimeStr := r.SingleTickets.UpdatedAt.Format("2006-01-02 15:04:05")
	fmt.Printf("Updated at: %s\n", updatedTimeStr)
}
