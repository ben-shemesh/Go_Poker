package deck

type Suite int

const (
	Spades Suite = iota
	Harts
	Diamonds
	Clubs
)

type Card struct {
	Suite Suite
	Value int
}

func NewCard (s Suite, v int ) Card {
	if v > 14 {
		panic("the value of the card cannot be higher than 13.")
	}
	return Card{
		Suite: s,
		Value: v,
	}
} 