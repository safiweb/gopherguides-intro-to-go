package newsmill

import (
	"context"
	"sync"
)

type NewsService struct {
	sync.RWMutex
	cancel      context.CancelFunc
	stopped     bool
	Subscribers map[string][]*Subscriber
	//sources *Source
}

//Start
func (ns *NewsService) Start(ctx context.Context) (context.Context, error) {

	ctx, cancel := context.WithCancel(ctx)

	ns.Lock()
	ns.cancel = cancel
	ns.Unlock()

	return ctx, nil
}

//Subscribe will subscribe a subscriber to the news service
func (ns *NewsService) Subscribe(ctx context.Context) error {
	return nil
}

//UnSubscribe will unsubscribe a subscriber from the news service
func (ns *NewsService) Unsubscribe(ctx context.Context) error {
	return nil
}

/*
//Read
func (ns *NewsService) Read() {}

//Stream
func (ns *NewsService) Stream() {}

//Clear
func (ns *NewsService) Clear() {}
*/

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
