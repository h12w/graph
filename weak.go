// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

import (
	"sort"

	"h12.io/go-nauty"
)

type WeakCompFinder struct {
	g  Graph       // Graph
	c  []Component // component numbers of each vertex
	cc int         // Max component sequence number
}

func NewWeakCompFinder(g Graph) *WeakCompFinder {
	ug := ToUndirected(g)
	f := &WeakCompFinder{
		g:  ug,
		c:  make([]Component, ug.VertexMax()+1),
		cc: 1}
	f.run()
	return f
}

// The component numbers of each vertices, starting from index 1
func (f *WeakCompFinder) Result() ([]Component, int) {
	return f.c, f.cc - 1
}

// Run the algorithm
func (f *WeakCompFinder) run() {
	for v := Vertex(1); v <= f.g.VertexMax(); v++ {
		if !f.assigned(v) {
			f.dfs(v)
			f.cc++
		}
	}
}

func (f *WeakCompFinder) dfs(v Vertex) {
	if !f.assigned(v) {
		f.c[v] = Component(f.cc)
		for _, w := range f.g.AdjacentVertices(v) {
			f.dfs(w)
		}
	}
}

// Whether or not a vertex has been assigned a component number
func (f *WeakCompFinder) assigned(v Vertex) bool {
	return f.c[v] != 0
}

func TraverseWeakComp(g Graph, visit func(int, VertexSlice)) {

	finder := NewWeakCompFinder(g)
	c, cc := finder.Result()

	comp := make([][]Vertex, cc+1)
	for i, weakId := range c {
		comp[weakId] = append(comp[weakId], Vertex(i))
	}

	for id, vs := range comp {
		visit(id, vs)
	}
}

type TopoCounter struct {
	G     *nauty.DenseGraph
	Count int
}

type TopoCounterSlice []TopoCounter

func (s TopoCounterSlice) Len() int {
	return len(s)
}
func (s TopoCounterSlice) Less(i, j int) bool {
	if s[i].Count < s[j].Count {
		return true
	} else if s[i].Count == s[j].Count {
		return s[i].G.N > s[j].G.N
	}
	return false
}
func (s TopoCounterSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type reverse struct {
	// This embedded Interface permits reverse to use the methods of
	// another Interface implementation.
	sort.Interface
}

// Less returns the opposite of the embedded implementation's Less method.
func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

func WeakCompToDenseGraph(g Graph, vs VertexSlice) *nauty.DenseGraph {
	dg := nauty.NewDenseGraph(len(vs))
	for i, v := range vs {
		for j, w := range vs {
			if g.HasEdge(v, w) {
				dg.AddEdge(i, j)
			}
		}
	}
	return dg
}

func AnalyzeWeakCompTopo(g Graph, topoMaxSize int) TopoCounterSlice {
	tcnt := make(map[string]TopoCounter)
	maxCompSize := 0
	TraverseWeakComp(g, func(weakId int, vs VertexSlice) {
		if len(vs) > maxCompSize {
			maxCompSize = len(vs)
		}
		if len(vs) > 1 && len(vs) < topoMaxSize {
			dg := WeakCompToDenseGraph(g, vs).ToCanonical()
			if _, ok := tcnt[dg.String()]; ok {
				c := tcnt[dg.String()]
				c.Count++
				tcnt[dg.String()] = c
			} else {
				tcnt[dg.String()] = TopoCounter{G: dg, Count: 1}
			}
		}
	})

	counters := make(TopoCounterSlice, 0, len(tcnt))
	for _, c := range tcnt {
		counters = append(counters, c)
	}
	sort.Sort(reverse{counters})
	return counters
}
