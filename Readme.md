# go-url-harvester

A simple concurrent CLI tool written in Go to fetch URLs at a controlled rate and write responses to an output file.

This project is intentionally kept small and pragmatic. The goal is to practice and solidify Go concurrency patterns (goroutines, channels, wait groups, and rate limiting) while still producing a usable tool.

---

## What it does

* Reads a list of URLs from a text file (one per line)
* Fetches URLs concurrently using a worker pool
* Enforces a **global requests-per-second rate limit** across all workers
* Appends fetch results to a single output file in log format
* Exits cleanly after all jobs are completed

---

## Usage

```bash
go run . --input targets.txt --rate 5 --output ./outputs/data.txt
```

### Flags

* `--input`
  Path to a text file containing URLs (one per line)

* `--rate`
  Maximum number of requests allowed per second (global across all workers)

* `--output`
  Path to the output file where results will be appended

---

## Output format

Each fetch result is appended as a log entry:

```
2026-02-10T09:15:32+05:30(https://example.com):<response body>
```

Entries are written in completion order, not input order.

---

## Concurrency model

* A fixed-size worker pool processes jobs from a shared channel
* A single `time.Ticker` acts as a global rate limiter
* Each job waits for exactly one tick before executing
* A `sync.WaitGroup` ensures the program exits only after all jobs finish

This design avoids sleeps, busy waiting, and goroutine leaks.

---

## Notes

* The tool uses append-only file writes (`O_APPEND`), which are safe for concurrent workers
* HTTP requests use the default client (no retries, no custom headers)
* Error handling favors visibility over recovery
* This is a learning-focused project, not a production crawler

---

## Future improvements (intentionally not implemented)

* Per-URL output files
* Request timeouts and retries
* Configurable worker count via flags
* Structured logging

These are left out on purpose to keep the project complete and frozen.
