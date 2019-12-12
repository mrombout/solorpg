package rollsvc

import (
	"math/rand"
	"time"

	"github.com/mrombout/solorpg/dice"
)

var emptyResult = RollResult{}

// RollResult contains the result of a single dice roll.
type RollResult struct {
	Result int
	Dice   []dice.NumeralDie
}

// RollService rolls one or more dice based on the given dice notation.
type RollService interface {
	Roll(diceNotation string, seed int64) (RollResult, error)
}

// RollServiceImpl is the default implementation for RollService.
type RollServiceImpl struct{}

// Roll rolls the dice given in the dice notation.
func (RollServiceImpl) Roll(diceNotation string, seed int64) (RollResult, error) {
	dice, err := dice.Parse(diceNotation)
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
