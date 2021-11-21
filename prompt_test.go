package main

import (
	"bytes"
	"testing"
)

func TestMainMenu(t *testing.T) {
	var tests = []struct {
		want  string
		input string
	}{
		{
			"1",
			"1\n",
		},
		{
			"2",
			"2\n",
		},
		{
			"3",
			"3\n",
		},
	}
	for _, test := range tests {
		var stdin bytes.Buffer
		stdin.Write([]byte(test.input))
		res := mainMenu(&stdin)
		if res != test.want {
			t.Errorf("get command = %s; want %s", res, test.want)
		}
	}
}
