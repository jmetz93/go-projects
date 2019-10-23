package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// make sure deck is created with x number of cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	// make sure first card in deck is ace of spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be the Ace of Spades, but got %v", d[0])
	}
	// make last card is four of clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be the Four of Clubs, but got %v", d[len(d)-1])
	}
}
func TestSaveDeckAndNewDeckFromFile(t *testing.T) {
	// delete any files in working directory with name "_decktesting"
	os.Remove("_decktesting")
	// create deck
	d := newDeck()
	// save to file "_decktesting"
	d.saveToFile("_decktesting")
	// load from file
	loadedDeck := newDeckFromFile("_decktesting")
	// assert deck length
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but go %v", len(loadedDeck))
	}
	// delete any files with name "_decktesting"
	os.Remove("_decktesting")
}
