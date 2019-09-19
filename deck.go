package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of stings
type deck []string

// function for creating deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// use underscore to let go know we don't need to use index
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// first receiver function (d refers to instance of deck, deck ties it to deck type)
// go best practice to refer to receiver as 1 letter variable name
func (d deck) print() { // receiver function for a deck, d refers to the deck
	for _, card := range d {
		fmt.Println(card)
	}
}
