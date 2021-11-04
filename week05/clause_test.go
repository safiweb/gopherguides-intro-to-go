package week05

import (
	"testing"
)

func TestClauses_String(t *testing.T) {

	t.Parallel()

	testcases := []struct {
		name string
		cls  Clauses
		want string
	}{
		{
			name: "no clauses provided",
			cls:  Clauses{},
			want: "",
		},
		{
			name: "clauses provided",
			cls:  Clauses{"animals": "Lion", "sports": "F1"},
			want: `"animals" = "Lion" and "sports" = "F1"`,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.cls.String()
			if got != tt.want {
				t.Fatalf("unexpected value, got: %v, exp: %v", got, tt.cls)
			}

		})
	}
}

func TestClauses_Match(t *testing.T) {

	t.Parallel()

	testcases := []struct {
		name string
		cls  Clauses
		m    Model
		want bool
	}{
		{
			name: "nil data provided",
			cls:  Clauses{"animals": "Lion", "sports": "F1", "person": "Anthony"},
			m:    Model{},
			want: false,
		},
		{
			name: "matching data provided",
			cls:  Clauses{"animals": "Lion", "sports": "F1", "person": "Anthony"},
			m:    Model{"sports": "F1", "person": "Anthony", "animals": "Lion"},
			want: true,
		},
		{name: "no matching data provided",
			cls:  Clauses{"animals": "Lion", "sports": "F1", "person": "Anthony"},
			m:    Model{"sports": "F1", "person": "Webbs", "animals": "Lion"},
			want: false,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.cls.Match(tt.m)
			if got != tt.want {
				t.Fatalf("unexpected value, got: %v, exp: %v", got, tt.want)
			}

		})
	}
}
