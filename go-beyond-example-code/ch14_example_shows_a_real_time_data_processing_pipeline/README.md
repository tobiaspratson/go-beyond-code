# example shows a real-time data processing pipeline using `select`. Data flows from generator → processor → consumer, with each component using `select` to handle data processing and shutdown signals.

**Source**: chapter14 (line 1391)

## Description

This program is from the Go Beyond golang reference book.

## Usage

```bash
# Run the program
go run main.go

# Build the program
go build -o ch14_example_shows_a_real_time_data_processing_pipeline main.go

# Run the built executable
./ch14_example_shows_a_real_time_data_processing_pipeline
```

## Original Context

example shows a real-time data processing pipeline using `select`. Data flows from generator → processor → consumer, with each component using `select` to handle data processing and shutdown signals.
