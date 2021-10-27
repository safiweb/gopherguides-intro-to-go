package assignment04

import "fmt"

type Band struct {
	Bandname string
	Members  int
}

func (b Band) Name() string {
	return b.Bandname
}

func (b Band) Perform(v Venue) error {
	if v.Audience == 0 {
		return fmt.Errorf("%v cannot perform for %d audience", b.Bandname, v.Audience)
	}

	fmt.Fprintf(v.Log, "%s has completed performing.\n", b.Bandname)

	return nil
}

func (b Band) Teardown(v Venue) error {
	if v.Audience == 0 {
		return fmt.Errorf("%v cannot complete teardown for %d audience", b.Bandname, v.Audience)
	}
	fmt.Fprintf(v.Log, "%s has completed teardown.\n", b.Bandname)
	return nil
}
