package main

func main() {
	// Declaring slices
	cards := deck{"Five of Diamonds", newCard()}
	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}
