package newsmill

import (
	"context"
	"io"
	"sync"
	"time"
)

const (
	File SourceType = "file"
	Mock SourceType = "mock"
	URL  SourceType = "url"
)

type SourceType string

//The Source to publish stories for any category, or categories, they wish to define
type Source struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	SourceType `json:"publishertype"`
	sync.RWMutex
	cancel context.CancelFunc
	//stopped bool
}

//News story/article that is published by the Publisher
type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Content     string    `json:"body"`
	PublishedAt time.Time `json:"publishedAt"`
	Category    []string  `json:"category"`
}

type Sources interface {
	Publish(ctx context.Context, cat []string, r io.Reader) ([]*Article, error)
	Stop()
}

func (s *Source) Publish(ctx context.Context, cat []string, r io.Reader) ([]Article, error) {

	return nil, nil
}

func (s *Source) Stop() {
	s.RLock()
	defer s.RUnlock()

	if s.cancel != nil {
		s.cancel()
	}
}
