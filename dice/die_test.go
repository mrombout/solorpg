package dice

import (
	"math/rand"
	"testing"
)

func TestParseValidNotation(t *testing.T) {
	testCases := []struct {
		str              string
		expectedNumDice  int
		expectedNumFaces int
	}{
		{str: "1d6", expectedNumDice: 1, expectedNumFaces: 6},
		{str: "2d20", expectedNumDice: 2, expectedNumFaces: 20},
		{str: "d20", expectedNumDice: 1, expectedNumFaces: 20},
		{str: "2d", expectedNumDice: 2, expectedNumFaces: 6},
	}

	for _, testCase := range testCases {
		t.Run(testCase.str, func(t *testing.T) {
			dice, err := Parse(testCase.str)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}

			actualNumDice := len(dice)
			if actualNumDice != testCase.expectedNumDice {
				t.Errorf("expected %d dice, but got %d", testCase.expectedNumDice, actualNumDice)
			}
		})
	}
}

func TestParseInvalidNotation(t *testing.T) {
	testCases := []struct {
		str string
	}{
		{str: ""},
	}

	for _, testCase := range testCases {
		t.Run(testCase.str, func(t *testing.T) {
			_, err := Parse(testCase.str)
			if err == nil {
				t.Errorf("expected an error")
			}
		})
	}
}

func TestRoll(t *testing.T) {
	testCases := []struct {
		seed         int64
		die          NumeralDie
		expectedRoll int
	}{
		{seed: 1, die: NumeralDie{Faces: 1}, expectedRoll: 1},
		{seed: 2, die: NumeralDie{Faces: 5}, expectedRoll: 3},
		{seed: 3, die: NumeralDie{Faces: 20}, expectedRoll: 17},
	}

	for _, testCase := range testCases {
		t.Run(string(testCase.die.Faces), func(t *testing.T) {
			rand.Seed(testCase.seed)
			actualRoll := testCase.die.Roll()
			if actualRoll != testCase.expectedRoll {
				t.Errorf("expected a roll of %d, but rolled %d", testCase.expectedRoll, actualRoll)
			}
		})
	}
}
