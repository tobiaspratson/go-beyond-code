# example demonstrates the basic timeout pattern using `time.After()`. The `select` statement waits for either a message from the channel or a timeout signal. Since the goroutine takes 2 seconds but the timeout is 1 second, the timeout case will execute.

**Source**: chapter14 (line 818)

## Description

This program is from the Go Beyond golang reference book.

## Usage

```bash
# Run the program
go run main.go

# Build the program
go build -o ch14_example_demonstrates_the_basic_timeout_pattern_usi main.go

# Run the built executable
./ch14_example_demonstrates_the_basic_timeout_pattern_usi
```

## Original Context

example demonstrates the basic timeout pattern using `time.After()`. The `select` statement waits for either a message from the channel or a timeout signal. Since the goroutine takes 2 seconds but the timeout is 1 second, the timeout case will execute.
