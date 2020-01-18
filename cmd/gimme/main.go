package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mrombout/solorpg/gimme"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: gimme <generator file>")
		os.Exit(1)
	}

	generatorFileName := os.Args[1]

	generator, err := gimme.NewGenerator(generatorFileName, time.Now().UTC().UnixNano())
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
