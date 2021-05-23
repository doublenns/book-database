package main

import (
	"fmt"
)

var menuItems = map[int]string{
	1: "View all books",
	2: "Add a book",
	3: "Edit a book",
	4: "Delete a book",
	5: "Search for a book",
	6: "Save and exit",
}

func printMenu() {
	for i := 1; i < 7; i++ {
		fmt.Printf("%d) %s\n", i, menuItems[i])
	}
	fmt.Print("\nChoose [1-6]: ")
}

// Could convert menuItems to map[string]string and sort using this func
// But instead just accessing the menuItems values by their int keys
// func sortMenuKeys(mi map[string]string) []string {
// 	keys := make([]string, 0)
// 	for k := range mi {
// 		keys = append(keys, k)
// 	}
// 	sort.Strings(keys)
// 	return keys
// }
