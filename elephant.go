package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		seqLength = 10
		seed      = time.Now().UnixNano()
		seq       = make([]int, seqLength)
		i         int
	)
	rand.Seed(seed)

	// Initialise sequence.
	for i = 0; i < seqLength; i++ {
		seq[i] = rand.Intn(10)
	}

	// Print sequence.
	for i = 0; i < seqLength; i++ {
		fmt.Print(seq[i])
	}
	fmt.Printf("\nSeed: %v", seed) // Print seed for replication.
}
