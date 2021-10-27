package assignment04

import (
	"bytes"
	"testing"
)

func TestArtistName(t *testing.T) {

	var tests = []struct {
		name string
		a    Artist
		exp  string
	}{
		{"failure", Artist{Artistname: "Rihanna"}, "Beyonce"},
		{"success", Artist{Artistname: "Rihanna"}, "Rihanna"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Artistname
			exp := tt.exp
			if got != exp {
				t.Fatalf("expected %s, got %s", exp, got)
			}
		})
	}

}

func TestArtistPerform(t *testing.T) {

	var tests = []struct {
		name string
		a    Artist
		v    Venue
	}{
		{"failure", Artist{Artistname: "Rihanna"}, Venue{}},
		{"success", Artist{Artistname: "Rihanna"}, Venue{Audience: 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			buff := bytes.Buffer{}
			tt.v.Log = &buff

			err := tt.a.Perform(tt.v)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			exp := tt.a.Artistname + " has completed performing.\n"
			got := buff.String()

			if exp != got {
				t.Fatalf("unexpected error, got: %v expected: %v", got, exp)
			}
		})
	}
}

func TestArtistSetup(t *testing.T) {

	var tests = []struct {
		name string
		a    Artist
		v    Venue
	}{
		{"failure", Artist{Artistname: "Rihanna"}, Venue{}},
		{"success", Artist{Artistname: "Rihanna"}, Venue{Audience: 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			buff := bytes.Buffer{}
			tt.v.Log = &buff

			err := tt.a.Setup(tt.v)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			exp := tt.a.Artistname + " has completed setup.\n"
			got := buff.String()

			if exp != got {
				t.Fatalf("unexpected error, got: %v expected: %v", got, exp)
			}
		})
	}
}
