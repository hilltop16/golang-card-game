package main

import "fmt"

type deck []string

func newDeck() deck {
	// create a deck slice deck := []string{}
	cards := deck{}
	// create suits slice
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	// create value slice
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}