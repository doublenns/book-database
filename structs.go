package main

type book struct {
	Id          int
	Title       string
	Author      string
	Description string
}

type bookDatabase []book

type menuOption struct {
	description string
	function    func()
}

var view = menuOption{
	description: "View all books",
	function:    viewBooks,
}

var add = menuOption{
	description: "Add a book",
	function:    library.addBook,
}

var edit = menuOption{
	description: "Edit a book",
	function:    library.editBook,
}

var search = menuOption{
	description: "Search for a book",
	function:    searchBook,
}

var exit = menuOption{
	description: "Save and exit",
	function:    saveAndExit,
}
