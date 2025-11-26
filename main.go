package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if !strings.HasSuffix(strings.ToLower(inputFile), ".txt") {
		fmt.Println("Error: input file must have .txt extension")
		return
	}

	if !strings.HasSuffix(strings.ToLower(outputFile), ".txt") {
		fmt.Println("Error: output file must have .txt extension")
		return
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	result := FixQuotes(lines)
	resultF := ""
	
	for i, r := range result {
	
		for _, k := range Clean(r) {
			
			resultF += k + " "
		}
		
		if i < len(result)-1 {
			resultF += "\n"
		}
	}

	err = os.WriteFile(outputFile, []byte(resultF), 0o644)
	if err != nil {
		fmt.Println("Error writing output:", err)
		return
	}

	
}
