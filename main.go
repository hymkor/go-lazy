package lazy

import (
	"sync"
)

type Of[T any] struct {
	New      func() T
	lockItem sync.Mutex
	value    T
}

func (this *Of[T]) Value() T {
	this.lockItem.Lock()
	defer this.lockItem.Unlock()
	if this.New != nil {
		this.value = this.New()
		this.New = nil
	}
	return this.value
}

func New[T any](newfunc func() T) *Of[T] {
	return &Of[T]{New: newfunc}
}
