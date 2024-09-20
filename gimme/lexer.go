package gimme

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type tokenType int

const (
	tokenError              tokenType = 0
	tableStart              tokenType = 1
	diceNotationStart       tokenType = 2
	diceNotationEnd         tokenType = 3
	text                    tokenType = 4
	variableAssignmentStart tokenType = 5
	variableAssignmentEnd   tokenType = 6
	optionNumber            tokenType = 7
	diceNotation            tokenType = 8
	ident                   tokenType = 9
	templateStart           tokenType = 10
	templateContent         tokenType = 11
)

type scanner interface {
	Err() error
	Scan() bool
	Text() string
}

type token struct {
	typ   tokenType
	value string
	line  int
}

var tableDefinitionRegex = regexp.MustCompile(`(table: )\((.*)\) (.*) \[(.*)\]`)
var optionDefinitionRegex = regexp.MustCompile(`(\d*)\. (.*)`)

func lex(scanner scanner) ([]token, error) {
	tokens := []token{}

	lineNo := 1
	for scanner.Scan() {
		line := scanner.Text()

		var currentTokens = []token{}
		switch {
		case isTableDefinition(line):
			currentTokens = lexTableDefinition(line, &lineNo)
		case isTableOption(line):
			currentTokens = lexTableOption(line, &lineNo)
		case isTemplateStart(line):
			currentTokens = lexTemplateStart(scanner, &lineNo)
		}

		tokens = append(tokens, currentTokens...)
	}

	if err := scanner.Err(); err != nil {
		return tokens, err
	}

	return tokens, nil
}

func isTableDefinition(line string) bool {
	return strings.HasPrefix(line, "table:")
}

func lexTableDefinition(line string, lineNo *int) []token {
	submatches := tableDefinitionRegex.FindSubmatch([]byte(line))

	diceNotationVal := submatches[2]
	tableName := submatches[3]
	variableName := submatches[4]

	currentLine := *lineNo
	(*lineNo)++

	return []token{
		{
			typ:  tableStart,
			line: currentLine,
		},
		{
			typ:  diceNotationStart,
			line: currentLine,
		},
		{
			typ:   diceNotation,
			value: string(diceNotationVal),
			line:  currentLine,
		},
		{
			typ:  diceNotationEnd,
			line: currentLine,
		},
		{
			typ:   text,
			value: string(tableName),
			line:  currentLine,
		},
		{
			typ:  variableAssignmentStart,
			line: currentLine,
		},
		{
			typ:   ident,
			value: string(variableName),
			line:  currentLine,
		},
		{
			typ:  variableAssignmentEnd,
			line: currentLine,
		},
	}
}

func isTableOption(line string) bool {
	r, _ := utf8.DecodeRuneInString(line)
	return unicode.IsDigit(r)
}

func lexTableOption(line string, lineNo *int) []token {
	submatches := optionDefinitionRegex.FindSubmatch([]byte(line))

	optionNumberVal := submatches[1]
	optionText := submatches[2]

	currentLine := *lineNo
	(*lineNo)++

	return []token{
		{
			typ:   optionNumber,
			value: string(optionNumberVal),
			line:  currentLine,
		},
		{
			typ:   text,
			value: string(optionText),
			line:  currentLine,
		},
	}
}

func isTemplateStart(line string) bool {
	return line == ">>>"
}

func lexTemplateStart(scanner scanner, lineNo *int) []token {
	template := ""

	for scanner.Scan() {
		line := scanner.Text()
		template += line
	}

	templateStartLine := *lineNo
	templateContentLine := templateStartLine + 1
	(*lineNo) += 2

	return []token{
		{
			typ:  templateStart,
			line: templateStartLine,
		},
		{
			typ:   templateContent,
			value: template,
			line:  templateContentLine,
		},
	}
}
