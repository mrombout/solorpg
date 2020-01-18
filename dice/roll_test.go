package dice

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/mrombout/solorpg/assert"
)

func TestRollValidNotation(t *testing.T) {
	testCases := []struct {
		diceNotation string
	}{
		{diceNotation: "1d6"},
	}

	for _, testCase := range testCases {
		t.Run(string(testCase.diceNotation), func(t *testing.T) {
			result, err := Roll(rand.New(rand.NewSource(0)), testCase.diceNotation)
			assert.Nil(t, err)
			assert.True(t, result.Result > 0, fmt.Sprintf("expected total result to be > 0, but it was %v", result))
			for index, die := range result.Dice {
				assert.True(t, die.Result > 0, fmt.Sprintf("expected die %v result to be > 0, but it was %v", index, die.Result))
			}
		})
	}
}

func TestRollInvalidNotation(t *testing.T) {
	testCases := []struct {
		diceNotation string
	}{
		{diceNotation: "coin flip"},
	}

	for _, testCase := range testCases {
		t.Run(string(testCase.diceNotation), func(t *testing.T) {
			_, err := Roll(rand.New(rand.NewSource(0)), testCase.diceNotation)
			assert.EqualError(t, err, "dice notation 'coin flip' is not valid")
		})
	}
}
