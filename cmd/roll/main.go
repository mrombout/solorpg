package main

import (
	"fmt"
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

	seed := time.Now().UTC().UnixNano()
	result, err := dice.Roll(diceNotation, seed)
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
