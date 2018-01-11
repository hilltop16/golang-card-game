package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(5)
	}
	s := deck(strings.Split(string(bs), ","))
	return s
}

//func (d deck) shuffle() deck {
//	cardIndices := rand.Perm(52)
//	newDeck := deck{}
//
//	for i := range d {
//		newDeck = append(newDeck, d[cardIndices[i]])
//	}
//
//	return newDeck
//}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1) // every time newPosition is assigned a new int
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}