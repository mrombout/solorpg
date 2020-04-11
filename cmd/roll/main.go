package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/mrombout/solorpg/dice"
)

func main() {
	diceNotation := "1d6"
	if len(os.Args) > 1 {
		diceNotation = os.Args[1]
	}

	rng := rand.New(rand.NewSource(0))
	result, err := dice.Roll(rng, diceNotation)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rolls := []string{}
	fmt.Printf("%d = ", result.Result)
	for _, dice := range result.Dice {
		rolls = append(rolls, fmt.Sprintf("%d[%s]", dice.Result, dice.Type()))
	}
	fmt.Println(strings.Join(rolls, " + "))
}
