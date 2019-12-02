package gimme

import (
	"strings"
	"text/template"
)

type templateParser struct {
}

func indefiniteArticleFunc(word string) string {
	// TODO: Replace with proper, less naive implementation
	firstLetter := []rune(word)[0]
	if firstLetter == 'a' || firstLetter == 'o' || firstLetter == 'i' || firstLetter == 'e' {
		return "an"
	}

	return "a"
}

func capitalizeFunc(word string) string {
	return strings.ToUpper(string([]rune(word)[0])) + word[1:]
}

// Parse parses a generator template and returns it if it's valid, otherwise it returns an error.
func (templateParser) Parse(templateContent string) (*template.Template, error) {
	generatorTemplate, err := template.New("template").Funcs(template.FuncMap{
		"A":          indefiniteArticleFunc,
		"An":         indefiniteArticleFunc,
		"Capitalize": capitalizeFunc,
	}).Parse(templateContent)
	if err != nil {
		return generatorTemplate, err
	}

	// TODO: Add functions etc.

	return generatorTemplate, nil
}
