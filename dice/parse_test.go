package dice

import "testing"

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
