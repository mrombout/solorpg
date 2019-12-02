package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mrombout/solorpg/gimme"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: gimme <generator file>")
		os.Exit(1)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	generatorFileName := os.Args[1]

	generator, err := gimme.NewGenerator(generatorFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result, err := generator.Generate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}
