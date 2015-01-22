package dfs

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
)

func undirectedDfs(g *graph.UnGraph, v graph.VertexId, fn func(graph.VertexId)) {
	stack := []graph.VertexId{v}
	visited := make(map[graph.VertexId]bool)

	for len(stack) > 0 {
		l := len(stack) - 1
		v = stack[l]
		stack = stack[:l]

		if _, ok := visited[v]; !ok {
			visited[v] = true
			fn(v)
			neighbours := g.GetNeighbours(v).VerticesIter()
			for neighbour := range neighbours {
				stack = append(stack, neighbour)
			}
		}
	}
}

func directedDfs(g *graph.DirGraph, v graph.VertexId, fn func(graph.VertexId)) {
	stack := []graph.VertexId{v}
	visited := make(map[graph.VertexId]bool)

	for len(stack) > 0 {
		l := len(stack) - 1
		v = stack[l]
		stack = stack[:l]

		if _, ok := visited[v]; !ok {
			visited[v] = true
			fn(v)
			neighbours := g.GetSuccessors(v).VerticesIter()
			for neighbour := range neighbours {
				stack = append(stack, neighbour)
			}
		}
	}
}
