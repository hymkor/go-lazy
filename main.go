package lazy

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item = generic.Type

type OfItem struct {
	New      func() Item
	lockItem sync.Mutex
	value    Item
}

func (this *OfItem) Value() Item {
	this.lockItem.Lock()
	defer this.lockItem.Unlock()
	if this.New != nil {
		this.value = this.New()
		this.New = nil
	}
	return this.value
}
