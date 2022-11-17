package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a file path:")

	var first string

	fmt.Scanln(&first)

	_, err := os.Stat(first)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File doesn't exist")
	} else {
		fmt.Println("File exists:", first)
	}
}
