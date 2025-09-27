# example shows how to handle multiple operations with timeouts. The first `select` will receive from `ch1` (fast operation) and timeout on `ch2` (slow operation). The second `select` waits for the remaining slow operation.

**Source**: chapter14 (line 933)

## Description

This program is from the Go Beyond golang reference book.

## Usage

```bash
# Run the program
go run main.go

# Build the program
go build -o ch14_example_shows_how_to_handle_multiple_operations_wi main.go

# Run the built executable
./ch14_example_shows_how_to_handle_multiple_operations_wi
```

## Original Context

example shows how to handle multiple operations with timeouts. The first `select` will receive from `ch1` (fast operation) and timeout on `ch2` (slow operation). The second `select` waits for the remaining slow operation.
