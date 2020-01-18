package dice

import "math/rand"

// RollMany rolls one or more dice and returns the total result.
func RollMany(rng *rand.Rand, dice []NumeralDie) int {
	totalResult := 0

	for _, die := range dice {
		totalResult += die.Roll(rng)
	}

	return totalResult
}

// RollOne rolls a single die and returns the result.
func RollOne(rng *rand.Rand, die NumeralDie) int {
	return die.Roll(rng)
}
