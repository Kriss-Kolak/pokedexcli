package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {

	//Lower the input
	lowered := strings.ToLower(text)
	//Split the input by white spaces
	splited := strings.Fields(lowered)

	return splited
}
