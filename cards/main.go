package main

import "fmt"

func main() {
	card, card2 := newCard()
	fmt.Println(card2 + card)
}

func newCard() (string, string) {
	return "Five of Diamonds", "hello"
}
