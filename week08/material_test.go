package week08

import (
	"testing"
	"time"
)

func TestMaterial_Duration(t *testing.T) {
	t.Parallel()

	mats := Materials{
		Oil:  1,
		Wood: 1,
	}

	exp := time.Duration(7) * time.Millisecond

	got := mats.Duration()
	if got != exp {
		t.Fatalf("expected %v, got %v", exp, got)
	}

}
