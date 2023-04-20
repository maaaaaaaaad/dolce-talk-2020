package main

import "fmt"

type Graph map[int][]int

func dfs(node int, parent int, graph Graph, values []int, sums []int) {
	sums[node] = values[node-1]
	for _, child := range graph[node] {
		if child == parent {
			continue
		}
		dfs(child, node, graph, values, sums)
		sums[node] += sums[child]
	}
}

func Solution(values []int, edges [][]int, queries [][]int) []int {
	n := len(values)
	graph := make(Graph)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	sums := make([]int, n+1)
	dfs(1, 0, graph, values, sums)
	answer := make([]int, 0, len(queries))
	for _, query := range queries {
		u, w := query[0], query[1]
		if w == -1 {
			answer = append(answer, sums[u])
		} else {
			diff := w - values[u-1]
			values[u-1] = w
			for u != 1 {
				sums[u] += diff
				parent := graph[u][0]
				if parent == query[1] {
					parent = graph[u][1]
				}
				u = parent
			}
			sums[1] += diff
		}
	}
	return answer
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6}
	edges := [][]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}, {3, 6}}
	queries := [][]int{{1, -1}, {2, 4}, {6, 1}, {4, 0}}
	answer := Solution(values, edges, queries)
	fmt.Println(answer)
}
