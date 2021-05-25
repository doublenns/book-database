package main

import "fmt"

func runSelection(o menuOption) {
	clearScreen()
	fmt.Println(o.description)
	fmt.Println()
	o.function()
}

func viewBooks() {

}

func addBook() {

}

func editBook() {

}

func deleteBook() {

}

func searchBook() {

}

func saveAndExit() {

}
