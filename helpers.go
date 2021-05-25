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
	} else if err != nil {
		fmt.Println("Error:", err)
	}

	err = json.Unmarshal(byteContacts, &library)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func saveToFile() {

}

func exitProgram() {
	fmt.Println("Quitting Go Library...")
	time.Sleep(2 * time.Second)
	quit = true
	os.Exit(0)
}
