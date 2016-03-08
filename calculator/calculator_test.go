package calculator

import "testing"

func TestAdd(t *testing.T) {
	// given
	a, b := 7, 9

	// when
	res := Add(a, b)

	// then
	if res != 16 {
		t.Errorf("Expected 16! got: %d", res)
	}
}
