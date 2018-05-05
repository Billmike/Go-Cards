package main

import (
	"fmt"
)

type deck []string

// Function that creates a new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}

	return cards
}

// Function that prints out the deck of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Function that handles dealing of cards

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
