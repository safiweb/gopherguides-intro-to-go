package assignment04

import (
	"bytes"
	"os"
	"testing"
)

func TestEntertain(t *testing.T) {
	v := Venue{Audience: 5, Log: os.Stdout}

	var tests = []struct {
		name     string
		audience int
		acts     []Entertainer
	}{
		{"failure", 0, []Entertainer{
			Artist{StageName: "Burna Boy"},
			Band{StageName: "Sauti Soul"},
		}},
		{"success", 20, []Entertainer{
			Artist{StageName: "Wizkid"},
			Band{StageName: "Destruction Boyz"},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, act := range tt.acts {
				err := v.Entertain(tt.audience, act)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}

		})
	}
}

func TestPlay(t *testing.T) {
	var tests = []struct {
		name string
		v    Venue
		acts []Entertainer
	}{
		{"failure", Venue{Audience: 0}, []Entertainer{
			Artist{StageName: "Burna Boy"},
			Band{StageName: "Sauti Soul"},
		}},
		{"success", Venue{Audience: 12}, []Entertainer{
			Artist{StageName: "Wizkid"},
			Band{StageName: "Destruction Boyz"},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, act := range tt.acts {

				buff := bytes.Buffer{}
				tt.v.Log = &buff

				err := tt.v.play(act)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}

}
