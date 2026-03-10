package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input-file>")
		return
	}

	inputFile := os.Args[1]
	

	if inputFile == ""{
		fmt.Println("Input file cannot be empty.")
		return
	}

    content, err := os.ReadFile("standard.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	words := strings.Split(inputFile, "\\n")

	for _, word := range words {

		if word == "" {
			fmt.Println()
			continue
		}

		for row := 0; row < 8; row++ {

			for _, char := range word {

			    index := (int(char)-32)*9 + row


			    fmt.Print(lines[index])
	        }
		    fmt.Println()

		}		

	}

}