package main

import (
	"fmt"
	"log"

	"github.com/howeyc/gopass"
)

// Get credential from User
// Return subdomain, email address and password
func getCredential() (string, string, string, error) {
	var subDomain string
	var emailAddress string
	var password string
	fmt.Println("Welcome to Zendesk Ticket Viewer!")
	fmt.Printf("Input your subdomain: ")
	fmt.Scanln(&subDomain)
	fmt.Printf("Input your email address: ")
	fmt.Scanln(&emailAddress)
	fmt.Printf("Input the password: ")
	passwordByte, err := gopass.GetPasswdMasked()
	password = string(passwordByte)
	if err != nil {
		log.Fatalln(err)
		return "", "", "", nil
	}
	return subDomain, emailAddress, password, nil
}

// Prompt of main menu
// Return the command user input
func mainMenu() string {
	fmt.Println("\n\nMenu")
	fmt.Println("* Press 1 to view all tickets")
	fmt.Println("* Press 2 to view individual ticket details")
	fmt.Println("* Press 3 to quit")
	var command string
	fmt.Scanln(&command)
	return command
}
