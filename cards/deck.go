package main

import "fmt"

// Create a new data type 'deck'
// which is a slice of string
type deck []string

// Receiver functions
// It is tied with the passed data
// Acts sort of like a class function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}