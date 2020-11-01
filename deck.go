package main

import "fmt"

// Create type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Aces", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven"}

	// anytime you have a variable you're not using you replace
	// with an underscore (i,j)
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

//  receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal cards
// returns multiple values 
func deal(d deck, handSize int)(deck, deck) {
	return d[:handSize], d[handSize:]
}
