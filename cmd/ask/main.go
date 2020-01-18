package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/mrombout/solorpg/ask"
)

var output io.Writer = os.Stdout

func main() {
	modifier := "+0"
	if len(os.Args) > 1 {
		modifier = os.Args[1]
	}

	if len(modifier) < 2 {
		fmt.Fprintln(output, "modifier not in format <+|-><number>")
		os.Exit(1)
	}

	modifierInt, err := strconv.Atoi(modifier)
	if err != nil {
		fmt.Fprintln(output, "modifier not a valid number")
		os.Exit(1)
	}

	rng := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Fprintln(output, ask.Ask(rng, modifierInt))
}
