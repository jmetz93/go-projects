package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

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
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// take deck and turn to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save deck to hardrive
func (d deck) saveToFile(filename string) error {
	// to convert string to byte slice just do []byte(stringToConvert)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// read content from "my_cards" file and convert back to deck format
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) // ReadFile returns a byte slice and an error so make variables for each
	if err != nil {                      // error handling
		// Log the error and quit program
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}
	// take byte slice, convert to string and split to array
	s := strings.Split(string(bs), ",")
	// convert to deck type and return
	return deck(s)
}

func (d deck) shuffleCards() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
