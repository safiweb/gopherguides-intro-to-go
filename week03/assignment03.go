package main

import (
	"fmt"
)

type Movie struct {
	Name    string
	Length  int
	rating  float32
	plays   int
	viewers int
}

type CritiqueFn func(m *Movie) (float32, error)

func (m *Movie) Rate(rating float32) error {

	if m.plays == 0 {
		return fmt.Errorf("can't review a movie without watching it first")
	}

	m.rating = m.rating + rating
	return nil

}

func (m *Movie) Play(viewers int) {
	m.viewers = m.viewers + viewers
	m.plays = m.plays + viewers
}

func (m Movie) Viewers() int {
	return m.viewers
}

func (m Movie) Plays() int {
	return m.plays
}

func (m Movie) Rating() float64 {
	ratings := m.rating / float32(m.plays)
	return float64(ratings)
}

func (m Movie) String() string {
	s := fmt.Sprintf("%v (%dm) %.1f%%", m.Name, m.Length, m.rating)
	return s
}

func main() {

}
