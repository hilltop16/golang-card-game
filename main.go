package main



func main() {
	//cards := deck{"ace of spade", showCard()}
	//cards = append(cards, "ace of diamond")
	//cards.print()
	//cards := newDeck()
	//hand, remainingCards := deal(cards, 3)
	//hand.print()
	//remainingCards.print()
	//cards.saveToFile("my_deck")
	//cards := newDeckFromFile("my_deck")
	//fmt.Println(cards)
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

//func showCard() string{
//	return "ace of heart"
//}