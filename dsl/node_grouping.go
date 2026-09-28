package dsl

import "reflect"

// internal grouping structures
type procBucket struct {
	seen  map[uintptr]struct{}
	order []CSS
}

func newProcBucket() *procBucket {
	return &procBucket{
		seen:  make(map[uintptr]struct{}),
		order: make([]CSS, 0),
	}
}

func (b *procBucket) Add(fn CSS) {
	ptr := reflect.ValueOf(fn).Pointer() // general CCS generation is cold path.

	if _, exists := b.seen[ptr]; exists {
		return
	}

	b.seen[ptr] = struct{}{}
	b.order = append(b.order, fn)
}

type group struct {
	Declarative map[string]string // prop -> value
	Procedural  *procBucket
}

func newGroup() *group {
	return &group{
		Declarative: make(map[string]string),
		Procedural:  newProcBucket(),
	}
}
