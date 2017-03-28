# README

Elephant is a simple memory training program implemented in Go 1.8.

## Building

Run `go fmt && go build`.

## Usage

Running `elephant` without flags seeds the RNG with the current time and generates a sequence of 10 digits.

### Flags

`-c` The size of chunks to break the sequence into (0 to disable chunking)

`-n` The length of the sequence to generate

`-s` The seed value for the RNG

`-t` The time limit in seconds for memorising the sequence (no limit by default)

## Cleaning

Run `go clean`.
