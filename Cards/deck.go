package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

//function to deal cards of particular hand size. This function does not take receiver

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//convert deck of cards to string

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 777)
}

func readFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Critical error occurred: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

//shuffle the deck of cards by swapping positions of current index with random index within deck length
func (d deck) shuffle() deck {

	//seed the random object with current time in nanoseconds
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
