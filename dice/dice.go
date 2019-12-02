package dice

import "fmt"

// RollMany rolls one or more dice and returns the total result.
func RollMany(dice []NumeralDie) int {
	totalResult := 0

	fmt.Println(dice)
	for _, die := range dice {
		fmt.Println(die)
		totalResult += die.Roll()
	}

	fmt.Println(totalResult)
	return totalResult
}

// RollOne rolls a single die and returns the result.
func RollOne(die NumeralDie) int {
	return die.Roll()
}
