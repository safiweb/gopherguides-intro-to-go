package assignment04

import (
	"bytes"
	"testing"
)

func TestBandName(t *testing.T) {

	var tests = []struct {
		name string
		b    Band
		exp  string
	}{
		{"failure", Band{Bandname: "Maroon 5"}, "Boyz 2 Men"},
		{"success", Band{Bandname: "Boyz 2 Men"}, "Boyz 2 Men"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.Bandname
			exp := tt.exp
			if got != exp {
				t.Fatalf("expected %s, got %s", exp, got)
			}
		})
	}

}

func TestBandPerform(t *testing.T) {

	var tests = []struct {
		name string
		b    Band
		v    Venue
	}{
		{"failure", Band{Bandname: "Boyz 2 Men"}, Venue{}},
		{"success", Band{Bandname: "Maroon 5"}, Venue{Audience: 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			buff := bytes.Buffer{}
			tt.v.Log = &buff

			err := tt.b.Perform(tt.v)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			exp := tt.b.Bandname + " has completed performing.\n"
			got := buff.String()

			if exp != got {
				t.Fatalf("unexpected error, got: %v expected: %v", got, exp)
			}
		})
	}
}

func TestBandTeardown(t *testing.T) {
	var tests = []struct {
		name string
		b    Band
		v    Venue
	}{
		{"failure", Band{Bandname: "Boyz 2 Men"}, Venue{}},
		{"success", Band{Bandname: "Maroon 5"}, Venue{Audience: 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			buff := bytes.Buffer{}
			tt.v.Log = &buff

			err := tt.b.Teardown(tt.v)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			exp := tt.b.Bandname + " has completed teardown.\n"
			got := buff.String()

			if exp != got {
				t.Fatalf("unexpected error, got: %v expected: %v", got, exp)
			}
		})
	}
}
