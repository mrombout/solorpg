package assert

import "testing"

func NotEmpty(t *testing.T, actual string) {
	t.Helper()

	if actual == "" {
		t.Errorf("expected string to not be empty, but it was")
	}
}
