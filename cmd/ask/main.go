package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mrombout/solorpg/ask"
)

func main() {
	modifier := "+0"
	if len(os.Args) > 1 {
		modifier = os.Args[1]
	}

	if len(modifier) < 2 {
		fmt.Println("modifier not in format <+|-><number>")
		os.Exit(1)
	}

	modifierInt, err := strconv.Atoi(modifier)
	if err != nil {
		fmt.Println("modified not a valid number")
	}

	fmt.Println(ask.Ask(modifierInt))
}
