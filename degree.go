// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

type GraphVectorFunc func(g Graph) []float64

func EdgeCount(g Graph) int {
	t := 0
	for v := Vertex(1); v <= g.VertexMax(); v++ {
		t += len(g.AdjacentVertices(v))
	}
	return t
}

func Outdegree(g Graph) []float64 {
	ks := make([]float64, g.VertexMax()+1)
	for v := Vertex(1); v <= g.VertexMax(); v++ {
		ks[v] = float64(len(g.AdjacentVertices(v)))
	}
	return ks
}

func Indegree(g Graph) []float64 {
	return Outdegree(Reverse(g))
}

func OutdegreeCount(g Graph) []float64 {
	counter := newDegreeCounter(int(g.VertexMax() + 1))
	for v := Vertex(1); v <= g.VertexMax(); v++ {
		vs := g.AdjacentVertices(v)
		k := len(vs)
		counter.Add(k)
	}
	return counter.ToCount()
}

func IndegreeCount(g Graph) []float64 {
	return OutdegreeCount(Reverse(g))
}

func IndegreePMF(g Graph) []float64 {
	return OutdegreePMF(Reverse(g))
}

func OutdegreePMF(g Graph) []float64 {
	counter := newDegreeCounter(int(g.VertexMax() + 1))
	for v := Vertex(1); v <= g.VertexMax(); v++ {
		vs := g.AdjacentVertices(v)
		k := len(vs)
		counter.Add(k)
	}
	return counter.ToPercent()
}

type DegreeCounter struct {
	C []int
}

func newDegreeCounter(n int) *DegreeCounter {
	return &DegreeCounter{make([]int, n)}
}

func (c *DegreeCounter) Add(k int) {
	c.C[k]++
}

func (c *DegreeCounter) ToPercent() []float64 {
	p := make([]float64, len(c.C))
	total := float64(itotal(c.C))
	for i := range p {
		p[i] = float64(c.C[i]) / total
	}
	return p
}

func (c *DegreeCounter) ToCount() []float64 {
	p := make([]float64, len(c.C))
	for i := range p {
		p[i] = float64(c.C[i])
	}
	return p
}

func itotal(a []int) int {
	t := 0
	for _, i := range a {
		t += i
	}
	return t
}
