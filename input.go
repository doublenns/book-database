package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func validateMenuSelection() (int, error) {
	errorMessage := fmt.Sprintf("Invalid menu item -- selection wasn't a number from 1 - %v", len(menuItems))

	rawInput := getInput()
	input, err := strconv.Atoi(rawInput)
	if err == nil {
		if menuItems[input].description != "" {
			return input, nil
		}
	}
	return 999, errors.New(errorMessage)
}

func validateBookSelection() (int, error) {
	errorMessage := fmt.Sprintf("Invalid book selection -- selection wasn't a number from 1 - %v", len(library))

	rawInput := getInput()
	if rawInput == "" {
		return -500, nil
	}
	input, err := strconv.Atoi(rawInput)
	if err == nil {
		if input >= 1 && input <= len(library) {
			return input, nil
		}
	}
	return 999, errors.New(errorMessage)
}
