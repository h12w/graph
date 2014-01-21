// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

/*
Gabow's path-based strong component algorithm

Gabow's algorithm is basicly Tarjan's algorithm by using a stack to track root vertices instead of calculating low point values.

The crux of the algorithm comes in determining whether a node is the root of a strongly connected component. The root node is simply the first node of the strongly connected component which is encountered during the depth-first traversal. When the recursion on a node's successors has finished (postorder), if the node is identified as the root node, then all nodes on the stack from the root upwards (to the top of the stack) form a complete strongly connected component.

Ref:
1. http://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm
2. http://en.wikipedia.org/wiki/Path-based_strong_component_algorithm
*/
type StrongCompFinder struct {
	g  Graph // Graph
	s  VertexStack     // Stack to keep track of visited but not assigned vertices
	r  VertexStack     // Stack to keep track of root vertices
	i  []int           // Preorder numbers of each vertex
	ic int             // Max preorder sequence number
	c  []Component     // component numbers of each vertex
	cc int             // Max component sequence number
}

func NewStrongCompFinder(g Graph) *StrongCompFinder {
	f := &StrongCompFinder{
		g:  g,
		s:  VertexStack{},
		r:  VertexStack{},
		i:  make([]int, g.VertexMax()+1),
		ic: 1,
		c:  make([]Component, g.VertexMax()+1),
		cc: 1}
	f.run()
	return f
}

// The component numbers of each vertices, starting from index 1
func (f *StrongCompFinder) Result() ([]Component, int) {
	return f.c, f.cc
}

// Run the algorithm
func (f *StrongCompFinder) run() {
	for v := Vertex(1); v <= f.g.VertexMax(); v++ {
		f.dfs(v)
	}
}

// High level algorithm formed by recursive depth first search
func (f *StrongCompFinder) dfs(v Vertex) {
	if !f.visited(v) {
		// Push in preorder of DFS
		f.pushVertex(v)
		for _, w := range f.g.AdjacentVertices(v) {
			f.dfs(w)
		}
		// Pop in postorder of DFS
		if f.isRoot(v) {
			f.popComponent(v)
		}
	} else if !f.assigned(v) {
		// contract when a vertex is visisted but not assigned
		f.contract(v)
	}
}

// Whether or not a vertex has been visited
func (f *StrongCompFinder) visited(v Vertex) bool {
	return f.i[v] != 0
}

// Whether or not a vertex has been assigned a component number
func (f *StrongCompFinder) assigned(v Vertex) bool {
	return f.c[v] != 0
}

// Whether or not a vertex is a root vertex on the top of stack r
func (f *StrongCompFinder) isRoot(v Vertex) bool {
	return v == f.r.Top()
}

// Contract stack r to contain only the root vertices
func (f *StrongCompFinder) contract(v Vertex) {
	if !f.r.Empty() {
		for f.i[f.r.Top()] > f.i[v] {
			f.r.Pop()
		}
	}
}

// Push vertex to the stack, and assign a sequence number to it
func (f *StrongCompFinder) pushVertex(v Vertex) {
	f.s.Push(v)
	f.r.Push(v)
	f.i[v] = f.ic
	f.ic++
}

// Pop a component from top of the stack to the root vertex
func (f *StrongCompFinder) popComponent(root Vertex) {
	for {
		v := f.s.Pop()
		f.c[v] = Component(f.cc)
		if v == root {
			break
		}
	}
	f.cc++
	// Pop the root vertex after the component is done.
	f.r.Pop()
}
