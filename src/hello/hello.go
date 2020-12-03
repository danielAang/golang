package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitorTimes = 3
const monitorDelay = 5

func main() {
	showIntro()
	for { // For without parameters, is equivalent to while (true)
		showMenu()
		command := readUserInput()
		switch command { // Do not need to use break inside switch-case statements
		case 1:
			startMonitoring()
			break
		case 2:
			fmt.Println("Showing logs")
		case 0:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid input. Exiting program...")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	userName := "Daniel"         // Declares and defines a value
	var appVersion float32 = 1.0 // Declares and enforces a type

	// When calling function from external package, capitalize first letter from the function
	fmt.Println("Hello, Sr.", userName)
	fmt.Println("Application version:", appVersion)
}

func showMenu() {
	fmt.Println("----------------------------")
	fmt.Println("|---- SELECT AN OPTION ----|")
	fmt.Println("----------------------------")
	fmt.Println("| 1) Monitor applications  |")
	fmt.Println("| 2) Show logs             |")
	fmt.Println("| 0) Exit                  |")
	fmt.Println("----------------------------")
}

func readUserInput() int {
	var command int
	fmt.Scan(&command) //Pointer to command variable
	return command
}

func startMonitoring() {
	fmt.Println("Monitoring applications...")
	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://wwww.alura.com.br/",
		"https://www.caelum.com.br/",
	}
	for i := 0; i < monitorTimes; i++ {
		for _, site := range sites {
			response, _ := http.Get(site)
			if 200 == response.StatusCode {
				fmt.Println("Website", site, "loaded with success:", response.Status)
			} else {
				fmt.Println("Failed to load website", site, ":", response.Status)
			}
		}
		time.Sleep(monitorDelay * time.Second)
	}
}
