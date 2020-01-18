package gimme

import (
	"fmt"

	"github.com/mrombout/solorpg/dice"
)

type tokenStack struct {
	tokens []token
}

func (t *tokenStack) peek() *token {
	if len(t.tokens) <= 0 {
		return nil
	}
	return &t.tokens[0]
}

func (t *tokenStack) pop() *token {
	if len(t.tokens) <= 0 {
		return nil
	}

	token := t.tokens[0]
	t.tokens = t.tokens[1:]

	return &token
}

// DiceNotationParser parses a dice notation string into a set of dice.
type DiceNotationParser interface {
}

// Parser parses a gimme generator definition file.
type Parser struct {
	DiceNotationParser DiceNotationParser
	templateParser     templateParser
}

// NewParser creates a new parser.
func NewParser(diceNotationParser DiceNotationParser) Parser {
	return Parser{
		DiceNotationParser: diceNotationParser,
		templateParser:     templateParser{},
	}
}

// Parse parses the given tokens into a valid generator, or returns an error when it is not able to.
//
// Use the Lexer to get a set of tokens that this parser understands from a string.
func (p *Parser) Parse(tokens []token) (Generator, error) {
	generator := Generator{}

	stack := tokenStack{
		tokens: tokens,
	}

	if err := p.parseTables(&stack, &generator); err != nil {
		return generator, err
	}
	if err := p.parseTemplate(&stack, &generator); err != nil {
		return generator, err
	}

	return generator, nil
}

func (p *Parser) parseTables(stack *tokenStack, generator *Generator) error {
	for isToken(stack, tableStart) {
		table, err := p.parseTable(stack)
		if err != nil {
			return err
		}

		generator.tables = append(generator.tables, table)
	}
	return nil
}

func (p *Parser) parseTable(stack *tokenStack) (table, error) {
	table := table{}

	_, err := acceptToken(stack, tableStart)
	if err != nil {
		return table, err
	}

	{
		diceSet, err := p.parseDiceNotation(stack)
		if err != nil {
			return table, err
		}
		table.diceSet = diceSet
	}

	{
		textToken, err := acceptToken(stack, text)
		if err != nil {
			return table, err
		}
		table.text = textToken.value
	}

	{
		_, err := acceptToken(stack, variableAssignmentStart)
		if err != nil {
			return table, err
		}

		variableNameToken, err := acceptToken(stack, ident)
		if err != nil {
			return table, err
		}
		table.variableName = variableNameToken.value

		_, err = acceptToken(stack, variableAssignmentEnd)
		if err != nil {
			return table, err
		}
	}

	{
		options, err := p.parseOptions(stack)
		if err != nil {
			return table, err
		}
		table.options = options
	}

	return table, nil
}

func (p *Parser) parseDiceNotation(stack *tokenStack) ([]dice.NumeralDie, error) {
	_, err := acceptToken(stack, diceNotationStart)
	if err != nil {
		return []dice.NumeralDie{}, err
	}

	diceNotationToken, err := acceptToken(stack, diceNotation)
	if err != nil {
		return []dice.NumeralDie{}, err
	}

	diceSet, err := dice.Parse(diceNotationToken.value)
	if err != nil {
		return []dice.NumeralDie{}, err
	}

	_, err = acceptToken(stack, diceNotationEnd)
	if err != nil {
		return []dice.NumeralDie{}, err
	}

	return diceSet, nil
}

func (p *Parser) parseOptions(stack *tokenStack) ([]option, error) {
	options := []option{}

	for isToken(stack, optionNumber) {
		option, err := p.parseOption(stack)
		if err != nil {
			return options, err
		}
		options = append(options, option)
	}
	return options, nil
}

func (p *Parser) parseOption(stack *tokenStack) (option, error) {
	option := option{}

	_, err := acceptToken(stack, optionNumber)
	if err != nil {
		return option, err
	}

	textToken, err := acceptToken(stack, text)
	if err != nil {
		return option, err
	}
	option.text = textToken.value

	return option, nil
}

func (p *Parser) parseTemplate(stack *tokenStack, generator *Generator) error {
	_, err := acceptToken(stack, templateStart)
	if err != nil {
		return err
	}

	templateContentToken, err := acceptToken(stack, templateContent)
	if err != nil {
		return err
	}

	template, err := p.templateParser.Parse(templateContentToken.value)
	if err != nil {
		return err
	}
	generator.template = template

	return nil
}

func acceptToken(stack *tokenStack, tokenType tokenType) (*token, error) {
	if !isToken(stack, tokenType) {
		return nil, fmt.Errorf("unexpected token %v, expected %v", stack.peek(), tokenType)
	}

	return stack.pop(), nil
}

func isToken(stack *tokenStack, expectedTokenType tokenType) bool {
	if len(stack.tokens) <= 0 {
		return false
	}

	token := stack.peek()
	return token.typ == expectedTokenType
}
