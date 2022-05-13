package main

import (
	"fmt"
)

func main() {
	
	// var card string = "Ace of Spades"

	/*  
	var				->	We're about to create a new variable
	card			->	The name of the variable will be 'greeting'
	string			->	Only a "string" will ever be assigned to this variable
	"Ace of Spades"	->	Assign the value "Ace of Spades" to this variable
	
	[!] Go is statically typed language
	*/


	// Because we assign value to the variable, we can do instead:
	card := "Ace of Spades"

	// [!] We only use this colon equals syntax when we are defining a new variable

	card = newCard()
	
	fmt.Println(card)


	// -----------------------------------------	Array, Slices -------------------------------------------

	cards := []string{newCard(), "Ace of Diamonds"}

	cards = append(cards, "Six of Spades") // we pass in our existing slice and then new element


	// ---------------------------------------------- For Loop ----------------------------------------------
	for i, card := range cards {
		fmt.Println(i, card)
	}
	/* 
	i			->	index of this element in the array
	card		->	current card we're iterating over
	range cards ->	take the slice of 'cards' and loop over it
	*/
}


func newCard() string {
	// We need to update our function declaration right here to inform the compiler that whenever the new
	// card function is executed it's going to return a value of type String
	
	return "King of Hearts"
}


// --------------------------------------------------------------------------------------------------------------------

/* Basic Go Types

	bool	-> true, false
	string	-> "Hi!", "Hows it going"
	int		-> 0, -1000, 99999
	float64	-> 10.000001, 0.00009, -100.003

*/


/* Two basic data structures of Go

	Array	->	Fixed length list of things

	Slice	->	An array that can grow or shrink
				[!] Every element in a slice must be of same type
*/