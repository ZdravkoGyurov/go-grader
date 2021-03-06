package assignment2_test

import (
	"grader/assignment2"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 3
	actual := assignment2.Add(1, 2)

	if expected != actual {
		t.Errorf("actual=%d, expected=%d", actual, expected)
	}
}
