package main

import (
	"fmt"
	"time"
)

var asciiTitle string = `
__                           
/__  _    |  o |_  ._ _. ._   
\_| (_)   |_ | |_) | (_| | \/ 
                           / 
`

func main() {
	fmt.Println(asciiTitle)
	time.Sleep(2 * time.Second)
	clearScreen()
	printMenu()

}
