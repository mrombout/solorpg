package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/mrombout/solorpg/dice"
)

var results = map[int]string{
	6: "Yes, and...",
	5: "Yes...",
	4: "Yes, but...",
	3: "No, but...",
	2: "No...",
	1: "No, and...",
}

func main() {
	modifier := "+0"
	if len(os.Args) > 1 {
		modifier = os.Args[1]
	}

	if len(modifier) < 2 {
		fmt.Println("modifier not in format <+|-><number>")
		os.Exit(1)
	}

	aspect := modifier[0]
	effect, err := strconv.Atoi(modifier[1:])
	if err != nil {
		fmt.Println("modified not a valid number")
	}

	diceArr := []dice.NumeralDie{
		dice.NumeralDie{
			Faces: 6,
		},
	}
	for i := 0; i < effect; i++ {
		diceArr = append(diceArr, dice.NumeralDie{
			Faces: 6,
		})
	}

	rand.Seed(time.Now().UTC().UnixNano())

	var highestRoll dice.NumeralDie
	var lowestRoll dice.NumeralDie
	for _, dice := range diceArr {
		result := dice.Roll()
		if result > highestRoll.Result {
			highestRoll = dice
		}
		if lowestRoll.Result == 0 || result < lowestRoll.Result {
			lowestRoll = dice
		}
	}

	var result int
	if aspect == '-' {
		result = lowestRoll.Result
	} else if aspect == '+' {
		result = highestRoll.Result
	}

	fmt.Println(results[result])
}
