package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
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

	for key := range dice {
		die := &dice[key]
		die.Roll()
	}

	totalRoll := 0
	for _, dice := range dice {
		totalRoll += dice.Result
	}

	rolls := []string{}
	fmt.Printf("%d = ", totalRoll)
	for _, dice := range dice {
		rolls = append(rolls, fmt.Sprintf("%d[%s]", dice.Result, dice.Type()))
	}
	fmt.Println(strings.Join(rolls, " + "))
}
