package main

import "fmt"

func main() {
	card := newCard() // shorthand variable declaration
	fmt.Println(card)
}

func newCard() string { // write string here to tell compiler that function will return a value of type string
	return "Five of Diamonds"
}
