package helpers

import "testing"

func TestMaxOne(t *testing.T) {
	mx := Max(0, 1)
	if mx != 1 {
		t.Error("Max did not return correct max value")
	}

}

func TestMaxTwo(t *testing.T) {
	mx := Max(1, 0)
	if mx != 1 {
		t.Error("Max did not return correct max value")
	}

}
