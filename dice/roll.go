package dice

import (
	"math/rand"
	"time"
)

// Roll rolls the dice given in the dice notation.
func Roll(diceNotation string, seed int64) (RollResult, error) {
	dice, err := Parse(diceNotation)
	if err != nil {
		return emptyResult, err
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

type rollRequest struct {
	DiceNotation string
}

type rollResponse struct {
	Result int
	Dice   []NumeralDie
	Err    string
}
