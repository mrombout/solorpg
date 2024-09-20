package dice

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/mrombout/solorpg/assert"
)

func TestRollMany(t *testing.T) {
	// Arrange
	rng := rand.New(rand.NewSource(0))
	diceSet := []NumeralDie{{Faces: 6}, {Faces: 6}}

	// Act
	result := RollMany(rng, diceSet)

	// Assert
	assert.True(t, result > 0, fmt.Sprintf("expected total result to be > 0, but it was %v", result))
	for index, die := range diceSet {
		assert.True(t, die.Result > 0, fmt.Sprintf("expected die %v result to be > 0, but it was %v", index, die.Result))
	}
}

func TestRollOne(t *testing.T) {
	// Arrange
	rng := rand.New(rand.NewSource(0))
	die := NumeralDie{Faces: 6}

	// Act
	result := RollOne(rng, &die)

	// Assert
	assert.True(t, result > 0, fmt.Sprintf("expected total result to be > 0, but it was %v", result))
	assert.True(t, die.Result > 0, fmt.Sprintf("expected die result to be > 0, but it was %v", result))
}
