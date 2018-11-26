package main

func postOrderingDFS(edges [][2]int) []int {
	graph := map[int][]int{}

	for _, edge := range edges {
		u, v := edge[0], edge[1]

		graph[u] = append(graph[u], v)
	}

	visited := []int{}
	visit(graph, &visited, edges[0][0])
	return visited
}

func visit(graph map[int][]int, visited *[]int, u int) {
	for _, v := range *visited {
		if v == u {
			return
		}
	}

	*visited = append(*visited, u)

	for _, v := range graph[u] {
		visit(graph, visited, v)
	}

	return
}
