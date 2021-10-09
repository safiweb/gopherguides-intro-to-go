package main

import "testing"

func TestAssertArrayEquals(t *testing.T) {
	exp := [4]string{"John", "Paul", "George", "Ringo"}
	act := [len(exp)]string{}

	//Copy exp values into the act variable
	for i, v := range exp {
		act[i] = v
	}

	//Check if act array content equals to exp
	for i, v := range act {
		if v != exp[i] {
			t.Fatalf("unexpected value, got: %s, exp: %s at index %v", v, exp[i], i)
		}
	}
}

func TestAssertSliceEquals(t *testing.T) {
	exp := []string{"John", "Paul", "George", "Ringo"}
	act := []string{}

	//Copy exp values into the act variable
	for _, v := range exp {
		act = append(act, v)
	}

	//Check if length of act and exp are the same
	if len(exp) != len(act) {
		t.Fatalf("unexpected value, got: %d, exp: %d", len(act), len(exp))
	}

	//Check if act slice content equals to exp
	for i, v := range act {
		if v != exp[i] {
			t.Fatalf("unexpected value, got: %s, exp: %s at index %v", v, exp[i], i)
		}
	}
}