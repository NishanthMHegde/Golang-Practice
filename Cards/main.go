package main

import "fmt"

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
	fmt.Println("Original Deck of cards")
	cards.printDeck()
	hand, remainingCards := deal(cards, 4)
	fmt.Println("Hand \n\n\n")
	hand.printDeck()
	fmt.Println("Remaining cards \n\n\n")
	remainingCards.printDeck()

	//to save deck of cards to file, we must first convert it to byte slice
	//to convert to byte slice, we must first convert it to string slice and then to a , seperated string
	cards.saveToFile("my_cards")
	fmt.Println("\n\n\n")
	fmt.Println("Cards read from file are ")
	read_cards := readFromFile("my_cards")
	read_cards.printDeck()

	fmt.Println("\n\n\n")

	fmt.Println("Shuffled Deck is ")
	cards = cards.shuffle()
	cards.printDeck()

}
