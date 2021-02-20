package main

import "testing"

func TestNewDeck(t *testing.T) {
	// check the deck has x number of cards
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck of len 20, but got: %v", len(d))
	}
}
