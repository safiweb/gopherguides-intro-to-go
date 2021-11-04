package assignment04

import (
	"fmt"
	"os"
	"testing"
)

func TestVenue_Entertain(t *testing.T) {

	v := Venue{Audience: 5, Log: os.Stdout}

	var tests = []struct {
		name     string
		audience int
		acts     []Entertainer
		err      error
	}{
		{"no entertainer", 0, []Entertainer{}, fmt.Errorf("there are no entertainers to perform")},
		{"venue and entertiners booked", 20, []Entertainer{Artist{StageName: "Wizkid"}, Band{StageName: "Destruction Boyz"}}, nil},
		{"empty venue for artist", 0, []Entertainer{Artist{StageName: "Rihanna"}}, fmt.Errorf("Rihanna: Rihanna has 0 audience to complete setup")},
		{"empty venue for band", 0, []Entertainer{Band{StageName: "MMM"}}, fmt.Errorf("MMM: MMM cannot perform for 0 audience")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Entertain(tt.audience, tt.acts...)
			if err != nil && err.Error() != tt.err.Error() {
				t.Fatalf("unexpected error: %v", err)
				return
			}
		})
	}
}
