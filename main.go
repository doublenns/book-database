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
		userSelection, err := validateMenuSelection()
		if err != nil {
			fmt.Println(err)
			fmt.Print("Please press enter and try again: ")
			getInput()
		} else {
			selection := menuItems[userSelection]
			runSelection(selection)
		}
	}
}
