package dice

import (
	"math/rand"
)

// Roll rolls the dice given in the dice notation.
func Roll(rng *rand.Rand, diceNotation string) (RollResult, error) {
	dice, err := Parse(diceNotation)
	if err != nil {
		return emptyResult, err
	}

	for key := range dice {
		die := &dice[key]
		die.Roll(rng)
	}

	totalRoll := 0
	for _, dice := range dice {
		totalRoll += dice.Result
	}

	return RollResult{
		Result: totalRoll,
		Dice:   dice,
	}, nil
}

// RollResult contains the result of a single dice roll.
type RollResult struct {
	Result int
	Dice   []NumeralDie
}

var emptyResult = RollResult{}
