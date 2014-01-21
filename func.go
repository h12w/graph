// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

func Reverse(g Graph) Graph {
	rg := NewAdjacencyGraph(int(g.VertexMax()))
	for v := Vertex(1); v <= g.VertexMax(); v++ {
		for _, w := range g.AdjacentVertices(v) {
			rg.AddEdge(w, v)
		}
	}
	return rg
}

func ToUndirected(g Graph) Graph {
	ug := NewAdjacencyGraph(int(g.VertexMax()))
	for v := Vertex(1); v <= g.VertexMax(); v++ {
		for _, w := range g.AdjacentVertices(v) {
			ug.AddEdge(v, w)
			ug.AddEdge(w, v)
		}
	}
	return ug
}

// remove isolated vertex
func Reduce(g Graph) Graph {
	vmap := make(map[Vertex]Vertex)
	TraverseWeakComp(g, func(id int, vs VertexSlice) {
		if len(vs) > 1 {
			for _, v := range vs {
				vmap[v] = Vertex(len(vmap) + 1)
			}
		}
	})
	rg := NewAdjacencyGraph(0)
	for i := Vertex(1); i <= g.VertexMax(); i++ {
		v := vmap[i]
		for _, j := range g.AdjacentVertices(i) {
			w := vmap[j]
			rg.AddEdge(v, w)
		}
	}
	return rg
}

func DFS(g Graph, visit func(v Vertex)) {
	visited := make([]bool, g.VertexMax()+1)
	for root := Vertex(1); root <= g.VertexMax(); root++ {
		stack := VertexStack{root}
		for !stack.Empty() {
			if v := stack.Pop(); !visited[v] {
				visit(v)
				visited[v] = true
				stack.PushMany(g.AdjacentVertices(v))
			}
		}
	}
}

