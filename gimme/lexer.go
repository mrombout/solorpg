package gimme

import "strings"
import "unicode"
import "unicode/utf8"
import "regexp"

type tokenType int

const (
	tokenError              tokenType = 0
	tableStart                        = 1
	diceNotationStart                 = 2
	diceNotationEnd                   = 3
	text                              = 4
	variableAssignmentStart           = 5
	variableAssignmentEnd             = 6
	optionNumber                      = 7
	diceNotation                      = 8
	ident                             = 9
	templateStart                     = 10
	templateContent                   = 11
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

var emptyToken = token{
	typ:   tokenError,
	value: "",
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
			currentTokens = lexTableDefinition(line)
		case isTableOption(line):
			currentTokens = lexTableOption(line)
		case isTemplateStart(line):
			currentTokens = lexTemplateStart(scanner)
		}

		for _, currentToken := range currentTokens {
			currentToken.line = lineNo
			tokens = append(tokens, currentToken)
		}

		lineNo++
	}

	if err := scanner.Err(); err != nil {
		return tokens, err
	}

	return tokens, nil
}

func isTableDefinition(line string) bool {
	return strings.HasPrefix(line, "table:")
}

func lexTableDefinition(line string) []token {
	submatches := tableDefinitionRegex.FindSubmatch([]byte(line))

	diceNotationVal := submatches[2]
	tableName := submatches[3]
	variableName := submatches[4]

	return []token{
		token{
			typ: tableStart,
		},
		token{
			typ: diceNotationStart,
		},
		token{
			typ:   diceNotation,
			value: string(diceNotationVal),
		},
		token{
			typ: diceNotationEnd,
		},
		token{
			typ:   text,
			value: string(tableName),
		},
		token{
			typ: variableAssignmentStart,
		},
		token{
			typ:   ident,
			value: string(variableName),
		},
		token{
			typ: variableAssignmentEnd,
		},
	}
}

func isTableOption(line string) bool {
	r, _ := utf8.DecodeRuneInString(line)
	return unicode.IsDigit(r)
}

func lexTableOption(line string) []token {
	submatches := optionDefinitionRegex.FindSubmatch([]byte(line))

	optionNumberVal := submatches[1]
	optionText := submatches[2]

	return []token{
		token{
			typ:   optionNumber,
			value: string(optionNumberVal),
		},
		token{
			typ:   text,
			value: string(optionText),
		},
	}
}

func isTemplateStart(line string) bool {
	return line == ">>>"
}

func lexTemplateStart(scanner scanner) []token {
	template := ""

	for scanner.Scan() {
		line := scanner.Text()
		template += line
	}

	return []token{
		token{
			typ: templateStart,
		},
		token{
			typ:   templateContent,
			value: template,
		},
	}
}
