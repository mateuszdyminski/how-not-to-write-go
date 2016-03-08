package logger

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	// given
	l, err := NewLogger("/tmp/test")
	if err != nil {
					t.Errorf("Could not create logger! Err: %v", err)
	}

	// when
	l.Log("Sample entry")

	// then
	_, err = os.Open("/tmp/test")
	if err != nil {
		t.Errorf("Open should exist! Err: %v", err)
	}
}
