// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

import (
	"io"
)

// Integer type starting from 1
type Vertex int

type Edge struct {
	V, W Vertex
}

type Graph interface {
	VertexMax() Vertex                     // maximum number of vertices in a graph
	AdjacentVertices(v Vertex) VertexSlice // adjacency vertices of vertex v
	AddEdge(v, w Vertex)                   // add edge v -> w
	HasEdge(v, w Vertex) bool              // test edge v -> w
	Deserialize(r io.Reader) error
	Serialize(w io.Writer) error
}

/*
type EdgeIter interface {
	Next() bool
	Edge() (v, w Vertex)
}
*/
