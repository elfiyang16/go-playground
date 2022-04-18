package main

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of spades" // init, dynamic inference, only used with new var
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards = append(cards, "Size of spades")
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// // greeting := "hellow"
	// // fmt.Println([]byte(greeting)) //[104 101 108 108 111 119]
	// newCards := newDeckFromFile("my_cards")
	// fmt.Println("new cards", newCards)
}

func newCard() string {
	return "Five of diamonds"
}
