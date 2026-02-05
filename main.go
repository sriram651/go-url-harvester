package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	pathToInputFile := flag.String("input", "", "The text file to be input")
	rateOfRequests := flag.Int("rate", 5, "requests per second")
	pathToOutput := flag.String("output", "", "The output text")

	flag.Parse()

	if *pathToInputFile == "" || *pathToOutput == "" {
		fmt.Println("Flags --input & --output are required")
		os.Exit(2)
	}

	fmt.Println("Below are the flag values")
	fmt.Println("input: ", *pathToInputFile)
	fmt.Println("rate: ", *rateOfRequests)
	fmt.Println("output: ", *pathToOutput)
}
