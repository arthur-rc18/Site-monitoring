package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	introduction()
	menuExhibition()
	command := commandReading()

	switch command {
	case 1:
		startingMonitoring()
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
	name := "Doug"
	version := 1.2
	fmt.Println("Hello, mr.", name)
	fmt.Println("Program's version", version)
}

func menuExhibition() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Logs exhibition")
	fmt.Println("0- Exit program")
}

func commandReading() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The chosen command was", command)

	return command
}

func startingMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://www.alura.com.br"

	http.Get(site)
}
