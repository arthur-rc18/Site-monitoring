package main

import (
	"fmt"
	"os"
)

func main() {

	introduction()
	menuExhibition()
	command := commandReading()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Exiting program")
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
		os.Exit(-1)
	}
}

func introduction() {
	name := "DOug"
	version := 1.2
	fmt.Println("Hello, mr.", name)
	fmt.Println("Program's version", version)
}

func menuExhibition() {
	fmt.Println("1- Starting monitoring")
	fmt.Println("2- Logs exhibition")
	fmt.Println("0- Exit program")
}

func commandReading() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The chosen command was", command)

	return command
}
