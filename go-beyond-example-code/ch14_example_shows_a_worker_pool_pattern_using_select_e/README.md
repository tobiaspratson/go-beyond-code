# example shows a worker pool pattern using `select`. Each worker listens for jobs on a shared channel and quit signals on a private channel. The `select` statement allows workers to handle both job processing and graceful shutdown.

**Source**: chapter14 (line 1168)

## Description

This program is from the Go Beyond golang reference book.

## Usage

```bash
# Run the program
go run main.go

# Build the program
go build -o ch14_example_shows_a_worker_pool_pattern_using_select_e main.go

# Run the built executable
./ch14_example_shows_a_worker_pool_pattern_using_select_e
```

## Original Context

example shows a worker pool pattern using `select`. Each worker listens for jobs on a shared channel and quit signals on a private channel. The `select` statement allows workers to handle both job processing and graceful shutdown.
