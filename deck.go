package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven"}
	// anytime you have a variable you're not using you replace
	// with an underscore (i,j)
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
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
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// convert deck to string
// type conversion
func (d deck) toString() string {
	// converts deck to slice of string
	// strings package converts to a single string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option 1 log the error and return a cal to newDeck()

		// option 2 log error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// random source to get truly random generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// iterated through elements inside of slice
	for i := range d {
		// generate random number
		newPosition := r.Intn(len(d) - 1)
		// take new position and assign to i
		// takes what is at i and assign to new position
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
