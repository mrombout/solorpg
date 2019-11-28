package rollsvc

import (
	"math/rand"
	"time"

	"github.com/mrombout/solorpg/dice"
)

var emptyResult = RollResult{}

type RollResult struct {
	Result int
	Dice   []dice.NumeralDie
}

type RollService interface {
	Roll(diceNotation string) (RollResult, error)
}

type RollServiceImpl struct{}

// Roll rolls the dice given in the dice notation.
func (RollServiceImpl) Roll(diceNotation string) (RollResult, error) {
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
