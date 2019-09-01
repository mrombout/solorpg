package dice

import (
	"fmt"
	"strconv"
)

// Parse parses a dice notation string and converts it to a slice of matching dice.
//
// See https://en.wikipedia.org/wiki/Dice_notation for more in the notation.
func Parse(diceNotation string) ([]NumeralDie, error) {
	match := diceRegexp.FindStringSubmatch(diceNotation)
	result := make(map[string]string)

	if len(match) < 5 {
		return nil, fmt.Errorf("dice notation '%s' is not valid", diceNotation)
	}

	for i, name := range diceRegexp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	numDice := 1
	if val, ok := result["num"]; ok && val != "" {
		numDice, _ = strconv.Atoi(val)
	}

	numFaces := 6
	if val, ok := result["faces"]; ok && val != "" {
		numFaces, _ = strconv.Atoi(val)
	}

	dice := []NumeralDie{}
	for i := 0; i < numDice; i++ {
		dice = append(dice, NumeralDie{
			Faces: numFaces,
		})
	}

	return dice, nil
}
