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
				token{typ: tableStart, line: 1},
				token{typ: diceNotationStart, line: 1},
				token{typ: diceNotation, value: "1d6", line: 1},
				token{typ: diceNotationEnd, line: 1},
				token{typ: text, value: "The NPC feels...", line: 1},
				token{typ: variableAssignmentStart, line: 1},
				token{typ: ident, value: "feels", line: 1},
				token{typ: variableAssignmentEnd, line: 1},
				token{typ: optionNumber, value: "1", line: 2},
				token{typ: text, value: "Angry", line: 2},
				token{typ: optionNumber, value: "2", line: 3},
				token{typ: text, value: "Sad", line: 3},
				token{typ: optionNumber, value: "3", line: 4},
				token{typ: text, value: "Bored", line: 4},
				token{typ: tableStart, line: 5},
				token{typ: diceNotationStart, line: 5},
				token{typ: diceNotation, value: "1d6", line: 5},
				token{typ: diceNotationEnd, line: 5},
				token{typ: text, value: "The NPC smells of...", line: 5},
				token{typ: variableAssignmentStart, line: 5},
				token{typ: ident, value: "smell", line: 5},
				token{typ: variableAssignmentEnd, line: 5},
				token{typ: optionNumber, value: "1", line: 6},
				token{typ: text, value: "Flowers", line: 6},
				token{typ: optionNumber, value: "2", line: 7},
				token{typ: text, value: "Fish", line: 7},
				token{typ: optionNumber, value: "3", line: 8},
				token{typ: text, value: "Fear", line: 8},
				token{typ: templateStart, line: 9},
				token{typ: templateContent, value: "An NPC feeling {{feels}} and smells of {{smell}}", line: 10},
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
