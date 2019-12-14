package main

import "fmt"

//Go is not an object oriented programming language
//Instead of creating a deck class, we create a deck 'type'
// This deck 'type' is a slice type which can have its own set of functions

type deck []string

//we create a receiver function which prints the deck of cards
func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//function to create a deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	//create a deck of values of suites
	for _, value := range cardValues {
		for _, suite := range cardSuites {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}
