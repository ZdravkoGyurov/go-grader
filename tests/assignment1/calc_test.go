package assignment1_test

import (
	"grader/assignment1"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 3
	actual := assignment1.Add(1, 2)

	if expected != actual {
		t.Errorf("actual=%d, expected=%d", actual, expected)
	}
}
