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
			currentTokens = lexTableDefinition(line, &lineNo)
		case isTableOption(line):
			currentTokens = lexTableOption(line, &lineNo)
		case isTemplateStart(line):
			currentTokens = lexTemplateStart(scanner, &lineNo)
		}

		for _, currentToken := range currentTokens {
			//currentToken.line = lineNo
			tokens = append(tokens, currentToken)
		}
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
		token{
			typ:  tableStart,
			line: currentLine,
		},
		token{
			typ:  diceNotationStart,
			line: currentLine,
		},
		token{
			typ:   diceNotation,
			value: string(diceNotationVal),
			line:  currentLine,
		},
		token{
			typ:  diceNotationEnd,
			line: currentLine,
		},
		token{
			typ:   text,
			value: string(tableName),
			line:  currentLine,
		},
		token{
			typ:  variableAssignmentStart,
			line: currentLine,
		},
		token{
			typ:   ident,
			value: string(variableName),
			line:  currentLine,
		},
		token{
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
		token{
			typ:   optionNumber,
			value: string(optionNumberVal),
			line:  currentLine,
		},
		token{
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
		token{
			typ:  templateStart,
			line: templateStartLine,
		},
		token{
			typ:   templateContent,
			value: template,
			line:  templateContentLine,
		},
	}
}
