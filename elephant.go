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
		chunkSize uint
		seqLength uint
		seed      int64
		timeLimit uint

		seq        []byte
		enteredSeq []byte
		i          uint
		err        error
	)

	flag.UintVar(&chunkSize, "c", 0, "size of chunks to break sequence into")
	flag.UintVar(&seqLength, "n", 10, "length of sequence to generate")
	flag.Int64Var(&seed, "s", time.Now().UnixNano(), "seed for RNG")
	flag.UintVar(&timeLimit, "t", 0, "time limit to memorise sequence")
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

	fmt.Println("Remember the sequence:")
	// Print sequence.
	for i = 0; i < seqLength; i++ {
		if chunkSize > 0 && i%chunkSize == 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%c", seq[i])
	}

	if timeLimit == 0 {
		fmt.Print("\nHit enter to continue...")
		fmt.Scanln()
	} else {
		time.Sleep(time.Second * time.Duration(timeLimit))
	}

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
