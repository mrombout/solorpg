package assert

import "testing"

func True(t *testing.T, expressionResult bool, message string) {
	if !expressionResult {
		t.Errorf(message)
	}
}
