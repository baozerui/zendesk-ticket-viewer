package main

import "time"

type ListTicketsResponse struct {
	Tickets []Ticket `json:"tickets"`
}

type ShowTicketResponse struct {
	SingleTickets Ticket `json:"ticket"`
}

type Ticket struct {
	RequesterID int       `json:"requester_id"`
	AssigneeID  int       `json:"assignee_id,omitempty"`
	Subject     string    `json:"subject"`
	Tags        []string  `json:"tags"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	ID          int       `json:"id,omitempty"`
}
