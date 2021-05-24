package main

type book struct {
	id          int
	title       string
	author      string
	description string
}

type option struct {
	description string
	// function    func()
}

var view = option{
	description: "View all books",
	// function:    viewBooks,
}

var add = option{
	description: "Add a book",
	// function:    addBook,
}

var edit = option{
	description: "Edit a book",
	// function:    editBook,
}

var del = option{
	description: "Delete a book",
	// function:    deleteBook,
}

var search = option{
	description: "Search for a book",
	// function:    searchBook,
}

var exit = option{
	description: "Save and exit",
	// function:    exit,
}
