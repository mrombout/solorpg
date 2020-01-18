package assert

import "testing"

func Equal(t *testing.T, expected interface{}, actual interface{}) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected '%+v', but got '%+v'", expected, actual)
	}
}

func Nil(t *testing.T, actual interface{}) {
	t.Helper()

	if actual != nil {
		t.Errorf("expected 'nil', but got '%+v'", actual)
	}
}

func NotNil(t *testing.T, actual interface{}) {
	t.Helper()

	if actual == nil {
		t.Errorf("expected '%+v' to not be nil, but it was", actual)
	}
}
