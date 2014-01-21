// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

import (
	"errors"
	"fmt"
)

func p(v ...interface{}) {
	fmt.Println(v...)
}

type IdMap struct {
	L []uint64
}

func (m *IdMap) Id(v Vertex) uint64 {
	return m.L[v]
}

func (m *IdMap) Map() map[uint64]Vertex {
	idMap := make(map[uint64]Vertex)
	for i, id := range m.L {
		idMap[id] = Vertex(i)
	}
	return idMap
}

type VertexStack []Vertex

func (s *VertexStack) Empty() bool {
	return len(*s) == 0
}

func (s *VertexStack) Top() Vertex {
	if s.Empty() {
		panic(errors.New("VertexStack is empty."))
	}
	return (*s)[len(*s)-1]
}

func (s *VertexStack) Push(v Vertex) {
	*s = append(*s, v)
}

func (s *VertexStack) PushMany(indices []Vertex) {
	*s = append(*s, indices...)
}

func (s *VertexStack) Pop() (v Vertex) {
	v, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return v
}

type VertexSlice []Vertex

func (s VertexSlice) Insert(v Vertex, i int) VertexSlice {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = v
	return s
}

func (s VertexSlice) Append(v Vertex) VertexSlice {
	return append(s, v)
}

func (s VertexSlice) Remove(i int) (Vertex, VertexSlice) {
	v := s[i]
	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	return v, s[:len(s)-1]
}

func (s VertexSlice) FastRemove(i int) (Vertex, VertexSlice) {
	v := s[i]
	s[i], s[len(s)-1] = s[len(s)-1], 0
	return v, s[:len(s)-1]
}


type VertexSortedSet struct {
	S VertexSlice
}

func (set *VertexSortedSet) Count() int {
	return len(set.S)
}

func (set *VertexSortedSet) Add(v Vertex) {
	if v == 0 {
		panic("vertex should not be zero.")
	}
	if i, ok := set.BinarySearch(v); !ok {
		set.S = set.S.Insert(v, i)
	}
}

func (set *VertexSortedSet) Remove(i int) Vertex {
	if i < 0 || i >= len(set.S) {
		panic("out of range.")
	}
	v, s := set.S.Remove(i)
	set.S = s
	return v
}

func (set *VertexSortedSet) Has(v Vertex) bool {
	_, ok := set.BinarySearch(v)
	return ok
}

func (set *VertexSortedSet) BinarySearch(k Vertex) (int, bool) {
	min, max := 0, len(set.S)-1
	for min <= max {
		mid := (max + min) / 2
		switch v := set.S[mid]; {
		case v < k:
			min = mid + 1
		case v > k:
			max = mid - 1
		default:
			return mid, true
		}
	}

	return min, false
}

func (set *VertexSortedSet) Slice() VertexSlice {
	return set.S
}

func imin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// -------------

type Component int

func ToCompIndex(c []Component, cc int) []VertexSlice {
	r := make([]VertexSlice, cc)
	for i, comp := range c {
		v := Vertex(i)
		r[comp] = r[comp].Append(v)
	}
	return r
}

func imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
