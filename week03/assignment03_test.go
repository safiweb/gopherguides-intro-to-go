package main

import (
	"fmt"
	"testing"
)

func TestRate(t *testing.T) {

	var tests = []struct {
		name string
		a    Movie
		want error
	}{
		{"success", Movie{Name: "Gladiator", Length: 155, plays: 4}, nil},
		{"failure", Movie{Name: "Léon: The Professional", Length: 110}, fmt.Errorf("can't review a movie without watching it first")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Rate(50.2)
			if got != tt.want {
				t.Errorf("no movies to play: %s", got)
			}
		})

	}
}

func TestPlay(t *testing.T) {
	var tests = []struct {
		name      string
		a         Movie
		b         int
		want_view int
		want_play int
	}{
		{"success", Movie{Name: "Gladiator", Length: 155}, 5, 5, 1},
		{"failure", Movie{Name: "Léon: The Professional", Length: 110}, 5, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Play(tt.b)
			if tt.a.viewers != tt.want_view {
				t.Errorf("the %s movie views %d, expected: %d.", tt.a.Name, tt.a.viewers, tt.want_view)
			}
			if tt.a.plays != tt.want_play {
				t.Errorf("the %s movie plays %d, expected: %d.", tt.a.Name, tt.a.plays, tt.want_play)
			}
		})

	}
}

func TestViewers(t *testing.T) {
	var tests = []struct {
		name string
		a    Movie
	}{
		{"failure", Movie{Name: "Gladiator", Length: 155, viewers: -1}},
		{"success", Movie{Name: "Léon: The Professional", Length: 110, viewers: 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Viewers()
			if got < 0 {
				t.Errorf("the %s movie viewers %v, is not a vieweble number.", tt.a.Name, got)
			}
		})

	}
}

func TestPlays(t *testing.T) {
	var tests = []struct {
		name string
		a    Movie
	}{
		{"failure", Movie{Name: "Gladiator", Length: 155, plays: -1}},
		{"success", Movie{Name: "Léon: The Professional", Length: 110, plays: 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Plays()
			if got < 0 {
				t.Errorf("the %s movie plays %v, is not a playable number.", tt.a.Name, got)
			}
		})

	}
}

func TestRating(t *testing.T) {
	var tests = []struct {
		name string
		a    Movie
	}{
		{"success", Movie{Name: "Gladiator", Length: 155, rating: 90.5, plays: 2}},
		{"failure", Movie{Name: "Léon: The Professional", Length: 110, rating: 0.0, plays: 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Rating()
			if got <= 0 {
				t.Errorf("the %s movie rating %.1f should be greater than 0, %v is not a rating number.", tt.a.Name, got, got)
			}
		})

	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		name string
		a    Movie
		want string
	}{
		{"success", Movie{Name: "Gladiator", Length: 155, rating: 90.5}, "Gladiator (155m) 90.5%"},
		{"failure", Movie{Name: "Léon: The Professional", Length: 110, rating: 4.5}, "Léon: The Professional (110m) 68.9%"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.String()
			if got != tt.want {
				t.Errorf("error outputing movie detail,got: %s, expected:%s ", got, tt.want)
			}
		})

	}
}

func TestPlayTheater(t *testing.T) {
	var tests = []struct {
		name string
		a    Theater
		b    Movie
		c    int
		want int
	}{
		{"failure", Theater{}, Movie{Name: "Gladiator", Length: 155}, -2, 0},
		{"success", Theater{}, Movie{Name: "Batman", Length: 155}, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Play(tt.c, &tt.b)
			got := tt.b.viewers
			if got != tt.want {
				t.Errorf("error playing movie detail,got: %v, expected:%v ", got, tt.want)
			}
		})
	}
}

func TestCritique(t *testing.T) {

	var tests = []struct {
		name    string
		a       Theater
		b       []*Movie
		cRating float32
		want    error
	}{
		{"success", Theater{}, []*Movie{{Name: "Gladiator", Length: 155}, {Name: "Léon: The Professional", Length: 110}}, 80.0, nil},
		{"failure", Theater{}, []*Movie{{Name: "Gladiator 2", Length: 155}, {Name: "Léon: The Professional", Length: 110}}, -1, nil},
		{"failure 2", Theater{}, []*Movie{}, 80, fmt.Errorf("no movies to play")},
	}

	var critiqueRating float32

	fn := func(a *Movie) (float32, error) {

		if a.Name == "" && a.Length == 0 {
			return 0, fmt.Errorf("the movie is nil/empty")
		}

		return float32(critiqueRating), nil
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if len(tt.b) == 0 {
				t.Errorf("No movies in theater")
			}

			for _, v := range tt.b {
				v.Play(1)
				//fn(v)
				critiqueRating = tt.cRating

				rate, _ := fn(v)

				if rate <= 0 {
					t.Errorf("Critique rating not eligible,got: %v, expected:>0", rate)
				}

				if err := v.Rate(rate); err != nil {
					t.Errorf("oh no, something went wrong! %v", err)
				}

			}
		})
	}

}
