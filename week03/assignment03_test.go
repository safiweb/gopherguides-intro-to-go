package main

import (
	"fmt"
	"math"
	"testing"
)

func TestMovie_Rate(t *testing.T) {

	var tests = []struct {
		name string
		a    *Movie
		want error
	}{
		{"success", &Movie{Name: "Gladiator", plays: 4}, nil},
		{"failure", &Movie{Name: "Léon: The Professional"}, fmt.Errorf("can't review a movie without watching it first")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Rate(50.2)
			if got != nil {
				if got.Error() != tt.want.Error() {
					t.Fatalf("no movies to play: %s", got)
				}
			}
		})

	}
}

func TestMovie_Play(t *testing.T) {

	var tests = []struct {
		name    string
		movie   *Movie
		viewers int
		exp     int
	}{
		{"empty movie", &Movie{}, 5, 5},
		{"existing movie", &Movie{Name: "Léon: The Professional", viewers: 2}, 5, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.movie.Play(tt.viewers)
			if tt.movie.viewers != tt.exp {
				t.Fatalf("the %s movie views %d, expected: %d.", tt.movie.Name, tt.movie.viewers, tt.exp)
			}
		})

	}
}

func TestViewers(t *testing.T) {
	var tests = []struct {
		name string
		a    Movie
		exp  int
	}{
		{"empty movie", Movie{}, 0},
		{"existing movie", Movie{Name: "Léon: The Professional", viewers: 4}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Viewers()
			if got != tt.exp {
				t.Fatalf("the %s movie viewers %v, is not a vieweble number.", tt.a.Name, got)
			}
		})

	}
}

func TestPlays(t *testing.T) {
	var tests = []struct {
		name string
		a    Movie
		exp  int
	}{
		{"empty movie", Movie{}, 0},
		{"existing movie", Movie{Name: "Léon: The Professional", plays: 4}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Plays()
			if got != tt.exp {
				t.Fatalf("the %s movie plays %v, is not a playable number.", tt.a.Name, got)
			}
		})

	}
}

func TestMovie_Rating(t *testing.T) {
	var tests = []struct {
		name string
		a    Movie
		exp  float64
	}{
		{"unreleased movie", Movie{Name: "Gladiator"}, 0},
		{"box office movie", Movie{Name: "Léon: The Professional", rating: 100.0, plays: 5}, 20.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.a.Rating()

			if math.IsNaN(got) {
				got = 0
			}

			if got != tt.exp {
				t.Fatalf("the %s movie rating %.1f should be greater than 0, %v is not a rating number.", tt.a.Name, got, got)
			}
		})

	}
}

func TestMovie_String(t *testing.T) {
	var tests = []struct {
		name string
		a    Movie
		want string
	}{
		{"unreleased movie", Movie{Name: "Gladiator", Length: 155}, "Gladiator (155m) 0.0%"},
		{"released movie", Movie{Name: "Gladiator", Length: 155, rating: 90.5}, "Gladiator (155m) 90.5%"},
		{"movie in writing stage", Movie{}, " (0m) 0.0%"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.String()
			if got != tt.want {
				t.Fatalf("error outputing movie detail,got: %s, expected:%s ", got, tt.want)
			}
		})

	}
}

func TestTheater_Play(t *testing.T) {

	var tests = []struct {
		name    string
		theater *Theater
		b       *Movie
		viewers int
		want    error
	}{
		{"closed theater", &Theater{}, &Movie{}, 0, nil},
		{"active theater", &Theater{}, &Movie{Name: "Batman"}, 5, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.theater.Play(tt.viewers, tt.b)
			if got != tt.want {
				t.Fatalf("error playing movie detail,got: %v, expected:%v ", got, tt.want)
			}
		})
	}
}

func TestTheater_Critique(t *testing.T) {

	var tests = []struct {
		name string
		a    *Theater
		b    []*Movie
		want error
	}{
		{"unreleased movies", &Theater{}, []*Movie{{Name: "Gladiator"}, {Name: "Léon: The Professional"}}, fmt.Errorf("no rate provided or exist")},
		{"released movies with zero play", &Theater{}, []*Movie{{Name: "Gladiator 2", rating: 90, plays: 0}}, nil},
		{"released movies with some play", &Theater{}, []*Movie{{Name: "Gladiator 2", rating: 90}}, fmt.Errorf("no movies to play")},
	}

	fn := func(a *Movie) (float32, error) {

		rating := a.rating

		return rating, nil
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.a.Critique(fn, tt.b)

			if got != nil {
				if tt.want.Error() != got.Error() {
					t.Fatalf("unexpected error, got: %v, expected:%v ", got, tt.want)
				}
			}
		})
	}

}
