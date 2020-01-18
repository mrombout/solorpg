package dice

import (
	"math/rand"
	"testing"
)

func TestDiceRoll(t *testing.T) {
	testCases := []struct {
		seed         int64
		die          NumeralDie
		expectedRoll int
	}{
		{seed: 1, die: NumeralDie{Faces: 1}, expectedRoll: 1},
		{seed: 2, die: NumeralDie{Faces: 5}, expectedRoll: 2},
		{seed: 3, die: NumeralDie{Faces: 20}, expectedRoll: 9},
	}

	for _, testCase := range testCases {
		t.Run(string(testCase.die.Faces), func(t *testing.T) {
			rng := rand.New(rand.NewSource(testCase.seed))
			actualRoll := testCase.die.Roll(rng)
			if actualRoll != testCase.expectedRoll {
				t.Errorf("expected a roll of %d, but rolled %d", testCase.expectedRoll, actualRoll)
			}
		})
	}
}

func TestDiceType(t *testing.T) {
	testCases := []struct {
		die          NumeralDie
		expectedType string
	}{
		{die: NumeralDie{Faces: 2}, expectedType: "d2"},
		{die: NumeralDie{Faces: 6}, expectedType: "d6"},
	}

	for _, testCase := range testCases {
		t.Run(string(testCase.expectedType), func(t *testing.T) {
			actualType := testCase.die.Type()
			if actualType != testCase.expectedType {
				t.Errorf("expected type %v, but was %v", testCase.expectedType, actualType)
			}
		})
	}
}
