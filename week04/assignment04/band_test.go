package assignment04

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBand_Name(t *testing.T) {

	var tests = []struct {
		name string
		b    Band
		exp  string
	}{
		{"no artist", Band{}, ""},
		{"existing artist", Band{StageName: "Boyz 2 Men"}, "Boyz 2 Men"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.StageName
			exp := tt.exp
			if got != exp {
				t.Fatalf("expected %s, got %s", exp, got)
			}
		})
	}

}

func TestBand_Perform(t *testing.T) {

	var tests = []struct {
		name string
		b    Band
		v    Venue
		exp  error
	}{
		{"no venue to perform", Band{StageName: "Boyz 2 Men"}, Venue{}, fmt.Errorf("Boyz 2 Men cannot perform for 0 audience")},
		{"has venue to perform", Band{StageName: "Maroon 5"}, Venue{Audience: 20}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			buff := bytes.Buffer{}
			tt.v.Log = &buff

			err := tt.b.Perform(tt.v)
			if err != nil {
				if err.Error() != tt.exp.Error() {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}

			exp := tt.b.StageName + " has completed performing.\n"
			got := buff.String()

			if exp != got {
				t.Fatalf("unexpected error, got: %v expected: %v", got, exp)
			}
		})
	}
}

func TestBand_Teardown(t *testing.T) {
	var tests = []struct {
		name string
		b    Band
		v    Venue
		err  error
	}{
		{"no stage venue to teardown", Band{StageName: "Boyz 2 Men"}, Venue{}, fmt.Errorf("Boyz 2 Men cannot complete teardown for 0 audience")},
		{"stage venue to teardown", Band{StageName: "Maroon 5"}, Venue{Audience: 20}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			buff := bytes.Buffer{}
			tt.v.Log = &buff

			err := tt.b.Teardown(tt.v)
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}

			exp := tt.b.StageName + " has completed teardown.\n"
			got := buff.String()

			if exp != got {
				t.Fatalf("unexpected error, got: %v expected: %v", got, exp)
			}
		})
	}
}
