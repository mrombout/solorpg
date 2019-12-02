package gimme

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/mrombout/solorpg/dice"
)

// Generator generates a semi-random string based on it's configuration.
type Generator struct {
	tables   []table
	template *template.Template
}

type table struct {
	diceSet      []dice.NumeralDie
	options      []option
	text         string
	variableName string
}

type option struct {
	text string
}

// NewGenerator creates a new generator based on the given file.
func NewGenerator(generatorFileName string) (Generator, error) {
	content, err := ioutil.ReadFile(generatorFileName)
	if err != nil {
		return Generator{}, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	tokens, err := lex(scanner)
	if err != nil {
		return Generator{}, err
	}

	parser := NewParser(nil)
	generator, err := parser.Parse(tokens)
	if err != nil {
		return Generator{}, err
	}

	return generator, nil
}

// Generate generates a semi-random string based on the generators configuration.
//
// Calling .Generate() multiple times will result in different results.
func (g *Generator) Generate() (string, error) {
	tableResolutions := map[string]string{}
	for _, table := range g.tables {
		val, err := g.resolveTable(table)
		if err != nil {
			return "", err
		}

		tableResolutions[table.variableName] = val
	}

	var generatedContent strings.Builder
	err := g.template.Execute(&generatedContent, tableResolutions)
	if err != nil {
		return generatedContent.String(), err
	}

	return generatedContent.String(), nil
}

func (g *Generator) resolveTable(table table) (string, error) {
	rollResult := dice.RollMany(table.diceSet)
	numOptions := len(table.options)

	if rollResult > numOptions {
		return "", fmt.Errorf("rolled %d, but there are only %d options", rollResult, len(table.options))
	}

	return table.options[rollResult].text, nil
}
