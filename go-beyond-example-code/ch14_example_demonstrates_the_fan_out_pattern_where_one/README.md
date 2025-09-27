# example demonstrates the fan-out pattern where one input channel is distributed to multiple output channels. The `select` with `default` ensures that if an output channel is full, the fan-out doesn't block and skips that channel.

**Source**: chapter14 (line 1462)

## Description

This program is from the Go Beyond golang reference book.

## Usage

```bash
# Run the program
go run main.go

# Build the program
go build -o ch14_example_demonstrates_the_fan_out_pattern_where_one main.go

# Run the built executable
./ch14_example_demonstrates_the_fan_out_pattern_where_one
```

## Original Context

example demonstrates the fan-out pattern where one input channel is distributed to multiple output channels. The `select` with `default` ensures that if an output channel is full, the fan-out doesn't block and skips that channel.
