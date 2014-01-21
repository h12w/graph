// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

import (
	"math/rand"
)

type VertexSequencer struct {
	m map[Vertex]Vertex
}

func NewVertexSequencer() *VertexSequencer {
	return &VertexSequencer{make(map[Vertex]Vertex)}
}

func (s *VertexSequencer) Has(v Vertex) bool {
	_, ok := s.m[v]
	return ok
}

func (s *VertexSequencer) Get(v Vertex) Vertex {
	if vv, ok := s.m[v]; ok {
		return vv
	}
	vv := Vertex(len(s.m) + 1)
	s.m[v] = vv
	return vv
}

type VertexSet struct {
	m map[Vertex]int // Vertex -> index in the array
	a VertexSlice
}

func NewVertexSet() *VertexSet {
	return &VertexSet{make(map[Vertex]int), VertexSlice{}}
}

func (s *VertexSet) Add(v Vertex) {
	if _, ok := s.m[v]; !ok {
		s.a = append(s.a, v)
		s.m[v] = len(s.a) - 1
	}
}

func (s *VertexSet) Remove(v Vertex) {
	if i, ok := s.m[v]; ok {
		s.m[s.a[len(s.a)-1]] = i
		_, s.a = s.a.FastRemove(i)
		delete(s.m, v)
	}
}

func (s *VertexSet) Has(v Vertex) bool {
	_, ok := s.m[v]
	return ok
}

func (s *VertexSet) Sample() Vertex {
	return s.a[rand.Intn(len(s.a))]
}

func (s *VertexSet) Count() int {
	return len(s.a)
}
