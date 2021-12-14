package newsmill

import "context"

type SourceConfig struct {
	articlesFile string
	Name         string
	Frequency    int
}

//Source interface exposes sources use cases.
type Source interface {
	Run(ctx context.Context, cfg SourceConfig) context.Context
	Fetch(articles ...Article) error
	Publish(categories ...Category) error
	Close()
}
