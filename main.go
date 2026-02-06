package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	pathToInputFile := flag.String("input", "", "The text file to be input")
	// rateOfRequests := flag.Int("rate", 5, "requests per second")
	pathToOutput := flag.String("output", "", "The output text")

	flag.Parse()

	if *pathToInputFile == "" || *pathToOutput == "" {
		fmt.Println("Flags --input & --output are required")
		os.Exit(2)
	}

	inputBytes, err := os.ReadFile(*pathToInputFile)

	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(2)
	}

	lines := strings.Lines(string(inputBytes))

	targetLinks := []string{}

	for line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			continue
		}

		targetLinks = append(targetLinks, trimmedLine)
		fmt.Println(trimmedLine)
	}

	fmt.Println("\nTotal target links loaded from the file:", len(targetLinks))
}
