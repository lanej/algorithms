package main

import (
	"math"
)

type stack []int

func (q *stack) Push(e int) {
	*q = append([]int{e}, *q...)
}

func (q *stack) Pop() int {
	if q.Empty() {
		return -1
	}

	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *stack) Empty() bool {
	return len(*q) < 1
}

func djikstra(start, end int, edges [][]int) []int {
	graph := map[int][][]int{}

	for _, edge := range edges {
		a, b, w := edge[0], edge[1], edge[2]

		graph[a] = append(graph[a], []int{b, w})
		graph[b] = append(graph[a], []int{a, w})
	}

	dist := map[int]int{}

	for k := range graph {
		dist[k] = math.MaxInt32
	}

	v := map[int]bool{}
	prev := map[int]int{}

	s := stack{}
	s.Push(start)

	for !s.Empty() {
		u := s.Pop()
		v[u] = true

		for _, edge := range graph[u] {
			d, w := edge[0], edge[1]

			if dist[d] > w {
				dist[d] = w
				prev[d] = u
			}

			if !v[d] {
				s.Push(d)
			}
		}
	}

	revPath := []int{end}
	curr := end

	for curr != start {
		if next, ok := prev[curr]; ok {
			revPath = append(revPath, next)
			curr = next
		}
	}

	path := make([]int, len(revPath))

	for i, j := 0, len(path)-1; i < len(path); i, j = i+1, j-1 {
		path[i] = revPath[j]
	}

	return path
}
