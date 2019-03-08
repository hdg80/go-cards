package main

import "fmt"

func main() {

	// var card string = "Ace of Spades"
	// card := newCard()
	// only have to use colon on initial declaration of variable
	// card = "Five of Diamonds"

	// Slice - variable length array
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// Using for loop to iterate over slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
