package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ttacon/chalk"
)

func runSelection(o menuOption) {
	clearScreen()
	fmt.Println(strings.Repeat("=", 4), o.description, strings.Repeat("=", 4))
	fmt.Println()
	o.function()
}
func (bdb *bookDatabase) printAllBooks() {
	for _, b := range *bdb {
		fmt.Printf("\t[%v] %v \n", b.Id, b.Title)
	}
}

func viewBooks() {
	library.printAllBooks()

	for {
		fmt.Println()
		fmt.Println("To view", chalk.Underline.TextStyle("details"), "enter the book ID, to return press <Enter>")
		fmt.Println()
		fmt.Print("Book ID: ")
		userSelection, err := validateBookSelection()

		if err == nil {
			if userSelection == -500 {
				return
			} else {
				bookIndex := userSelection - 1
				printBookSelection(library, bookIndex)
			}
		} else {
			fmt.Println(err)
		}
	}
}

func printBookSelection(bdb bookDatabase, i int) {
	fmt.Println()
	fmt.Printf("\tID: %d \n", library[i].Id)
	fmt.Printf("\tTitle: %v \n", library[i].Title)
	fmt.Printf("\tAuthor: %v \n", library[i].Author)
	fmt.Printf("\tDescription: %v \n", library[i].Description)
}

func (bdb *bookDatabase) addBook() {
	fmt.Println("Please enter the following information:")
	fmt.Println()
	fmt.Print("Title: ")
	title := getInput()
	fmt.Print("Author: ")
	author := getInput()
	fmt.Print("Description: ")
	description := getInput()

	newBook := book{
		Id:          (len(*bdb) + 1),
		Title:       title,
		Author:      author,
		Description: description,
	}

	(*bdb) = append((*bdb), newBook)

	saveToFile()
	fmt.Println()
	fmt.Printf("Book [%v] Saved\n", newBook.Id)
	time.Sleep(3 * time.Second)
}

func (bdb *bookDatabase) editBook() {
	(*bdb).printAllBooks()

	for {
		fmt.Println("Enter the book ID of the book you want to edit; to return press <Enter>")
		fmt.Println()
		fmt.Print("Book ID: ")
		userSelection, err := validateBookSelection()

		if err == nil {
			if userSelection == -500 {
				return
			} else {
				bookIndex := userSelection - 1
				fmt.Println()
				fmt.Println("Input the following information. To leave a field unchanged, hit <Enter>")
				fmt.Println()
				fmt.Printf("Title: [%v]: ", (*bdb)[bookIndex].Title)
				newTitle := getInput()
				fmt.Printf("Author: [%v]: ", (*bdb)[bookIndex].Author)
				newAuthor := getInput()
				fmt.Printf("Description: [%v]: ", (*bdb)[bookIndex].Description)
				newDescription := getInput()

				if newTitle != "" {
					(*bdb)[bookIndex].Title = newTitle
				}
				if newAuthor != "" {
					(*bdb)[bookIndex].Author = newAuthor
				}
				if newDescription != "" {
					(*bdb)[bookIndex].Description = newDescription
				}

				fmt.Println()
				fmt.Println("Book saved")
			}
		} else {
			fmt.Println(err)
		}
	}
}

func searchBook() {

}

func saveAndExit() {
	saveToFile()
	exitProgram()
}
