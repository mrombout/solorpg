package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

var diceRegexp = regexp.MustCompile(`(?P<num>\d+)?(?P<type>(?:d|d%|dF))(?P<faces>\d+)?(?P<mod>[\+\*](?P<mult>\d+))?`)

// NumeralDie represents the most common die with even numeral sides.
type NumeralDie struct {
	faces int
}

// Roll a die and return the result.
func (d *NumeralDie) Roll() int {
	if d.faces == 1 {
		return 1
	}

	result := rand.Intn(d.faces - 1)
	return result + 1
}

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
			faces: numFaces,
		})
	}

	return dice, nil
}
