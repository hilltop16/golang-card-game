package main


func main() {
	//cards := deck{"ace of spade", showCard()}
	//cards = append(cards, "ace of diamond")
	//cards.print()
	cards := newDeck()
	hand, remainingCards := deal(cards, 3)
	hand.print()
	remainingCards.print()
}

//func showCard() string{
//	return "ace of heart"
//}