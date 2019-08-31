package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mrombout/solorpg/dice"
)

func main() {
	diceNotation := "1d6"
	if len(os.Args) > 1 {
		diceNotation = os.Args[1]
	}

	dice, err := dice.Parse(diceNotation)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	rolls := []int{}
	for _, dice := range dice {
		rolls = append(rolls, dice.Roll())
	}

	totalRoll := 0
	for _, roll := range rolls {
		totalRoll += roll
	}

	fmt.Println(totalRoll)
}
