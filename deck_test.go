package main

import(
"testing"
"os"
)

// test functions start lowercase
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 44 {
		t.Errorf("Expected deck length of 44, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Eleven of Clubs" {
		t.Errorf("Expected last card of Eleven of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	// removed any testing file
	os.Remove("_decktesting")
	// create new deck
	deck := newDeck()
	// save deck
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 44 {
		t.Errorf("Expected 44 cards in deck but got, %v ", len(loadedDeck))
	}
	os.Remove("_decktesting")
}