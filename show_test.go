package main

import (
	"bytes"
	"testing"
	"time"
)

func TestShowDetailedTicket(t *testing.T) {
	var tests = []ShowTicketResponse{
		{
			SingleTickets: &Ticket{
				RequesterID: 1,
				AssigneeID:  2,
				Subject:     "test1",
				Tags:        []string{"test"},
				Description: "just for test",
				UpdatedAt:   time.Time{},
				CreatedAt:   time.Time{},
				ID:          1,
			},
		},
		{
			SingleTickets: &Ticket{
				RequesterID: 3,
				AssigneeID:  5,
				Subject:     "test2",
				Tags:        []string{"test"},
				Description: "just for test",
				UpdatedAt:   time.Time{},
				CreatedAt:   time.Time{},
				ID:          2,
			},
		},
	}
	for _, test := range tests {
		showDetailedTicket(test)
	}
}

func TestShowAllTickets(t *testing.T) {
	var tests = []ListTicketsResponse{
		{
			Tickets: []*Ticket{
				{
					RequesterID: 1,
					AssigneeID:  2,
					Subject:     "test1",
					Tags:        []string{"test"},
					Description: "just for test",
					UpdatedAt:   time.Time{},
					CreatedAt:   time.Time{},
					ID:          1,
				},
				{
					RequesterID: 3,
					AssigneeID:  4,
					Subject:     "test2",
					Tags:        []string{"test"},
					Description: "just for test",
					UpdatedAt:   time.Time{},
					CreatedAt:   time.Time{},
					ID:          2,
				},
			},
		},
	}
	for _, test := range tests {
		var stdin bytes.Buffer
		stdin.Write([]byte("1\n"))
		stdin.Write([]byte("2\n"))
		stdin.Write([]byte("4\n"))
		stdin.Write([]byte("3\n"))
		showAllTickets(&stdin, test)
	}
}
