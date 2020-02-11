package main

import "testing"

func TestGenerateRandomNumber(t *testing.T) {
	random := generateRandomNumber()
	if random < 0 {
		t.Errorf("Random was incorrect, got: %d, want: (0-10000)", random)
	}
}
