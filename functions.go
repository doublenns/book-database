package main

import (
	"fmt"
	"strings"
	"time"
)

func runSelection(o menuOption) {
	clearScreen()
	fmt.Println(strings.Repeat("=", 4), o.description, strings.Repeat("=", 4))
	fmt.Println()
	o.function()
}

func viewBooks() {

}

func addBook() {
	fmt.Println("Please enter the following information:")
	fmt.Println()
	fmt.Print("Title: ")
	title := getInput()
	fmt.Print("Author: ")
	author := getInput()
	fmt.Print("Description: ")
	description := getInput()

	newBook := book{
		Id:          (len(library) + 1),
		Title:       title,
		Author:      author,
		Description: description,
	}

	library = append(library, newBook)

	saveToFile()
	fmt.Printf("Book [%v] Saved\n", newBook.Id)
	time.Sleep(5 * time.Second)
}

func editBook() {

}

func deleteBook() {

}

func searchBook() {

}

func saveAndExit() {

}
