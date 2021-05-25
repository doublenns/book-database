package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/inancgumus/screen"
)

func clearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

func readLibraryFromFile() {
	byteLibrary, readErr := os.ReadFile("library.json")
	if os.IsNotExist(readErr) {
		return // The file not yet existing isn't an issue
	} else if readErr == nil {
		parseErr = json.Unmarshal(byteLibrary, &library)
		if parseErr == nil {
			return
		}
	}

	// If error reading file or if unable to parse contents of the file
	fmt.Println("Error:", err)
	os.Exit(1)
}

func saveToFile() {
	res, _ := json.MarshalIndent(library, "", " ")
	err := ioutil.WriteFile("library.json", res, 0644)
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
