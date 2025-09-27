# example shows how channel capacity affects non-blocking operations. Unbuffered channels (capacity 0) block on send until someone receives, while buffered channels can hold a limited number of messages.

**Source**: chapter14 (line 730)

## Description

This program is from the Go Beyond golang reference book.

## Usage

```bash
# Run the program
go run main.go

# Build the program
go build -o ch14_example_shows_how_channel_capacity_affects_non_blo main.go

# Run the built executable
./ch14_example_shows_how_channel_capacity_affects_non_blo
```

## Original Context

example shows how channel capacity affects non-blocking operations. Unbuffered channels (capacity 0) block on send until someone receives, while buffered channels can hold a limited number of messages.
