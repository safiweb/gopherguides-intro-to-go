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

func (t *Theater) Play(viewers int, m ...*Movie) error {

	if len(m) == 0 {
		return fmt.Errorf("no movies to play")
	}

	for _, v := range m {
		v.Play(viewers)
	}

	return nil
}

func (t *Theater) Critique(fn CritiqueFn, m []*Movie) error {
	for _, v := range m {
		v.Play(1)
		if rate, err := fn(v); err != nil {
			return fmt.Errorf("no rate provided or exist")
		} else {
			if err := v.Rate(rate); err != nil {
				return fmt.Errorf("oh no, something went wrong! %v", err)
			}
		}
	}
	return nil
}

func main() {
	testCases := []Movie{
		{Name: "Gladiator", Length: 155, plays: 4},
		{Name: "Léon: The Professional", Length: 110},
	}

	for _, v := range testCases {
		got := v.Rate(10.2)
		want := error(nil)
		fmt.Println(v.plays)
		if got != want {
			fmt.Printf("no movies to play,play:%s", got)
		}
	}
}
