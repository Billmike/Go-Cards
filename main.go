package main

import (
	"fmt"
)

func main() {

	cards := []string{"Ace of spades", newCard()}
	cards = append(cards, "Two of hearts")

	for index, card := range cards {
		fmt.Println(index, card)
	}

}

func newCard() string {
	return "Five of diamonds"
}
