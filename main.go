package lazy

import (
	"sync"
)

type Of[T any] struct {
	New   func() T
	once  sync.Once
	value T
}

func (this *Of[T]) Value() T {
	if this.New != nil {
		this.once.Do(func() {
			this.value = this.New()
			this.New = nil
		})
	}
	return this.value
}

func New[T any](newfunc func() T) *Of[T] {
	return &Of[T]{New: newfunc}
}
