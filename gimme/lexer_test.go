package gimme

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestLex(t *testing.T) {
	testCases := []struct {
		line           string
		expectedTokens []token
	}{
		{
			`table: (1d6) The NPC feels... [feels]
1. Angry
2. Sad
3. Bored
table: (1d6) The NPC smells of... [smell]
1. Flowers
2. Fish
3. Fear
>>>
An NPC feeling {{feels}} and smells of {{smell}}`,
			[]token{
				{typ: tableStart, line: 1},
				{typ: diceNotationStart, line: 1},
				{typ: diceNotation, value: "1d6", line: 1},
				{typ: diceNotationEnd, line: 1},
				{typ: text, value: "The NPC feels...", line: 1},
				{typ: variableAssignmentStart, line: 1},
				{typ: ident, value: "feels", line: 1},
				{typ: variableAssignmentEnd, line: 1},
				{typ: optionNumber, value: "1", line: 2},
				{typ: text, value: "Angry", line: 2},
				{typ: optionNumber, value: "2", line: 3},
				{typ: text, value: "Sad", line: 3},
				{typ: optionNumber, value: "3", line: 4},
				{typ: text, value: "Bored", line: 4},
				{typ: tableStart, line: 5},
				{typ: diceNotationStart, line: 5},
				{typ: diceNotation, value: "1d6", line: 5},
				{typ: diceNotationEnd, line: 5},
				{typ: text, value: "The NPC smells of...", line: 5},
				{typ: variableAssignmentStart, line: 5},
				{typ: ident, value: "smell", line: 5},
				{typ: variableAssignmentEnd, line: 5},
				{typ: optionNumber, value: "1", line: 6},
				{typ: text, value: "Flowers", line: 6},
				{typ: optionNumber, value: "2", line: 7},
				{typ: text, value: "Fish", line: 7},
				{typ: optionNumber, value: "3", line: 8},
				{typ: text, value: "Fear", line: 8},
				{typ: templateStart, line: 9},
				{typ: templateContent, value: "An NPC feeling {{feels}} and smells of {{smell}}", line: 10},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.line, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(testCase.line))
			tokens, err := lex(scanner)

			if err != nil {
				t.Fatalf("expected error to be nil, but got %s", err)
			}
			for i := 0; i < len(testCase.expectedTokens); i++ {
				actualToken := tokens[i]
				expectedToken := testCase.expectedTokens[i]

				if !reflect.DeepEqual(expectedToken, actualToken) {
					t.Errorf("expected token %d to equal %v, but was %v", i, expectedToken, actualToken)
				}
			}
			if len(tokens) != len(testCase.expectedTokens) {
				t.Fatalf("expected to have lexed exactly %d tokens, but was %d (%T)", len(testCase.expectedTokens), len(tokens), tokens)
			}
		})
	}
}
