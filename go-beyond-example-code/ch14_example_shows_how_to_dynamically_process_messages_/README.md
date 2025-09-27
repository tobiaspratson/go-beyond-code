# example shows how to dynamically process messages from multiple channels using nested `select` statements. The outer `select` handles shutdown, while the inner loop tries each channel in order.

**Source**: chapter14 (line 335)

## Description

This program is from the Go Beyond golang reference book.

## Usage

```bash
# Run the program
go run main.go

# Build the program
go build -o ch14_example_shows_how_to_dynamically_process_messages_ main.go

# Run the built executable
./ch14_example_shows_how_to_dynamically_process_messages_
```

## Original Context

example shows how to dynamically process messages from multiple channels using nested `select` statements. The outer `select` handles shutdown, while the inner loop tries each channel in order.
