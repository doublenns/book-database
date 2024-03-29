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
	byteLibrary, err := os.ReadFile("library.json")
	if os.IsNotExist(err) {
		return // The file not yet existing isn't an issue
	} else if err == nil {
		err = json.Unmarshal(byteLibrary, &library)
		if err == nil {
			return
		}
	}

	// If error reading file or if unable to parse contents of the file
	fmt.Println("Error when trying to read from file library.json:")
	fmt.Println("Error:", err)
	os.Exit(1)
}

func saveToFile() {
	res, _ := json.MarshalIndent(library, "", " ")
	err := ioutil.WriteFile("library.json", res, 0644)
	if err != nil {
		fmt.Println("Error when trying to save to file library.json:")
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func exitProgram() {
	fmt.Println("Quitting Go Library...")
	time.Sleep(2 * time.Second)
	quit = true
}
