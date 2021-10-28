package assignment04

import "fmt"

type Artist struct {
	StageName string
	Genre     string
}

func (a Artist) Name() string {
	return a.StageName
}

func (a Artist) Perform(v Venue) error {
	if v.Audience == 0 {
		return fmt.Errorf("%v cannot perform for %d audience", a.StageName, v.Audience)
	}

	fmt.Fprintf(v.Log, "%s has completed performing.\n", a.StageName)

	return nil
}

func (a Artist) Setup(v Venue) error {
	if v.Audience == 0 {
		return fmt.Errorf("%v has %d audience to complete setup", a.StageName, v.Audience)
	}

	fmt.Fprintf(v.Log, "%s has completed setup.\n", a.StageName)

	return nil
}
