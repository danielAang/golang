package main

import (
	"fmt"
	"os"
)

func main() {
	showIntro()
	showMenu()
	command := readUserInput()
	switch command { // Do not need to use break inside switch-case statements
	case 1:
		fmt.Println("Monitoring applications...")
		break
	case 2:
		fmt.Println("Showing logs")
	case 0:
		fmt.Println("Exiting program")
		os.Exit(0)
	default:
		fmt.Println("Invalid input. Exiting program")
		os.Exit(-1)
	}
	/*
		Prints address
		fmt.Println(&command)
	*/
	/*
		reflect.TypeOf returns variable type
		fmt.Println(reflect.TypeOf(command))
	*/
	/*
		If else
		if command == 1 {
			fmt.Println("Monitoring applications...")
		} else if command == 2 {
			fmt.Println("Showing logs")
		} else if command == 0 {
			fmt.Println("Exiting program")
		} else {
			fmt.Println("Invalid input. Exiting program")
		}
	*/
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
