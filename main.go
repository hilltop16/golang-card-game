package main


func main() {
	cards := deck{"ace of spade", showCard()}
	cards = append(cards, "ace of diamond")
	cards.print()
}

func showCard() string{
	return "ace of heart"
}