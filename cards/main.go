package main

import (
	"fmt"
)

func main() {
	//card, card2 := newCard()
	//fmt.Println(card2 + card)
	array()
}

func newCard() string {
	return "Five of Diamonds"
}

func array() {
	cards := []string{newCard(), "Five of diamonds"}
	cards = append(cards, "Six of spades")

	for _, card := range cards {
		fmt.Println(card)
	}

	fmt.Println(cards)
}
