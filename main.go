package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// Determines number of concurrent go-routines
var CONCURRENT_WORKERS int = 10

func main() {
	pathToInputFile := flag.String("input", "", "The text file to be input")
	rateOfRequests := flag.Int("rate", 5, "requests per second")
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

	var urlFetchWaitGroup sync.WaitGroup

	// This creates a channel - a conveyer belt where you send stuff (jobs)
	jobs := make(chan string)

	// Calculate the interval between each tick based on the rate of requests
	requestIntervalInMs := int(time.Second) / *rateOfRequests

	ticker := time.NewTicker(time.Duration(requestIntervalInMs))

	// Loop deploys the workers
	for i := 0; i < CONCURRENT_WORKERS; i++ {
		// This is the worker function that waits for the job just beside the conveyer belt
		go func() {
			for job := range jobs {
				<-ticker.C

				fmt.Println("fetching", job, "at", time.Now().Unix())

				responseBody, fetchErr := fetchDataFromUrl(job)

				if fetchErr != nil {
					fmt.Println(fetchErr)
				}

				writeResponseToFile(responseBody)

				// Report to the owner that a job is done
				urlFetchWaitGroup.Done()
			}
		}()
	}

	// This is the owner function that puts the boxes (jobs) into the channel whenever the worker is ready to pick it up.
	for _, link := range targetLinks {
		urlFetchWaitGroup.Add(1)
		jobs <- link
	}

	// This notifies the owner that there are no more boxes to put on the belt,
	// so he shuts off the belt once the worker receives the last job
	close(jobs)

	// Wait for all the jobs to be done before exiting
	urlFetchWaitGroup.Wait()
}

func fetchDataFromUrl(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Error fetching from url: %s", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("Fetch failed with status code %d", res.StatusCode)
	}

	body, readBodyErr := io.ReadAll(res.Body)

	if readBodyErr != nil {
		return nil, fmt.Errorf("Error while trying to parse body: %s", readBodyErr)
	}

	fmt.Println("Body:", string(body))

	return body, nil
}

func writeResponseToFile(responseBody []byte) {
	// Execute file creation, write into it
}
