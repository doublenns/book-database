package main

import (
	"fmt"
)

var menuItems = map[int]option{
	1: view,
	2: add,
	3: edit,
	4: del,
	5: search,
	6: exit,
}

var quit bool = false

func printMenu() {
	for !quit {
		for i := 1; i < 7; i++ {
			fmt.Printf("%d) %s\n", i, menuItems[i].description)
		}
		fmt.Print("\nChoose [1-6]: ")
	}
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
