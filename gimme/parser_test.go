package gimme

import (
	"testing"
)

func TestTokenStackPeekEmptyTokenListReturnsNil(t *testing.T) {
	tokenStack := tokenStack{
		tokens: []token{},
	}

	result := tokenStack.peek()

	if result != nil {
		t.Errorf("expected .peek() to return nil, but was %v", *result)
	}
}

func TestTokenStackPeekReturnsFirstToken(t *testing.T) {
	var expectedToken tokenType = text
	tokenStack := tokenStack{
		tokens: []token{
			{
				typ: text,
			},
			{
				typ: variableAssignmentStart,
			},
		},
	}

	result := tokenStack.peek()

	if result.typ != expectedToken {
		t.Errorf("expected .peek() to return Text, but it was %v", result.typ)
	}
}

func TestTokenStackPopEmptyTokenListReturnsNil(t *testing.T) {
	tokenStack := tokenStack{
		tokens: []token{},
	}

	result := tokenStack.pop()

	if result != nil {
		t.Errorf("expected .pop() to return nil , but was %v", *result)
	}
}

func TestTokenStackPopRemovesAndReturnsFirstToken(t *testing.T) {
	var expectedFirstToken tokenType = variableAssignmentStart
	var expectedSecondToken tokenType = ident
	tokenStack := tokenStack{
		tokens: []token{
			{
				typ: expectedFirstToken,
			},
			{
				typ: expectedSecondToken,
			},
		},
	}

	result1 := tokenStack.pop()
	result2 := tokenStack.pop()

	if result1.typ != expectedFirstToken {
		t.Errorf("expected the first .pop() to return %v, but was %v", expectedFirstToken, result1.typ)
	}
	if result2.typ != expectedSecondToken {
		t.Errorf("expected the second .pop() to return %v, but was %v", expectedSecondToken, result2.typ)
	}
	if len(tokenStack.tokens) != 0 {
		t.Errorf("expected the token stack to be empty, but wasn't")
	}
}

func TestAcceptTokenReturnsErrorWhenTokenIsUnexpected(t *testing.T) {
	tokenStack := tokenStack{
		tokens: []token{
			{
				typ: text,
			},
		},
	}

	_, err := acceptToken(&tokenStack, variableAssignmentStart)
	if err == nil {
		t.Errorf("expected acceptToken to return an error, but it didn't")
	}
}

func TestAcceptTokenReturnsFirstWhenTokenIsExpected(t *testing.T) {
	var expectedTokenType tokenType = text
	tokenStack := tokenStack{
		tokens: []token{
			{
				typ: expectedTokenType,
			},
		},
	}

	token, err := acceptToken(&tokenStack, text)
	if err != nil {
		t.Errorf("unexpected error %v returned by acceptToken(), expected nil", err)
	}

	if token.typ != expectedTokenType {
		t.Errorf("expected acceptToken() to return a token of type %v, but was %v", expectedTokenType, token.typ)
	}
}

func TestParseTemplateValid(t *testing.T) {
	parser := NewParser(nil)
	generator := Generator{}
	tokenStack := tokenStack{
		tokens: []token{
			{typ: templateStart},
			{typ: templateContent, value: `{{.SomeValue}}`},
		},
	}

	err := parser.parseTemplate(&tokenStack, &generator)

	if err != nil {
		t.Errorf("unexpected error %v returned by parseTemplate(), expected nil", err)
	}
	if generator.template == nil {
		t.Errorf("expected parseTemplate() to populate generator.template")
	}
}

func TestParseTemplateInvalid(t *testing.T) {
	testCases := map[string]struct {
		tokens []token
	}{
		"missing start": {
			tokens: []token{
				{typ: templateContent, value: `{{.SomeValue}}`},
			},
		},
		"missing content": {
			tokens: []token{
				{typ: templateStart},
			},
		},
		"invalid template": {
			tokens: []token{
				{typ: templateStart},
				{typ: templateContent, value: `{{ missing moustache!`},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			parser := NewParser(nil)
			generator := Generator{}
			tokenStack := tokenStack{
				tokens: testCase.tokens,
			}

			err := parser.parseTemplate(&tokenStack, &generator)

			if err == nil {
				t.Error("expected parseTemplate() to return an error, but was nil")
			}
			// TODO: Assert specific error
		})
	}
}

func TestParseOptionValid(t *testing.T) {
	parser := NewParser(nil)
	expectedOptionText := `Knight`
	tokenStack := tokenStack{
		tokens: []token{
			{typ: optionNumber},
			{typ: text, value: expectedOptionText},
		},
	}

	option, err := parser.parseOption(&tokenStack)

	if err != nil {
		t.Errorf("unexpected error %v returned by parseOption(), expected nil", err)
	}
	if option.text != expectedOptionText {
		t.Errorf("expected parseTemplate() to populate generator.template")
	}
}

func TestParseOptionInvalid(t *testing.T) {
	testCases := map[string]struct {
		tokens []token
	}{
		"missing number": {
			tokens: []token{
				{typ: text, value: "Lorum ipsum dolor sit amet."},
			},
		},
		"missing text": {
			tokens: []token{
				{typ: optionNumber},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			parser := NewParser(nil)
			tokenStack := tokenStack{
				tokens: testCase.tokens,
			}

			_, err := parser.parseOption(&tokenStack)

			if err == nil {
				t.Error("expected parseTemplate() to return an error, but was nil")
			}
			// TODO: Assert specific error
		})
	}
}

func TestParseOptionsValid(t *testing.T) {
	parser := NewParser(nil)
	expectedOptionText1 := `Knight`
	expectedOptionText2 := `Thief`
	tokenStack := tokenStack{
		tokens: []token{
			{typ: optionNumber},
			{typ: text, value: expectedOptionText1},
			{typ: optionNumber},
			{typ: text, value: expectedOptionText2},
		},
	}

	options, err := parser.parseOptions(&tokenStack)

	if err != nil {
		t.Errorf("unexpected error %v returned by parseOption(), expected nil", err)
	}
	if len(options) != 2 {
		t.Fatalf("expected parseOptions() to have parsed 2 options.")
	}
	if options[0].text != expectedOptionText1 {
		t.Errorf("expected parseOptions() to populate generator.template")
	}
	if options[1].text != expectedOptionText2 {
		t.Errorf("expected parseOptions() to populate generator.template")
	}
}

func TestParseOptionsInvalid(t *testing.T) {
	testCases := map[string]struct {
		tokens []token
	}{
		"wrong text token": {
			tokens: []token{
				{typ: optionNumber},
				{typ: text, value: `Knight`},
				{typ: optionNumber},
				{typ: ident, value: `Mage`},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			parser := NewParser(nil)
			tokenStack := tokenStack{
				tokens: testCase.tokens,
			}

			_, err := parser.parseOptions(&tokenStack)

			if err == nil {
				t.Error("expected parseOptions() to return an error, but was nil")
			}
			// TODO: Assert specific error
		})
	}
}

func TestParseTableValid(t *testing.T) {
	parser := NewParser(nil)
	tokenStack := tokenStack{
		tokens: []token{
			{typ: tableStart},
			{typ: diceNotationStart},
			{typ: diceNotation, value: `1d6`},
			{typ: diceNotationEnd},
			{typ: text, value: `The NPC is a...`},
			{typ: variableAssignmentStart},
			{typ: ident, value: `occupation`},
			{typ: variableAssignmentEnd},
			{typ: optionNumber},
			{typ: text, value: `Knight`},
			{typ: optionNumber},
			{typ: text, value: `Mage`},
		},
	}

	_, err := parser.parseTable(&tokenStack)

	if err != nil {
		t.Errorf("unexpected error %q returned by parseTable(), expected nil", err)
	}

	// TODO: More assertions
}

func TestParseTableInvalid(t *testing.T) {
	testCases := map[string]struct {
		tokens []token
	}{
		"missing table start": {
			tokens: []token{
				{typ: diceNotation, value: `1d6`},
				{typ: text, value: `The NPC is a...`},
				{typ: variableAssignmentStart},
				{typ: ident, value: `occupation`},
				{typ: variableAssignmentEnd},
				{typ: optionNumber},
				{typ: text, value: `Knight`},
				{typ: optionNumber},
				{typ: text, value: `Mage`},
			},
		},
		"missing dice notation": {
			tokens: []token{
				{typ: tableStart},
				{typ: text, value: `The NPC is a...`},
				{typ: variableAssignmentStart},
				{typ: ident, value: `occupation`},
				{typ: variableAssignmentEnd},
				{typ: optionNumber},
				{typ: text, value: `Knight`},
				{typ: optionNumber},
				{typ: text, value: `Mage`},
			},
		},
		"missing text": {
			tokens: []token{
				{typ: tableStart},
				{typ: diceNotation, value: `1d6`},
				{typ: variableAssignmentStart},
				{typ: ident, value: `occupation`},
				{typ: variableAssignmentEnd},
				{typ: optionNumber},
				{typ: text, value: `Knight`},
				{typ: optionNumber},
				{typ: text, value: `Mage`},
			},
		},
		"missing variable assignment start": {
			tokens: []token{
				{typ: tableStart},
				{typ: diceNotation, value: `1d6`},
				{typ: text, value: `The NPC is a...`},
				{typ: ident, value: `occupation`},
				{typ: variableAssignmentEnd},
				{typ: optionNumber},
				{typ: text, value: `Knight`},
				{typ: optionNumber},
				{typ: text, value: `Mage`},
			},
		},
		"missing variable ident": {
			tokens: []token{
				{typ: tableStart},
				{typ: diceNotation, value: `1d6`},
				{typ: text, value: `The NPC is a...`},
				{typ: variableAssignmentStart},
				{typ: variableAssignmentEnd},
				{typ: optionNumber},
				{typ: text, value: `Knight`},
				{typ: optionNumber},
				{typ: text, value: `Mage`},
			},
		},
		"missing variable assignment end": {
			tokens: []token{
				{typ: tableStart},
				{typ: diceNotation, value: `1d6`},
				{typ: text, value: `The NPC is a...`},
				{typ: variableAssignmentStart},
				{typ: ident, value: `occupation`},
				{typ: optionNumber},
				{typ: text, value: `Knight`},
				{typ: optionNumber},
				{typ: text, value: `Mage`},
			},
		},
		"missing options text": {
			tokens: []token{
				{typ: tableStart},
				{typ: diceNotation, value: `1d6`},
				{typ: text, value: `The NPC is a...`},
				{typ: variableAssignmentStart},
				{typ: ident, value: `occupation`},
				{typ: variableAssignmentEnd},
				{typ: optionNumber},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			parser := NewParser(nil)
			tokenStack := tokenStack{
				tokens: testCase.tokens,
			}

			_, err := parser.parseTable(&tokenStack)

			if err == nil {
				t.Error("expected parseTable() to return an error, but was nil")
			}
			// TODO: Assert specific error
		})
	}
}

func TestParseTablesValid(t *testing.T) {
	parser := NewParser(nil)
	generator := Generator{}
	tokenStack := tokenStack{
		tokens: []token{
			{typ: tableStart},
			{typ: diceNotationStart},
			{typ: diceNotation, value: `1d6`},
			{typ: diceNotationEnd},
			{typ: text, value: `The NPC is a...`},
			{typ: variableAssignmentStart},
			{typ: ident, value: `occupation`},
			{typ: variableAssignmentEnd},
			{typ: optionNumber},
			{typ: text, value: `Knight`},
			{typ: optionNumber},
			{typ: text, value: `Mage`},
			{typ: tableStart},
			{typ: diceNotationStart},
			{typ: diceNotation, value: `1d6`},
			{typ: diceNotationEnd},
			{typ: text, value: `The NPC is a...`},
			{typ: variableAssignmentStart},
			{typ: ident, value: `occupation`},
			{typ: variableAssignmentEnd},
			{typ: optionNumber},
			{typ: text, value: `Knight`},
			{typ: optionNumber},
			{typ: text, value: `Mage`},
		},
	}

	err := parser.parseTables(&tokenStack, &generator)

	if err != nil {
		t.Errorf("unexpected error %q returned by parseTables(), expected nil", err)
	}

	// TODO: More assertions
}

func TestParseTablesInvalid(t *testing.T) {
	testCases := map[string]struct {
		generator Generator
		tokens    []token
	}{
		"invalid second table": {
			generator: Generator{},
			tokens: []token{
				{typ: tableStart},
				{typ: diceNotation, value: `1d6`},
				{typ: text, value: `The NPC is a...`},
				{typ: variableAssignmentStart},
				{typ: ident, value: `occupation`},
				{typ: variableAssignmentEnd},
				{typ: optionNumber},
				{typ: text, value: `Knight`},
				{typ: optionNumber},
				{typ: text, value: `Mage`},
				{typ: tableStart},
				{typ: diceNotation, value: `1d6`},
				{typ: text, value: `The NPC is a...`},
				{typ: ident, value: `occupation`},
				{typ: variableAssignmentEnd},
				{typ: optionNumber},
				{typ: text, value: `Knight`},
				{typ: optionNumber},
				{typ: text, value: `Mage`},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			parser := NewParser(nil)
			tokenStack := tokenStack{
				tokens: testCase.tokens,
			}

			err := parser.parseTables(&tokenStack, &testCase.generator)

			if err == nil {
				t.Error("expected parseTables() to return an error, but was nil")
			}
			// TODO: Assert specific error
		})
	}
}

func TestParseValid(t *testing.T) {
	parser := NewParser(nil)
	tokens := []token{
		{typ: tableStart},
		{typ: diceNotationStart},
		{typ: diceNotation, value: `1d6`},
		{typ: diceNotationEnd},
		{typ: text, value: `The NPC is a...`},
		{typ: variableAssignmentStart},
		{typ: ident, value: `occupation`},
		{typ: variableAssignmentEnd},
		{typ: optionNumber},
		{typ: text, value: `Knight`},
		{typ: optionNumber},
		{typ: text, value: `Mage`},
		{typ: tableStart},
		{typ: diceNotationStart},
		{typ: diceNotation, value: `1d6`},
		{typ: diceNotationEnd},
		{typ: text, value: `The NPC is a...`},
		{typ: variableAssignmentStart},
		{typ: ident, value: `occupation`},
		{typ: variableAssignmentEnd},
		{typ: optionNumber},
		{typ: text, value: `Knight`},
		{typ: optionNumber},
		{typ: text, value: `Mage`},
		{typ: templateStart},
		{typ: templateContent, value: `TODO: Actually fill a template`},
	}

	generator, err := parser.Parse(tokens)

	if err != nil {
		t.Errorf("unexpected error %v returned by Parse(), expected nil", err)
	}

	if len(generator.tables) != 2 {
		t.Error("expected Parse() to have parsed two tables.", err)
	}

	if generator.template == nil {
		t.Error("expected Parse() to have parsed a template.", err)
	}

	// TODO: More assertions
}
