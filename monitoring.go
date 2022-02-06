package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	monitorings = 2
	waitingTime = 2
)

func main() {

	showIntro()

	for {
		showMenu()
		prompt := chooseCommand()

		switch prompt {
		case 1:
			monitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("I do not understand this =( ")
			os.Exit(-1)
		}
	}

}

func showIntro() {
	name := "X"
	version := 1.1
	fmt.Println("Hey", name)
	fmt.Println("Versions is", version)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Quit")
}

func chooseCommand() int {
	var prompt int
	fmt.Scan(&prompt)

	return prompt
}

// monitoring will get the sites.
func monitoring() {

	// sites := []string{"https://www.alura.com.br", "https://www.google.com", "https://www.caelum.com.br"}
	sites := readingFromFile()

	fmt.Println("Monitoring starting...")
	for i := 0; i < monitorings; i++ {

		for _, site := range sites {
			fmt.Println("Testing site: ", site)
			testSite(site)
		}
		time.Sleep(waitingTime * time.Second)
		fmt.Println()
	}
	fmt.Println()
}

// testSite will receive a slice with a list of sites
// and test each one of them
func testSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Err:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Status:", resp.StatusCode)
		createLog(site, true)
	} else {
		fmt.Println("Site:", site, "Status:", resp.StatusCode)
		createLog(site, false)
	}
}

// readingFromFile will open a file and return a slice of sites
func readingFromFile() []string {

	var sites []string
	file, err := os.Open("sites.txt")
	// file, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Err:", err)
	}

	reader := bufio.NewReader(file)

	for {

		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}

	file.Close()

	return sites
}

// showLogs will show the logs
func showLogs() {

	fmt.Println("Logs...")

	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println(string(file))

}

// createLog will crate a log.txt file
func createLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Err:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	// fmt.Println(file)
	file.Close()
}
