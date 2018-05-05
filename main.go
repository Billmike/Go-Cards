package main

func main() {

	cards := newDeck()
	cards.saveToFile("my_file")
}

func newCard() string {
	return "Five of diamonds"
}
