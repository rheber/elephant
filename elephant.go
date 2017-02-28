package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		seqLength uint
		seed      int64
		seq       []int
		i         uint
	)

	flag.UintVar(&seqLength, "n", 10, "length of sequence to generate")
	flag.Int64Var(&seed, "s", time.Now().UnixNano(), "seed for RNG")
	flag.Parse()

	if seqLength == 0 {
		panic("Sequence length must be positive")
	}

	seq = make([]int, seqLength)
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
