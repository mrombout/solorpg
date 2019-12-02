package gimme

import (
	"testing"

	"github.com/mrombout/solorpg/dice"
)

func TestGenerate(t *testing.T) {
	// TODO: Seeded
	template, err := templateParser{}.Parse(`{{index . "race" | A | Capitalize}} {{index . "race"}} {{index . "occupation"}}`)
	if err != nil {
		t.Fatal(err)
	}

	generator := Generator{
		tables: []table{
			table{
				diceSet: []dice.NumeralDie{
					dice.NumeralDie{Faces: 6},
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
					dice.NumeralDie{Faces: 6},
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
