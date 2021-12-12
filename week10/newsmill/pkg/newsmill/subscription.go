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

/*
//Subscriber listen for news
func (sub *Subscription) listen(ctx context.Context, s *Service) {
	for {

		// listen for messages on different channels
		select {
		case <-ctx.Done(): // listen context cancellation
			return
		case article, ok := <-s.Broadcast(): // listen for articles

			// check if the channel is closed or not
			if !ok {
				continue
			}

			fmt.Println(article)


				// try to build the product
				err := p.Build(e, m.Warehouse)
				if err != nil {
					// if there is an error, send it to the manager
					m.Errors() <- err
					continue
				}

				// if there is no error, send the product back to the manager
				m.Complete(e, p)

		}
	}
}
*/
