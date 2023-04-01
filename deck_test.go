package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected length to be 52 but got %v ", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck, _ := newDeckFromFile("_decktesting")

	if len(loadedDeck) != len(deck) {
		t.Errorf("deck length mistmatch %v %v", len(deck), len(loadedDeck))
	}
	os.Remove("_decktesting")
}
