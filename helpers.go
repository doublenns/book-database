package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/inancgumus/screen"
)

func clearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

func readLibraryFromFile() {
	byteContacts, err := os.ReadFile("library.json")
	if os.IsNotExist(err) {
		return
	} else if err == nil {
		err = json.Unmarshal(byteContacts, &library)
		return
	}

	// If error reading file or if unable to parse contents of the file
	fmt.Println("Error:", err)
	os.Exit(1)
}

func saveToFile() {
	file, _ := json.MarshalIndent(library, "", " ")
	err := os.WriteFile("library.json", file, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func exitProgram() {
	fmt.Println("Quitting Go Library...")
	time.Sleep(2 * time.Second)
	quit = true
	os.Exit(0)
}
