package main

import "fmt"

func main() {

	command := "Close"

	switch command {
	case "create":
		fmt.Println("Creating.....")
	case "open":
		fmt.Println("Opening......")
	case "Close":
		fmt.Println("Closing......")

	default:
		fmt.Println("Unrecognized command")
	}

}
