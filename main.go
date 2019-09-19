package main

func main() {
	cards := deck{newCard(), newCard()}    // in between curly braces are the elements in the array/slice
	cards = append(cards, "Six of Spades") // adding element to end of array

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	cards.print()
}

func newCard() string { // need to tell compiler what type of data value will be returned by function
	return "Five of Diamonds"
}
