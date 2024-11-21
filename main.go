package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Get input from user
		scanner.Scan()
		input := scanner.Text()

		command, exists := getCommands()[input]
		if exists {
			// if input is a command, call the function callback for that command
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown Command")
		}
	}
}
