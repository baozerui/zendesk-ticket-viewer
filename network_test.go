package main

import "testing"

func TestNetwork(t *testing.T) {
	var tests = []struct {
		url      string
		email    string
		password string
	}{
		{
			url:      "https://subdomain.zendesk.com/api/v2/tickets.json",
			email:    "{email_address}",
			password: "{password}",
		},
		{
			url:      "https://subdomain.zendesk.com/api/v2/tickets/1.json",
			email:    "{email_address}",
			password: "{password}",
		},
	}
	for _, test := range tests {
		_, err := makeRequest(test.url, test.email, test.password)
		if err != nil {
			t.Errorf("fail to get request: %s", test.url)
		}
	}
}
