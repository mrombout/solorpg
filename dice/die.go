package dice

import (
	"fmt"
	"math/rand"
	"regexp"
)

var diceRegexp = regexp.MustCompile(`(?P<num>\d+)?(?P<type>(?:d|d%|dF))(?P<faces>\d+)?(?P<mod>[\+\*](?P<mult>\d+))?`)

// NumeralDie represents the most common die with even numeral sides.
type NumeralDie struct {
	Faces  int
	Result int
}

// Roll a die and return the result.
func (d *NumeralDie) Roll() int {
	if d.Faces == 1 {
		return 1
	}

	d.Result = rand.Intn(d.Faces-1) + 1
	return d.Result
}

// Type returns the type of the die in dice notation.
func (d *NumeralDie) Type() string {
	return fmt.Sprintf("d%d", d.Faces)
}
