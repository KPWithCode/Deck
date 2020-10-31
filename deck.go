package main

import "fmt"
// Create type of deck 
// which is a slice of strings
type deck []string

//  receiver function
func (d deck) print() {
for i, card := range d {
	fmt.Println(i, card)
}
}
