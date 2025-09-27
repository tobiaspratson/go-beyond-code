# example demonstrates how to process multiple channels non-blocking. The function checks each channel in a loop, processing any available messages, and only waits when all channels are empty.

**Source**: chapter14 (line 646)

## Description

This program is from the Go Beyond golang reference book.

## Usage

```bash
# Run the program
go run main.go

# Build the program
go build -o ch14_example_demonstrates_how_to_process_multiple_chann main.go

# Run the built executable
./ch14_example_demonstrates_how_to_process_multiple_chann
```

## Original Context

example demonstrates how to process multiple channels non-blocking. The function checks each channel in a loop, processing any available messages, and only waits when all channels are empty.
