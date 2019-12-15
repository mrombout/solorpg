package ask

import (
	"math/rand"
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

// Ask returns a response to a yes/no question, and possibly a consequence.
func Ask(modifier int) string {
	diceArr := []dice.NumeralDie{
		dice.NumeralDie{
			Faces: 6,
		},
	}
	for i := 0; i < modifier; i++ {
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
	if modifier < 0 {
		result = lowestRoll.Result
	} else {
		result = highestRoll.Result
	}

	return results[result]
}
