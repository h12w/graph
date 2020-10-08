// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

import (
	"h12.io/go-gsl/stats"
)

// Assortativity Coefficient of graph g from vertex v to w
// cv[v]: scalar characteristic of node v
// cw[w]: scalar characteristic of node w
func AssortativityCoefficient(g Graph, cv, cw []float64) float64 {
	m := EdgeCount(g)
	x, y := make([]float64, m), make([]float64, m)
	i := 0
	for v := Vertex(1); v <= g.VertexMax(); v++ {
		for _, w := range g.AdjacentVertices(v) {
			x[i], y[i] = cv[v], cw[w]
			i++
		}
	}
	return stats.Correlation(x, 1, y, 1, len(x))
}


func AssortativityCoefficient2(g Graph, cv, cw func(Vertex) float64) float64 {
	m := EdgeCount(g)
	x, y := make([]float64, m), make([]float64, m)
	i := 0
	for v := Vertex(1); v <= g.VertexMax(); v++ {
		for _, w := range g.AdjacentVertices(v) {
			x[i], y[i] = cv(v), cw(w)
			i++
		}
	}
	return stats.Correlation(x, 1, y, 1, len(x))
}

