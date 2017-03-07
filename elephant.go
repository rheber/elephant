package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	var (
		seqLength  uint
		seed       int64
		seq        []byte
		enteredSeq []byte
		i          uint
		err        error
	)

	flag.UintVar(&seqLength, "n", 10, "length of sequence to generate")
	flag.Int64Var(&seed, "s", time.Now().UnixNano(), "seed for RNG")
	flag.Parse()

	if seqLength == 0 {
		panic("Sequence length must be positive")
	}

	seq = make([]byte, seqLength)
	rand.Seed(seed)

	// Initialise sequence.
	for i = 0; i < seqLength; i++ {
		seq[i] = '0' + byte(rand.Intn(10))
	}

	// Print sequence.
	for i = 0; i < seqLength; i++ {
		fmt.Printf("%c", seq[i])
	}
	fmt.Print("\nHit enter to continue...")
	fmt.Scanln()

	// Clear screen.
	var cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		panic(err) // Unable to clear screen.
	}

	fmt.Println("Enter the sequence:")
	fmt.Scanln(&enteredSeq)

	if bytes.Compare(seq, enteredSeq) != 0 {
		fmt.Println("Incorrect!")
	} else {
		fmt.Println("Correct!")
	}
	fmt.Printf("Seed: %v", seed) // Print seed for replication.
}
