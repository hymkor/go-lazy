package lazy_test

import (
	"testing"

	"github.com/hymkor/go-lazy"
)

func TestLazyOfString(t *testing.T) {
	const sample = "x"

	called := 0

	data := lazy.Of[string]{
		New: func() string { called++; return sample },
	}

	if called != 0 {
		t.Fatal("failed 1")
	}

	if data.Value() != sample {
		t.Fatal("failed 2")
	}

	if called != 1 {
		t.Fatal("failed 3")
	}

	if data.Value() != sample {
		t.Fatal("failed 4")
	}

	if called != 1 {
		t.Fatal("failed 5")
	}

	data2 := lazy.New(func() string { return sample })
	if data2.Value() != sample {
		t.Fatal("failed 6")
	}
	// println( data2.Value())
}
