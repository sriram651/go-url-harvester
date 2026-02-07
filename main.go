package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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
	}

	fmt.Println("\nLoaded", len(targetLinks), "targets")

	// This creates a channel - a conveyer belt where you send stuff (jobs)
	jobs := make(chan string)

	// This is the worker function that waits for the job just beside the conveyer belt
	go func() {
		for job := range jobs {
			fmt.Println("fetching", job)
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	// This is the owner function that puts the boxes (jobs) into the channel whenever the worker is ready to pick it up.
	for _, link := range targetLinks {
		jobs <- link
	}

	// This notifies the owner that there are no more boxes to put on the belt,
	// so he shuts off the belt once the worker receives the last job
	close(jobs)
}
