package main

import (
	"fmt"
	"time"
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
	fmt.Println(asciiTitle)
	var additionalPrompt string
	if len(library) == 0 {
		additionalPrompt = " -- This is likely the first time you're running this program"
	}
	fmt.Printf("Read %v books from file %v", len(library), additionalPrompt)
	time.Sleep(3 * time.Second)

	for !quit {
		clearScreen()
		printMenu()
		userChoice, err := validateMenuSelection()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Press any key to try again.")
			getInput()
		} else {
			fmt.Printf("Your choice was: %d) %s\n", userChoice, menuItems[userChoice].description)
			fmt.Println("Select a new menu item.")
			getInput()
		}
	}
}
