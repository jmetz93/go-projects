package main

func main() {
	// cards := deck{newCard(), newCard()}    // in between curly braces are the elements in the array/slice
	// cards = append(cards, "Six of Spades") // adding element to end of array
	// cards := newDeck()
	cards := newDeckFromFile("my_cards")
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards") // save deck to a file named "my_cards"
	cards.print()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// hand, remainingDeck := deal(cards, 5)
	// fmt.Println("hand")
	// hand.print()
	// fmt.Println("remaining deck")
	// remainingDeck.print()
	// fmt.Println("cards")
	// cards.print()

}

// func newCard() string { // need to tell compiler what type of data value will be returned by function
// 	return "Five of Diamonds"
// }
