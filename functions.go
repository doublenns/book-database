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
	fmt.Println()

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
				fmt.Printf("\tTitle: [%v]: ", (*bdb)[bookIndex].Title)
				newTitle := getInput()
				fmt.Printf("\tAuthor: [%v]: ", (*bdb)[bookIndex].Author)
				newAuthor := getInput()
				fmt.Printf("\tDescription: [%v]: ", (*bdb)[bookIndex].Description)
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
				fmt.Println()
			}
		} else {
			fmt.Println(err)
		}
	}
}

func searchBook() {
	var matches bookDatabase
	fmt.Println("Type in a keyword to search for")
	fmt.Println()
	fmt.Print("\tSearch: ")
	searchString := getInput()
	searchString = strings.ToLower(searchString)

	ch := make(chan book)

	for _, book := range library {
		go determineMatch(book, searchString, &matches, ch)
	}

	for i := 1; i <= len(library); i++ {
		<-ch
	}

	fmt.Println()
	if len(matches) > 0 {
		fmt.Println("The following book(s) matched your query. Enter the book ID to see more details, or <Enter> to return.")
		fmt.Println()
		for _, book := range matches {
			fmt.Printf("\t[%d] %v\n", book.Id, book.Title)
		}
		for {
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
	} else {
		fmt.Println("No books matched your search query. Returning to main menu...")
		time.Sleep(5 * time.Second)
	}
}

func determineMatch(b book, s string, m *bookDatabase, c chan book) {
	match := false
	if strings.Contains(strings.ToLower(b.Title), s) {
		match = true
	} else if strings.Contains(strings.ToLower(b.Author), s) {
		match = true
	} else if strings.Contains(strings.ToLower(b.Description), s) {
		match = true
	}
	if match {
		*m = append(*m, b)
	}
	c <- b
}

func saveAndExit() {
	saveToFile()
	fmt.Println("Library saved.")
	fmt.Println()
	exitProgram()
}
