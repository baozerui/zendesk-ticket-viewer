package main

import "testing"

func TestNetwork(t *testing.T) {
	var tests = []struct {
		url string
	}{
		{
			url: "https://xxx.zendesk.com/api/v2/tickets.json",
		},
		{
			url: "https://xxx.zendesk.com/api/v2/tickets/1.json",
		},
	}
	for _, test := range tests {
		_, err := makeRequest(test.url)
		if err != nil {
			t.Errorf("fail to get request: %s", test.url)
		}
	}
}
