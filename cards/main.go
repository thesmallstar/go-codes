package main

func main() {
	cards := newDeck()
	hand, rem := deal(cards, 5)
	hand.print()
	rem.print()
}
