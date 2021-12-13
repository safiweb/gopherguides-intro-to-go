package newsmill

// Subscription to receive news stories for the categories they are subscribed to.
type Subscription string

// IsValid returns an error if the subscription is not valid.
func (sub Subscription) IsValid() error {

	if len(sub) > 0 {
		return nil
	}

	return ErrInvalidSubscription(sub)
}
