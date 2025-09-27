# example demonstrates non-blocking operations on a buffered channel. The channel has a capacity of 2, so the first two sends succeed, but the third send fails because the channel is full.

**Source**: chapter14 (line 683)

## Description

This program is from the Go Beyond golang reference book.

## Usage

```bash
# Run the program
go run main.go

# Build the program
go build -o ch14_example_demonstrates_non_blocking_operations_on_a_ main.go

# Run the built executable
./ch14_example_demonstrates_non_blocking_operations_on_a_
```

## Original Context

example demonstrates non-blocking operations on a buffered channel. The channel has a capacity of 2, so the first two sends succeed, but the third send fails because the channel is full.
