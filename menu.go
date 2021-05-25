package main

import (
	"fmt"
	"time"
)

var menuItems = map[int]option{
	1: view,
	2: add,
	3: edit,
	4: del,
	5: search,
	6: exit,
}

func printGreeting() {
	clearScreen()
	fmt.Println(asciiTitle)
	var additionalPrompt string
	if len(library) == 0 {
		additionalPrompt = " -- This is likely the first time you're running this program"
	}
	fmt.Printf("Read %v books from file %v", len(library), additionalPrompt)
	time.Sleep(3 * time.Second)
}
func printMenu() {
	for i := 1; i < 7; i++ {
		fmt.Printf("%d) %s\n", i, menuItems[i].description)
	}
	fmt.Print("\nChoose [1-6]: ")
}
