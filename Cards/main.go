package main

func main() {
	//Since Go is statically typed, we need to declare variables along with data type as below
	// var card string = "Ace of Spades"
	//We can also declare variables as below directly and Go will understand the data type and in future
	//the variable should always hold the data type it first used
	// card := newCard()

	//there are 2 types of arrays in Go
	//Array is an array of fixed length
	//slice is an array which can grow/shrink

	//creating a slice
	cards := deck{}
	cards = newDeck()

	// fmt.Println(cards)
	//let us iterate through the slice of cards
	cards.printDeck()

}

//functions should also have a return type
func newCard() string {
	return "Ace of Spades"
}
