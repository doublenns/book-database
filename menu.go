package main

import (
	"fmt"
	"time"
)

var menuItems = map[int]menuOption{
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
	plural := "books"
	if len(library) == 0 {
		additionalPrompt = " -- This is likely the first time you're running this program"
	} else if len(library) == 1 {
		plural = "book"
	}
	fmt.Printf("Loaded %v %v from file %v", len(library), plural, additionalPrompt)
	time.Sleep(3 * time.Second)
}
func printMenu() {
	for i := 1; i < 7; i++ {
		fmt.Printf("%d) %s\n", i, menuItems[i].description)
	}
	fmt.Print("\nChoose [1-6]: ")
}
