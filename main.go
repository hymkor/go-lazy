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

type Two[T any, E any] struct {
	New    func() (T, E)
	once   sync.Once
	value1 T
	value2 E
}

func (this *Two[T, E]) Values() (T, E) {
	if this.New != nil {
		this.once.Do(func() {
			this.value1, this.value2 = this.New()
			this.New = nil
		})
	}
	return this.value1, this.value2
}
