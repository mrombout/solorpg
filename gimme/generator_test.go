package gimme

import (
	"math/rand"
	"testing"

	"github.com/mrombout/solorpg/dice"
)

func TestGenerate(t *testing.T) {
	template, err := templateParser{}.Parse(`{{index . "race" | A | Capitalize}} {{index . "race"}} {{index . "occupation"}}`)
	if err != nil {
		t.Fatal(err)
	}

	generator := Generator{
		tables: []table{
			table{
				diceSet: []dice.NumeralDie{
					dice.NumeralDie{Faces: 2},
				},
				options: []option{
					option{text: "orcish"},
					option{text: "dwarven"},
				},
				text:         "The NPC is...",
				variableName: "race",
			},
			table{
				diceSet: []dice.NumeralDie{
					dice.NumeralDie{Faces: 2},
				},
				options: []option{
					option{text: "warrior"},
					option{text: "paladin"},
				},
				text:         "The NPC's occupation is...",
				variableName: "occupation",
			},
		},
		template: template,
		rng:      rand.New(rand.NewSource(0)),
	}

	actualResult, err := generator.Generate()

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	expectedResult := "An orcish warrior"
	if actualResult != expectedResult {
		t.Fatalf("expected '%s', but was '%s'", expectedResult, actualResult)
	}
}
