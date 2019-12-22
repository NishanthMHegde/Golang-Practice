package main

import "fmt"

//Interfaces : Go does not allow two functions with same name to be defined if they have the same implementations even though
// they might have different function argument data types.
//However, the two functions with same name can exist if they both have different argument types as a receiver function.

type englishBot struct{}
type frenchBot struct{}

// we create an interface named Bot
//This interface will contain a function which will have the ability to be called by our printGreeting() function
// the getGreeting() function can take in an argument of type Bot
//printGreeting() function has to take the parameter as Bot type which can be english or French bot and pass it on to the getGreeting() function
type Bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	fb := frenchBot{}
	printGreeting(eb)
	printGreeting(fb)
}

//the below functions have custom implementations
// the below functions have same name but take in receiver arguments for bot type, so no error is thrown
func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (fb frenchBot) getGreeting() string {
	return "Salut!"
}

//*****IMPORTANT****
//the below implementation of printGreeting will throw error because the functions have same name
// and also they do not take in the bot type as receiver argument

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(fb frenchBot) {
// 	fmt.Println(fb.getGreeting())
// }

//to solve above issue, we use an interface named Bot. This Bot interface will define a function which can take in any types of bots
//printGreeting() should now take in the input parameter as Bot interface instead of specific bots

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}
