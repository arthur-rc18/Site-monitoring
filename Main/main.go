package main

import (
	"fmt"
	"net/http" // The net/http library is responsible for http connections
	"os"
	"time"
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
	site := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br",
		"https://www.caelum.com.br"} // Setting the URLs

	for i, site := range site {
		fmt.Println("Testing site", i, ":", site)
		sites(site)
		time.Sleep(2 * time.Second)
	}

}

func sites(site string) {

	resp, _ := http.Get(site)

	// Get issues a GET to the specified URL
	if resp.StatusCode == 200 { // StatusCode is a function from the net/http library responsible for the http statusCode
		fmt.Println("Site:", site, "loaded with success")
	} else {
		fmt.Println("Site", site, "failed load. Status Code:", resp.StatusCode)
	}
}
