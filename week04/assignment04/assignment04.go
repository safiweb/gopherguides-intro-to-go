package assignment04

import (
	"fmt"
	"io"
)

type Entertainer interface {
	Name() string
	Perform(v Venue) error
}

type Setuper interface {
	Setup(v Venue) error
}

type Teardowner interface {
	Teardown(v Venue) error
}

type Venue struct {
	Audience int
	Log      io.Writer
}

func (v Venue) Name() string {
	return "Drake"
}

func (v Venue) Perform(p Venue) error {
	if p.Audience <= 0 {
		return fmt.Errorf("%v cannot perform for 0 audience", p.Name())
	}
	return nil
}

func (v Venue) Setup(p Venue) error {
	if p.Audience <= 0 {
		return fmt.Errorf("%v has no audience to complete setup", p.Name())
	}

	return nil
}

func (v *Venue) Entertain(audience int, acts ...Entertainer) error {
	if len(acts) == 0 {
		return fmt.Errorf("there are no entertainers to perform")
	}

	v.Audience = audience
	for _, act := range acts {
		if err := v.play(act); err != nil {
			return err
		}
	}

	return nil
}

func (v Venue) play(act Entertainer) error {

	name := act.Name()

	if s, ok := act.(Setuper); ok {
		if err := s.Setup(v); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
		fmt.Fprintf(v.Log, "%s has completed setup.\n", name)
	}

	if err := act.Perform(v); err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}

	fmt.Fprintf(v.Log, "%s has performed for %d people.\n", name, v.Audience)

	if t, ok := act.(Teardowner); ok {
		if err := t.Teardown(v); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
		fmt.Fprintf(v.Log, "%s has completed teardown.\n", name)
	}

	return nil
}
