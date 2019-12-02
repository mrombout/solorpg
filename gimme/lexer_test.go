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
				token{typ: tableStart},
				token{typ: diceNotationStart},
				token{typ: diceNotation, value: "1d6"},
				token{typ: diceNotationEnd},
				token{typ: text, value: "The NPC feels..."},
				token{typ: variableAssignmentStart},
				token{typ: ident, value: "feels"},
				token{typ: variableAssignmentEnd},
				token{typ: optionNumber, value: "1"},
				token{typ: text, value: "Angry"},
				token{typ: optionNumber, value: "2"},
				token{typ: text, value: "Sad"},
				token{typ: optionNumber, value: "3"},
				token{typ: text, value: "Bored"},
				token{typ: tableStart},
				token{typ: diceNotationStart},
				token{typ: diceNotation, value: "1d6"},
				token{typ: diceNotationEnd},
				token{typ: text, value: "The NPC smells of..."},
				token{typ: variableAssignmentStart},
				token{typ: ident, value: "smell"},
				token{typ: variableAssignmentEnd},
				token{typ: optionNumber, value: "1"},
				token{typ: text, value: "Flowers"},
				token{typ: optionNumber, value: "2"},
				token{typ: text, value: "Fish"},
				token{typ: optionNumber, value: "3"},
				token{typ: text, value: "Fear"},
				token{typ: templateStart},
				token{typ: templateContent, value: "An NPC feeling {{feels}} and smells of {{smell}}"},
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
