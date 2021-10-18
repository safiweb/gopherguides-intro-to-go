package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRate(t *testing.T) {

	var tests = []struct {
		a    Movie
		want error
	}{
		{Movie{Name: "Gladiator", Length: 155, plays: 4}, nil},
		{Movie{Name: "Léon: The Professional", Length: 110}, fmt.Errorf("can't review a movie without watching it first")},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := tt.a.Rate(50.2)
			if got != tt.want {
				t.Errorf("no movies to play: %s", got)
			}
		})

	}
}

func TestPlay(t *testing.T) {
	var tests = []struct {
		a    Movie
		b    int
		want int
	}{
		{Movie{Name: "Gladiator", Length: 155}, 5, 5},
		{Movie{Name: "Léon: The Professional", Length: 110}, 5, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tt.a.Play(tt.b)
			if tt.a.viewers != tt.want {
				t.Errorf("the %s movie views %d, expected: %d.", tt.a.Name, tt.a.viewers, tt.want)
			}
			if tt.a.plays != tt.want {
				t.Errorf("the %s movie plays %d, expected: %d.", tt.a.Name, tt.a.plays, tt.want)
			}
		})

	}
}

func TestViewers(t *testing.T) {
	var tests = []struct {
		a Movie
	}{
		{Movie{Name: "Gladiator", Length: 155, viewers: -1}},
		{Movie{Name: "Léon: The Professional", Length: 110, viewers: 4}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := tt.a.Viewers()
			if reflect.TypeOf(got) != reflect.TypeOf(0) || got < 0 {
				t.Errorf("the %s movie viewers %v, is not a vieweble number.", tt.a.Name, got)
			}
		})

	}
}

func TestPlays(t *testing.T) {
	var tests = []struct {
		a Movie
	}{
		{Movie{Name: "Gladiator", Length: 155, plays: -1}},
		{Movie{Name: "Léon: The Professional", Length: 110, plays: 4}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := tt.a.Plays()
			if reflect.TypeOf(got) != reflect.TypeOf(0) || got < 0 {
				t.Errorf("the %s movie plays %v, is not a playable number.", tt.a.Name, got)
			}
		})

	}
}

func TestRating(t *testing.T) {
	var tests = []struct {
		a Movie
	}{
		{Movie{Name: "Gladiator", Length: 155, rating: 90.5, plays: 2}},
		{Movie{Name: "Léon: The Professional", Length: 110, rating: 0.0, plays: 5}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := tt.a.Rating()
			if got <= 0 {
				t.Errorf("the %s movie rating %.1f should be greater than 0, %v is not a rating number.", tt.a.Name, got, got)
			}
		})

	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		a    Movie
		want string
	}{
		{Movie{Name: "Gladiator", Length: 155, rating: 90.5}, "Gladiator (155m) 90.5%"},
		{Movie{Name: "Léon: The Professional", Length: 110, rating: 4.5}, "Léon: The Professional (110m) 68.9%"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := tt.a.String()
			if got != tt.want {
				t.Errorf("error outputing movie detail,got: %s, expected:%s ", got, tt.want)
			}
		})

	}
}

func TestPlayTheater(t *testing.T) {
	var tests = []struct {
		a    Theater
		b    Movie
		want error
	}{
		{Theater{}, Movie{Name: "Gladiator", Length: 155}, nil},
		{Theater{}, Movie{}, fmt.Errorf("no movies to play")},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if err := tt.a.Play(15, &tt.b); err != nil {
				t.Errorf("no movies to play")
			}
		})
	}
}

func TestCritique(t *testing.T) {

	var tests = []struct {
		a    Theater
		b    []*Movie
		want error
	}{
		{Theater{}, []*Movie{{Name: "Gladiator", Length: 155}, {Name: "Léon: The Professional", Length: 110}}, nil},
		{Theater{}, []*Movie{{Name: "Gladiator 2", Length: 155}, {Name: "Léon: The Professional", Length: 110}}, nil},
		{Theater{}, []*Movie{}, fmt.Errorf("no movies to play")},
	}

	var critiqueRating float32

	fn := func(a *Movie) (float32, error) {

		if a.Name == "" && a.Length == 0 {
			return 0, fmt.Errorf("the movie is nil/empty")
		}

		return float32(critiqueRating), nil
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {

			for _, v := range tt.b {
				v.Play(1)
				//fn(v)
				critiqueRating -= 2
				if rate, err := fn(v); err != nil || rate <= 0 {
					t.Errorf("no rate provided or exist")
				} else {
					if err := v.Rate(rate); err != nil {
						t.Errorf("oh no, something went wrong! %v", err)
					}
				}

			}
		})
	}

}
