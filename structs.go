package main

type book struct {
	id          int
	title       string
	author      string
	description string
}

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
	function:    addBook,
}

var edit = menuOption{
	description: "Edit a book",
	function:    editBook,
}

var del = menuOption{
	description: "Delete a book",
	function:    deleteBook,
}

var search = menuOption{
	description: "Search for a book",
	function:    searchBook,
}

var exit = menuOption{
	description: "Save and exit",
	function:    exitProgram,
	// function:    saveAndExit,
}
