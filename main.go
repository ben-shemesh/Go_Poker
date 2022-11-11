package main

import (
	"fmt"
	"github.com/ben-shemesh/Go_Poker/deck"
)

func main() {
	card := deck.NewCard(deck.Spades, 1)
	fmt.Println(card)
}