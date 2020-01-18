package assert

import "testing"

func EqualError(t *testing.T, actualError error, expectedErrorStr string) {
	actualErrorStr := "nil"
	if actualError != nil {
		actualErrorStr = actualError.Error()
	}

	if actualErrorStr != expectedErrorStr {
		t.Errorf("expected '%+v' to equal '%+v'", actualErrorStr, expectedErrorStr)
	}
}
