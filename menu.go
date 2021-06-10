package main

import (
	"fmt"
	"strings"
	"time"
)

var menuItems = map[int]menuOption{
	1: view,
	2: add,
	3: edit,
	4: search,
	5: exit,
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
	fmt.Printf("Loaded %v %v into the library %v", len(library), plural, additionalPrompt)
	time.Sleep(3 * time.Second)
}
func printMenu() {
	fmt.Println(strings.Repeat("=", 4), "Book Manager", strings.Repeat("=", 4))
	fmt.Println()
	for key, _ := range menuItems {
		fmt.Printf("\t%d) %s\n", key, menuItems[key].description)
	}
	fmt.Printf("\nChoose [1-%d]: ", len(menuItems))
}
