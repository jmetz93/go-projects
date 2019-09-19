package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of stings
type deck []string

func (d deck) print() { // receiver function for a deck, d refers to the deck
	for i, card := range d {
		fmt.Println(i, card)
	}
}
