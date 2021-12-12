package newsmill

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// Config holds the Service configuration settings
type Config struct{}

// Status holds the Service status and history
type Status struct{}

// Service is responsible for receiving subcriptions and broadcasting new stories to subcribers.
// Service is also responsible to get and manage new stories sources, create the new stories.
type Service struct {
	sync.RWMutex
	sync.WaitGroup
	broadcast         chan interface{}                  // broadcast channel
	errs              chan error                        // errors channel
	subs              map[Subscription]chan interface{} // subscriptions channel
	cancel            context.CancelFunc
	once              sync.Once
	stopped           bool
	Status            bool `json:"status"` // Status
	stateFile         string
	SubcribedArticles []Article `json:"articles"` //list of published and subscribed articles
}

// NewService will create a new Service.
// ALWAYS to be used to create a new Service
func NewService() *Service {
	return &Service{
		broadcast:         make(chan interface{}, 1),
		subs:              make(map[Subscription]chan interface{}),
		errs:              make(chan error),
		SubcribedArticles: []Article{},
		stateFile:         "testdata/news/status.json",
	}
}

// Start will start dispatching news for the given topics
func (s *Service) Start(ctx context.Context, topics ...string) (context.Context, error) {

	if len(topics) == 0 {
		return nil, ErrInvalidTopicCount(len(topics))
	}

	// create a new cancellation context
	ctx, cancel := context.WithCancel(ctx)

	// hold onto the cancel function so it can be called
	// by m.Stop()
	s.Lock()
	s.cancel = cancel
	s.Unlock()

	// launch a goroutine to listen context cancellation
	go func(ctx context.Context) {

		// listen for context cancellation
		// this could come from the external context
		// passed to s.Start()
		<-ctx.Done()

		// call the cancel function
		cancel()

		// call Stop()
		s.Stop()

	}(ctx)

	s.StateFileLoad()

	go s.disptach(ctx)

	return ctx, nil

}

// Subscribe a subscriber/subscription to the service
// Subscribe returns a new channel on which to receive articles on a certain topic.
func (s *Service) Subscribe(topic Subscription) (<-chan interface{}, error) {

	if err := topic.IsValid(); err != nil {
		return nil, err
	}

	s.Lock()
	if _, ok := s.subs[topic]; ok {
		return nil, ErrSubscriptionExist(fmt.Sprintf("subscription %s already exist", topic))
	}
	s.Unlock()

	s.Lock()
	ch := make(chan interface{}, 1)
	s.subs[topic] = ch
	s.Unlock()

	return ch, nil

}

//  Unsubscribe removes a subscriber from the service
//  Unsubscribe removes the subscribed topic channel from the subscription
//  meaning there will be no more news articles sent.
func (s *Service) Unsubscribe(topic Subscription) error {

	ch, ok := s.subs[topic]

	if !ok {
		return ErrSubscriptionNotFound(topic)
	}

	s.Lock()
	delete(s.subs, topic)
	close(ch)
	s.Unlock()

	return nil

}

// dispatch news articles to subscribers
// It get news articles from the broadcast channel and dispatches
func (s *Service) disptach(ctx context.Context) {

	for {
		//for _, ch := range s.subs {

		// listen for messages on different channels
		select {
		case <-ctx.Done(): // listen context cancellation
			return
		case article, ok := <-s.broadcastCh(): // listen for news articles

			// check if the channel is closed or not
			if !ok {
				continue
			}

			s.Publish(fmt.Sprint(article))

		}
		//}
	}
}

// Publish publishes new stories to the subscribers
func (s *Service) Publish(topic string) {

	s.Lock()
	if s.stopped {
		s.Unlock()
		return
	}
	s.Unlock()

	s.Lock()
	t := Subscription(topic)

	select {
	case s.subs[t] <- topic:
		s.SubcribedArticles = append(s.SubcribedArticles, *Article1)
	default:
	}
	s.Unlock()
}

// Broadcast will return a channel that can be listened to
// for new articles to be read.
func (s *Service) broadcastCh() chan interface{} {
	return s.broadcast
}

// Errors will return a channel that can be listened to
// and can be used to receive errors from the new service.
func (s *Service) Errors() chan error {
	return s.errs
}

//The resources include subscribers and news sources
func (s *Service) Stop() {

	s.Lock()
	if s.stopped {
		s.Unlock()
		return
	}
	s.Unlock()

	s.once.Do(func() {

		s.Lock()
		defer s.Unlock()

		s.StateFileLoad()

		s.cancel()

		s.stopped = true

		// close all channels
		for _, ch := range s.subs {
			if ch != nil {
				close(ch)
			}
		}

		close(s.broadcast)

		close(s.errs)

	})
}

// StateFileLoad data from a json file.
func (s *Service) StateFileLoad() error {
	if s.stateFile == "" {
		return nil
	}

	if buf, err := ioutil.ReadFile(s.stateFile); os.IsNotExist(err) {
		return s.StateFileSave()
	} else if err != nil {
		return err
	} else if err := json.Unmarshal(buf, s); err != nil {
		return err
	}

	return nil
}

// StateGetJSON returns the state data in json format.
func (s *Service) StateGetJSON() (string, error) {
	s.RLock()
	defer s.RUnlock()

	b, err := json.Marshal(s)

	return string(b), err
}

// StateFileSave writes out the state file.
func (s *Service) StateFileSave() error {
	if s.stateFile == "" {
		return nil
	}

	s.RLock()
	defer s.RUnlock()

	if buf, err := json.Marshal(s); err != nil {
		return fmt.Errorf("marshaling json: %w", err)
	} else if err = ioutil.WriteFile(s.stateFile, buf, 0o600); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}
