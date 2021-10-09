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