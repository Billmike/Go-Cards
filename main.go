package main

func main() {

	cards := newDeck()
	hands, remainingDeck := deal(cards, 5)

	hands.print()
	remainingDeck.print()

}

func newCard() string {
	return "Five of diamonds"
}
