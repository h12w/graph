// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

import (
	"encoding/gob"
	"io"
)

// A list of adjacent list represented with vertex index
// index 0 is kept unused so that index 0 is invalid.
type adjacencyGraph struct {
	L []VertexSortedSet
}

func NewAdjacencyGraph(n int) Graph {
	return &adjacencyGraph{make([]VertexSortedSet, n+1)}
}

/*
func (g *adjacencyGraph) Edges() EdgeIter {
	return newAgEdgeIter(g)
}
*/

func (g *adjacencyGraph) AdjacentVertices(v Vertex) VertexSlice {
	return g.L[v].Slice()
}

func (g *adjacencyGraph) VertexMax() Vertex {
	return Vertex(len(g.L) - 1)
}

func (g *adjacencyGraph) HasEdge(v, w Vertex) bool {
	if int(v) >= len(g.L) {
		return false
	}
	return g.L[v].Has(w)
}

func (g *adjacencyGraph) AddEdge(v, w Vertex) {
	if maxV := imax(int(v), int(w)); maxV >= len(g.L) {
		g.L = append(g.L, make([]VertexSortedSet, maxV-len(g.L)+1)...)
	}
	g.L[v].Add(w)
}

func (g *adjacencyGraph) Deserialize(r io.Reader) error {
	return gob.NewDecoder(r).Decode(g)
}

func (g *adjacencyGraph) Serialize(w io.Writer) error {
	return gob.NewEncoder(w).Encode(g)
}

/*
type agEdgeIter struct {
	g    *adjacencyGraph
	i, j int
}

func newAgEdgeIter(g *adjacencyGraph) *agEdgeIter {
	return &agEdgeIter{g, 1, -1}
}

func (e *agEdgeIter) Edge() (v, w Vertex) {
	return Vertex(e.i), e.g.L[e.i][e.j]
}

func (e *agEdgeIter) Next() bool {
	if e.i == len(e.g.L) {
		return false
	}
	e.j++
	if e.j == len(e.g.L[e.i]) {
		e.j = 0
		e.i++
	}
	for e.i < len(e.g.L) && len(e.g.L[e.i]) == 0 {
		e.i++
	}
	return e.i != len(e.g.L)
}
*/
