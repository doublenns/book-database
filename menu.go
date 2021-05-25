package main

import "fmt"

var menuItems = map[int]option{
	1: view,
	2: add,
	3: edit,
	4: del,
	5: search,
	6: exit,
}

func printMenu() {
	for i := 1; i < 7; i++ {
		fmt.Printf("%d) %s\n", i, menuItems[i].description)
	}
	fmt.Print("\nChoose [1-6]: ")
}
