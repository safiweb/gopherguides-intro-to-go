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

type Theater struct{}

func (m *Movie) Rate(rating float32) error {

	if m.plays == 0 {
		return fmt.Errorf("can't review a movie without watching it first")
	}

	m.rating += rating
	return nil

}

func (m *Movie) Play(viewers int) {
	m.viewers += viewers
	m.plays++
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

func (t *Theater) Play(viewers int, movies ...*Movie) error {

	if len(movies) == 0 {
		return fmt.Errorf("no movies to play")
	}

	for _, v := range movies {
		v.Play(viewers)
	}

	return nil
}

func (t *Theater) Critique(fn CritiqueFn, m []*Movie) error {
	for _, v := range m {
		v.Play(1)

		rate, _ := fn(v)
		if rate <= 0 {
			return fmt.Errorf("no rate provided or exist")
		}

		if err := v.Rate(rate); err != nil {
			return fmt.Errorf("oh no, something went wrong! %w", err)
		}
	}

	return nil
}

func main() {}
