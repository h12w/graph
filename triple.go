// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

import (
	"fmt"
)

type EdgeDirection byte

const (
	None EdgeDirection = iota
	Forward
	Backward
	Both
)

func (e EdgeDirection) Reverse() EdgeDirection {
	switch e {
	case Forward:
		return Backward
	case Backward:
		return Forward
	}
	return e
}

type Triple struct {
	V byte
}

func NewTriple(a, b, c EdgeDirection) Triple {
	return Triple{byte(a<<4 | b<<2 | c)}
}

func (t Triple) Edges() (a, b, c EdgeDirection) {
	return EdgeDirection((t.V >> 4) & 0x3),
		EdgeDirection((t.V >> 2) & 0x3),
		EdgeDirection(t.V & 0x3)
}

func (t Triple) String() string {
	a, b, c := t.Edges()
	return fmt.Sprintf("[%d %d %d]", a, b, c)
}

func (t Triple) Reverse() Triple {
	a, b, c := t.Edges()
	return NewTriple(a.Reverse(), c.Reverse(), b.Reverse())
}

func (t Triple) Shift() Triple {
	a, b, c := t.Edges()
	return NewTriple(c, a, b)
}

func (t Triple) Min() Triple {
	ts := t.Shift()
	tss := ts.Shift()
	return tmin(t, ts, tss)
}

func (t Triple) Less(o Triple) bool {
	return t.V < o.V
}

func (t Triple) ToCanonical() Triple {
	return tmin(t.Min(), t.Reverse().Min())
}

type TripleSlice []Triple

func (s TripleSlice) Len() int {
	return len(s)
}
func (s TripleSlice) Less(i, j int) bool {
	return s[i].V < s[j].V
}
func (s TripleSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func tmin(ts ...Triple) Triple {
	min := ts[0]
	for _, t := range ts[1:] {
		if t.Less(min) {
			min = t
		}
	}
	return min
}
