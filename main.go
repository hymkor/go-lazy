package lazy

import (
	"github.com/cheekybits/genny/generic"
)

type Item = generic.Type

type LazyItem struct {
	initializer func() Item
	value       Item
}

func (this *LazyItem) Value() Item {
	if this.initializer != nil {
		this.value = this.initializer()
		this.initializer = nil
	}
	return this.value
}

func (this *LazyItem) IsCreated() bool {
	return this.initializer != nil
}

func OfItem(f func() Item) LazyItem {
	return LazyItem{initializer: f}
}
