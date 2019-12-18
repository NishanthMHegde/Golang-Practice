package main

import (
	"os"
	"testing"
)

//test to check if the new deck has the required number of cards
//Always start the name of the fuction with Test, with T in caps
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("The length of card was not 52, found length %v", len(d))
	}
}

//we now test if the cards in different positions are as how we expect
func TestCardPosition(t *testing.T) {
	d := newDeck()
	if d[0] != "Ace of Spades" {
		t.Errorf("The first card was not Ace of Spades, but rather it was %v", d[0])
	}
	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("The last card was not King of Hearts, but rather it was %v", d[0])
	}
}

//we test whether our save to file and read from file are working. We save our deck to _decktesting and then remove it after the test

func TestSaveToFileAndReadFromFile(t *testing.T) {
	//remove created deck file
	os.Remove("_decktesting")
	cards := newDeck()
	cards.saveToFile("_decktesting")

	//now try to read the file
	readCards := readFromFile("_decktesting")
	if len(readCards) != 52 {
		t.Errorf("Read from file failed")
	}
}
