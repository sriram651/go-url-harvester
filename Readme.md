# go-api-harvester

A small CLI tool to process a list of URLs concurrently with a global rate limit.

The focus of this project is demonstrating Go concurrency patterns (worker pools, channels, wait groups) with a shared rate limiter.

## What it does

* Reads a list of targets from a text file (one per line)
* Distributes work across multiple worker goroutines
* Enforces a global requests-per-second rate limit
* Processes each target and logs execution time

## Usage

```bash
go run . --input targets.txt --rate 5 --output ./data
```

### Flags

* `--input`  Path to a text file containing targets (one per line)
* `--rate`   Maximum number of jobs allowed per second (global)
* `--output` Output directory (placeholder for future use)

## Concurrency model

* A fixed-size worker pool consumes jobs from a shared channel
* A single `time.Ticker` acts as a global rate limiter
* Each job waits for one tick before executing
* A `sync.WaitGroup` ensures the program exits only after all jobs finish

## Exit behavior

* Exits only after all queued jobs are processed
* Blocks cleanly using a wait group (no sleeps or busy waiting)

## Notes

* URL fetching is not implemented yet; current processing simulates work
* This project is intended as a learning exercise for Go concurrency patterns
