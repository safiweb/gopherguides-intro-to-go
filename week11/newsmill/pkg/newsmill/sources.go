package newsmill

import (
	"context"
	"sync"
)

type MockSource struct {
	cancel context.CancelFunc
	sync.RWMutex
}

// Start the mock source
func (m *MockSource) Start(ctx context.Context) context.Context {
	m.RLock()
	ctx, m.cancel = context.WithCancel(ctx)
	m.RUnlock()
	return ctx
}

// Stop the mock source
func (m *MockSource) Stop() {
	m.RLock()
	defer m.RUnlock()

	if m.cancel != nil {
		m.cancel()
	}

}

type FileSource struct {
	cancel context.CancelFunc
	sync.RWMutex
}

// Start the mock source
func (f *FileSource) Start(ctx context.Context) context.Context {
	f.RLock()
	ctx, f.cancel = context.WithCancel(ctx)
	f.RUnlock()
	return ctx
}

// Stop the mock source
func (f *FileSource) Stop() {
	f.RLock()
	defer f.RUnlock()

	if f.cancel != nil {
		f.cancel()
	}

}
