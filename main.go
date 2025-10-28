package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		userInput := s.Text()
		cleaned := cleanInput(userInput)
		var command string = ""
		if len(cleaned) == 0 {
			continue
		}
		command = cleaned[0]
		fc, ok := commandMap[command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := fc.callback()
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
