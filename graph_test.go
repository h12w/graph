// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

import (
	"reflect"
	"sort"
	"testing"
)

func Test_VertexSlice(t *testing.T) {
	s := VertexSlice{}
	s = s.Append(1)
	if !reflect.DeepEqual(s, VertexSlice{1}) {
		t.Fail()
	}
}

func Test_triple(t *testing.T) {
	tcnt := make(map[Triple]int)
	for i := byte(0); i < 64; i++ {
		t := Triple{i}
		tcnt[t.ToCanonical()]++
		if t != t.ToCanonical() {
			//p(t, t.ToCanonical())
		}
	}
	s := make(TripleSlice, 0, len(tcnt))
	cs := make([]int, 0, len(tcnt))
	for t, c := range tcnt {
		s = append(s, t)
		cs = append(cs, c)
	}
	sort.Sort(s)
	/*
		for i := range s {
			p(s[i], cs[i])
		}
		p(len(s))
	*/
}

func Test_strong(t *testing.T) {
	graph := sample1()
	finder := NewStrongCompFinder(graph)

	c, cc := finder.Result()
	r := ToCompIndex(c, cc)
	if !reflect.DeepEqual(r, []VertexSlice{{0},
		{2},
		{10, 11, 12, 13},
		{1, 3, 4, 5, 6, 7},
		{8, 9},
	}) {
		t.Fail()
	}
}

func binarySearch(t *testing.T, s VertexSortedSet, v Vertex, index int, found bool) {
	if i, f := s.BinarySearch(v); i != index || f != found {
		t.Fatalf("Search %v in %v, expected %v, got %v", v, s, index, i)
	}
}

func Test_BinarySearch(t *testing.T) {
	binarySearch(t, VertexSortedSet{VertexSlice(nil)}, 0, 0, false)
	binarySearch(t, VertexSortedSet{VertexSlice{}}, 0, 0, false)

	binarySearch(t, VertexSortedSet{VertexSlice{1, 3}}, 1, 0, true)
	binarySearch(t, VertexSortedSet{VertexSlice{1, 3}}, 3, 1, true)

	binarySearch(t, VertexSortedSet{VertexSlice{1, 3}}, 0, 0, false)
	binarySearch(t, VertexSortedSet{VertexSlice{1, 3}}, 2, 1, false)
	binarySearch(t, VertexSortedSet{VertexSlice{1, 3}}, 4, 2, false)
}

func Test_weak(t *testing.T) {
	g := sample2()
	finder := NewWeakCompFinder(g)
	c, cc := finder.Result()
	if !reflect.DeepEqual(c, []Component{0, 1, 1, 1, 1, 1, 1}) || cc != 1 {
		t.Fail()
	}
}

/*
Strong components:
[2]
[10 11 12 13]
[1 3 4 5 6 7]
[8 9]
*/
func sample1() Graph {
	return &adjacencyGraph{
		L: toSortedSets([]VertexSlice{{},
			{2, 6, 7},
			{},
			{1, 4},
			{3, 6},
			{3, 4, 12},
			{5},
			{5, 10},
			{7, 9},
			{8, 10},
			{11, 12},
			{13},
			{13},
			{10},
		},
		)}
}

func sample2() Graph {
	return &adjacencyGraph{
		L: toSortedSets([]VertexSlice{{},
			{2},
			{},
			{5},
			{},
			{4, 6},
			{2},
		},
	)}
}

func toSortedSets(s []VertexSlice) []VertexSortedSet {
	r := make([]VertexSortedSet, len(s))
	for i := range r {
		r[i] = VertexSortedSet{s[i]}
	}
	return r
}
