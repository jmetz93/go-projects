package main

import "fmt"

func main() {
	// cards := deck{newCard(), newCard()}    // in between curly braces are the elements in the array/slice
	// cards = append(cards, "Six of Spades") // adding element to end of array
	// cards := newDeck()
	cards := newDeckFromFile("my_cards")
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards") // save deck to a file named "my_cards"
	cards.print()
	cards.shuffleCards()
	fmt.Println("----------------------")
	cards.print()

}
