package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// check the deck has x number of cards
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck of len 20, but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spade, but go: %v", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected last card of Five of Clubs but got: %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// Start by delete deck testing text file if exists
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	// Retrieve file
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected len of 20 but got: %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
