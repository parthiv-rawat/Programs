package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 24 {
		t.Errorf("Expected Deck Length of 16, but got %v", len(d))
	}
}
