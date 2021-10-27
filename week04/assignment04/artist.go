package assignment04

import "fmt"

type Artist struct {
	Artistname string
	Genre      string
}

func (a Artist) Name() string {
	return a.Artistname
}

func (a Artist) Perform(v Venue) error {
	if v.Audience == 0 {
		return fmt.Errorf("%v cannot perform for %d audience", a.Artistname, v.Audience)
	}

	fmt.Fprintf(v.Log, "%s has completed performing.\n", a.Artistname)

	return nil
}

func (a Artist) Setup(v Venue) error {
	if v.Audience == 0 {
		return fmt.Errorf("%v has %d audience to complete setup", a.Artistname, v.Audience)
	}

	fmt.Fprintf(v.Log, "%s has completed setup.\n", a.Artistname)

	return nil
}
