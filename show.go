package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

// Show all the tickets user have
// Page through tickets when more than 25 are returned
func showAllTickets(stdin io.Reader, r ListTicketsResponse) {
	reader := bufio.NewReader(stdin)
	total_page := int(math.Ceil(float64(len(r.Tickets)) / float64(PAGE_SIZE)))
	if len(r.Tickets) == 0 {
		fmt.Println("There is no ticket")
		return
	}
	fmt.Printf("\n\nThere are total %v tickets\n", len(r.Tickets))
	fmt.Printf("There are %v pages\n", total_page)
	cur := 0
	for {
		fmt.Printf("\nThis is page %v\n", cur+1)
		var tickets_num int
		if len(r.Tickets)-(cur*PAGE_SIZE) >= PAGE_SIZE {
			tickets_num = PAGE_SIZE
		} else {
			tickets_num = len(r.Tickets) - (cur * PAGE_SIZE)
		}
		fmt.Printf("\nThis are %v tickets in this page\n", tickets_num)
		for i := cur * PAGE_SIZE; i < (cur+1)*PAGE_SIZE && i < len(r.Tickets); i++ {
			createdTime := r.Tickets[i].CreatedAt
			createdTimeStr := createdTime.Format("2006-01-02 15:04:05")
			fmt.Printf("%v\nTicket with subject '%v' requested by %v and assigned to %v on %v\n",
				r.Tickets[i].ID, r.Tickets[i].Subject, r.Tickets[i].RequesterID,
				r.Tickets[i].AssigneeID, createdTimeStr)
		}
		fmt.Println()
		if cur != total_page-1 {
			fmt.Println("* Press 1 to see next page")
		}
		if cur != 0 {
			fmt.Println("* Press 2 to see previous page")
		}
		fmt.Println("* Press 3 to return to main menu")
		var command string
		var isExit bool
		for {
			command, _ = reader.ReadString('\n')
			command = strings.TrimSpace(command)
			if command == OPTION_NEXT_PAGE {
				if cur != total_page-1 {
					cur++
					break
				} else {
					fmt.Println("It's the last page")
				}
			} else if command == OPTION_PRE_PAGE {
				if cur != 0 {
					cur--
					break
				} else {
					fmt.Println("It's the first page")
				}
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

// Display individual ticket details
func showDetailedTicket(r ShowTicketResponse) {
	if r.SingleTickets == nil {
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
