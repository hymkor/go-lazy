package lazy

import (
	"github.com/cheekybits/genny/generic"
)

type Item = generic.Type

type OfItem struct {
	New   func() Item
	value Item
}

func (this *OfItem) Value() Item {
	if this.New != nil {
		this.value = this.New()
		this.New = nil
	}
	return this.value
}
