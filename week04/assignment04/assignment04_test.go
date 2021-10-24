package assignment04_test

import (
	. "gopherguides-intro-to-go/week04/assignment04"
	"os"
	"testing"
)

var globalVenueTests = []struct {
	name string
	v    Venue
	ok   error
}{
	{"success", Venue{Audience: 5, Log: os.Stdout}, nil},
	{"failure", Venue{Log: os.Stdout}, nil},
}

var globalEntertainerTests = []struct {
	name     string
	audience int
	acts     []Entertainer
	ok       error
}{
	{"failure", 0, []Entertainer{
		Venue{Audience: 0, Log: os.Stdout},
	}, nil},
	{"success", 20, []Entertainer{
		Venue{Audience: 5, Log: os.Stdout},
	}, nil},
}

func TestName(t *testing.T) {

	var tests = []struct {
		name string
		v    Venue
		want string
	}{
		{"success", Venue{}, "Drake"},
		{"failure", Venue{}, "Kanye"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Name()
			if got != tt.want {
				t.Errorf("unexpected value, got: %s, exp: %s", got, tt.want)
			}
		})

	}
}

func TestPerform(t *testing.T) {
	venue := Venue{}
	for _, tt := range globalVenueTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.v.Audience <= 0 {
				t.Errorf("%v cannot perform for 0 audience", tt.v.Name())
			}
			err := venue.Perform(tt.v)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestSetup(t *testing.T) {
	venue := Venue{}
	for _, tt := range globalVenueTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.v.Audience <= 0 {
				t.Errorf("%v cannot perform for 0 audience", tt.v.Name())
			}
			err := venue.Perform(tt.v)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestEntertain(t *testing.T) {
	v := Venue{Audience: 5, Log: os.Stdout}
	for _, tt := range globalEntertainerTests {
		t.Run(tt.name, func(t *testing.T) {

			if len(tt.acts) == 0 {
				t.Errorf("there are no entertainers to perform")
			}

			for _, act := range tt.acts {
				err := v.Entertain(tt.audience, act)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}

		})
	}
}

func TestPlay(t *testing.T) {
	var tests = []struct {
		name string
		v    Venue
		act  Entertainer
	}{
		{"success", Venue{Audience: 10, Log: os.Stdout}, Venue{}},
		{"failure", Venue{Audience: 0, Log: os.Stdout}, Venue{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			name := tt.act.Name()

			if name == "" {
				t.Errorf("name cannot be empty")
			}

			if s, ok := tt.act.(Setuper); ok {
				if err := s.Setup(tt.v); err != nil {
					t.Errorf("%s: %v", name, err)
				}
			}

			if err := tt.act.Perform(tt.v); err != nil {
				t.Errorf("%s: %v", name, err)
			}

			if k, ok := tt.act.(Teardowner); ok {
				if err := k.Teardown(tt.v); err != nil {
					t.Errorf("%s: %v", name, err)
				}
			}

		})
	}

}
