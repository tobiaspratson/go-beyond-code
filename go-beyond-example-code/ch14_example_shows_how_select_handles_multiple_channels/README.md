# example shows how `select` handles multiple channels in a loop. Each iteration waits for any of the three channels to be ready, processes the message, and continues to the next iteration.

**Source**: chapter14 (line 199)

## Description

This program is from the Go Beyond golang reference book.

## Usage

```bash
# Run the program
go run main.go

# Build the program
go build -o ch14_example_shows_how_select_handles_multiple_channels main.go

# Run the built executable
./ch14_example_shows_how_select_handles_multiple_channels
```

## Original Context

example shows how `select` handles multiple channels in a loop. Each iteration waits for any of the three channels to be ready, processes the message, and continues to the next iteration.
