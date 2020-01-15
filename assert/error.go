package assert

import "testing"

func EqualError(t *testing.T, expectedError error, actualErrorStr string) {
	expectedErrorStr := expectedError.Error()
	if expectedErrorStr != actualErrorStr {
		t.Errorf("expected '%+v' to equal '%+v'", actualErrorStr, expectedErrorStr)
	}
}
