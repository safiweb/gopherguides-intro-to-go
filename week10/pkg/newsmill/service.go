package newsmill

import (
	"context"
	"sync"
)

/*
The news service should be able to be stopped and started multiple times.
The news service should be able to be stopped by the end user.
The news service should be able to be stopped by the news service itself.
The news service should load any saved state from the backup file when it is started.
The news service should not be able to be stopped by the news sources.
The news service should not be able to be stopped by the subscribers.
The news service should periodically save the state of the news service to a backup file, in JSON format.
The news service should provide access to historical news stories by ID number, or range of ID numbers.
The news service should save the state of the news service to the backup file, in JSON format, when it is stopped.
The news service should stop all sources and subscribers when it is stopped.
*/

// Service configuration
type Config struct{}

// Service status
type State struct{}

// topics channel
// subscribers channel
// unsubscribe channel
// stop channel

// Service is responsible for receiving new stories and broadcasting to the news subscribers.
type Service struct {
	config      Config
	mu          sync.RWMutex
	subs        map[string][]*Subscription
	subscribe   chan *Subscription
	unsubscribe chan *Subscription
	publish     chan interface{}
	closing     chan bool
	state       State
	errs        chan error
	cancel      context.CancelFunc
	stopped     bool
}

// NewService creates a new service.
func NewService(cfg Config) *Service {
	return &Service{}
}

/*
//Start
func (ns *NewsService) Start(ctx context.Context) (context.Context, error) {

	ctx, cancel := context.WithCancel(ctx)

	ns.Lock()
	ns.cancel = cancel
	ns.Unlock()

	return ctx, nil
}

//Subscribe will subscribe a subscriber to the news service
func (ns *NewsService) Subscribe(ctx context.Context, topic string) (<-chan interface{}, error) {

	ns.Lock()
	defer ns.Unlock()

	ch := make(chan interface{}, 1)
	ns.subs[topic] = append(ns.subs[topic], ch)

	return ch, nil
}

//Unsubscribe will unsubscribe a subscriber from the news service
func (ns *NewsService) Unsubscribe(ctx context.Context, topic string) error {
	return nil
}

//Publish will publish topics to the news service for subscribers to read
func (ns *NewsService) Publish(topic string, msg interface{}) {

	ns.RLock()
	defer ns.RUnlock()

	if ns.stopped {
		return
	}

	for _, ch := range ns.subs[topic] {
		go func(ch chan interface{}) {
			ch <- msg
		}(ch)
	}

}

//Stop will stop the news service and all resources
//The resources include subscribers and news sources
func (ns *NewsService) Stop() {

	ns.Lock()
	if ns.stopped {
		ns.Unlock()
		return
	}
	ns.Unlock()

	ns.cancel()

	ns.stopped = true
}
*/
