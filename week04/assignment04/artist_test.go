package assignment04

import (
	"bytes"
	"fmt"
	"testing"
)

func TestArtist_Name(t *testing.T) {

	var tests = []struct {
		name string
		a    Artist
		exp  string
	}{
		{"non existing artist", Artist{}, ""},
		{"existing artist", Artist{StageName: "Rihanna"}, "Rihanna"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.StageName
			exp := tt.exp
			if got != exp {
				t.Fatalf("expected %s, got %s", exp, got)
			}
		})
	}

}

func TestArtist_Perform(t *testing.T) {

	var tests = []struct {
		name string
		a    Artist
		v    Venue
		exp  error
	}{
		{"no venue to perform", Artist{StageName: "Rihanna"}, Venue{}, fmt.Errorf("Rihanna cannot perform for 0 audience")},
		{"venue to perform", Artist{StageName: "Beyonce"}, Venue{Audience: 20}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			buff := bytes.Buffer{}
			tt.v.Log = &buff

			err := tt.a.Perform(tt.v)

			if err != nil {
				if err.Error() != tt.exp.Error() {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}

			exp := tt.a.StageName + " has completed performing.\n"
			got := buff.String()

			if exp != got {
				t.Fatalf("unexpected error, got: %v expected: %v", got, exp)
			}
		})
	}
}

func TestArtist_Setup(t *testing.T) {

	var tests = []struct {
		name string
		a    Artist
		v    Venue
		exp  error
	}{
		{"no venue to set up", Artist{StageName: "Rihanna"}, Venue{}, fmt.Errorf("Rihanna has 0 audience to complete setup")},
		{"artist has venue to set up", Artist{StageName: "Rihanna"}, Venue{Audience: 20}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			buff := bytes.Buffer{}
			tt.v.Log = &buff

			err := tt.a.Setup(tt.v)
			if err != nil {
				if err.Error() != tt.exp.Error() {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}

			exp := tt.a.StageName + " has completed setup.\n"
			got := buff.String()

			if exp != got {
				t.Fatalf("unexpected error, got: %v expected: %v", got, exp)
			}
		})
	}
}
