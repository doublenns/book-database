package main

import (
	"fmt"
)

var asciiTitle string = `
__                           
/__  _    |  o |_  ._ _. ._   
\_| (_)   |_ | |_) | (_| | \/ 
                           / 
`

var library []book
var quit bool = false

func main() {
	readLibraryFromFile()
	printGreeting()

	for !quit {
		clearScreen()
		printMenu()
		userChoice, err := validateMenuSelection()
		if err != nil {
			fmt.Println(err)
			fmt.Print("Please press enter and try again: ")
			getInput()
		} else {
			fmt.Printf("Your choice was: %d) %s\n", userChoice, menuItems[userChoice].description)
			fmt.Println("Select a new menu item.")
			getInput()
		}
	}
}
