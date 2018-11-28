package main

type heap struct {
	data []int
}

func (h *heap) Push(e int) {
	h.data = append(h.data, e)

	h.up(len(h.data) - 1)
}

func (h *heap) Pop() int {
	if len(h.data) < 1 {
		return -1
	}

	end := len(h.data) - 1

	h.data[0], h.data[end] = h.data[end], h.data[0]
	h.down(0, end)

	root := h.data[end]
	h.data = h.data[:end]

	return root
}

func (h *heap) Peek() int {
	return h.data[0]
}

func (h *heap) down(i0, n int) bool {
	i := i0

	for {
		// find child
		j1 := 2*i + 1

		// past the end or j1 < 0 after int overflow
		if j1 >= n || j1 < 0 {
			break
		}

		// assume left child
		j := j1

		// set right child if larger
		// j2 = 2*i + 2
		if j2 := j1 + 1; j2 < n && h.data[j2] < h.data[j1] {
			j = j2
		}

		// target is greater than both children, so we're done
		if h.data[j] >= h.data[i] {
			break
		}

		// swap parent with child
		h.data[i], h.data[j] = h.data[j], h.data[i]

		// mark new child position
		i = j
	}

	// if any swaps happened, this will return true
	return i > i0
}

func (h *heap) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || h.data[j] > h.data[i] {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		j = i
	}
}

func heapify(n []int) *heap {
	h := heap{data: []int{}}

	for _, e := range n {
		h.Push(e)
	}

	return &h
}
