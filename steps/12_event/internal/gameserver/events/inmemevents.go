package events

import (
	"sync"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

var _ Eventer = &InMemoryEvents{}

// InMemoryEvents todo: needs to be scalable, use redis instead
type InMemoryEvents struct {
	mu        sync.RWMutex
	listeners map[EventType][]Listener
}

func NewInMemoryEvents() *InMemoryEvents {
	return &InMemoryEvents{
		listeners: make(map[EventType][]Listener),
	}
}

func (e *InMemoryEvents) Dispatch(t EventType, info EventInfo) error {
	info.Type = t
	e.mu.RLock()
	listeners := append(e.listeners[t], e.listeners[EventAny]...)
	for _, listener := range listeners {
		go listener.callback(info)
	}
	e.mu.RUnlock()
	return nil
}

func (e *InMemoryEvents) listenerCount(t EventType) int {
	e.mu.RLock()
	defer e.mu.RUnlock()

	return len(e.listeners[t])
}

func (e *InMemoryEvents) Clean(t EventType) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.listeners[t] = make([]Listener, 0)
	return nil
}

func (e *InMemoryEvents) Close() error {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.listeners = make(map[EventType][]Listener)
	return nil
}

func (e *InMemoryEvents) Register(t EventType, callback Callback) (func(), error) {
	e.mu.Lock()

	uid := uuid.New().String()
	e.listeners[t] = append(e.listeners[t], Listener{
		callback: callback,
		uuid:     uid,
	})

	e.mu.Unlock()

	return func() {
		e.mu.Lock()
		e.listeners[t] = lo.Filter(e.listeners[t], func(item Listener, index int) bool {
			return uid != item.uuid
		})
		e.mu.Unlock()
	}, nil
}
