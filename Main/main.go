package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http" // The net/http library is responsible for http connections
	"os"
	"strings"
	"time"
)

// Constants

const monit int = 5

func main() {

	introduction()
	readFileSite()
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
	// site := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br",
	// 	"https://www.caelum.com.br"} // Setting the URLs

	site := readFileSite()

	for i := 0; i < monit; i++ {
		for i, urls := range site {
			fmt.Println("Testing site", i, ":", urls)
			sites(urls)

		}
		time.Sleep(3 * time.Second)
		fmt.Println("")
	}

}

func sites(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Occurred an error:", err)
	}

	// Get issues a GET to the specified URL
	if resp.StatusCode == 200 { // StatusCode is a function from the net/http library responsible for the http statusCode
		fmt.Println("Site:", site, "loaded with success")
	} else {
		fmt.Println("Site", site, "failed load. Status Code:", resp.StatusCode)
	}
}

func readFileSite() []string {

	var sites []string

	arch, err := os.Open("sites.txt")

	// Handling if the file is nil
	if err != nil {
		fmt.Println("Occurred an error:", err)

	}

	reader := bufio.NewReader(arch)

	for {
		// It will read the line til the end
		line, err := reader.ReadString('\n')

		// TrimSpace returns a slice of the string s, with all leading and trailing white space removed.
		// This will make the code not jump a line.
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		// When readed all the lines, will return an error of EOF
		// And here this error is being treated
		if err == io.EOF {
			break

		}
	}

	// Closing the file
	arch.Close()

	return sites
}
