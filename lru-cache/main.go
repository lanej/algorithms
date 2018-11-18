package main

import "fmt"

// LRUCache ...
type LRUCache struct {
	cache map[int]*entry
	queue queue
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache: make(map[int]*entry, capacity),
		queue: queue{capacity: capacity},
	}
}

// Get ...
func (lru *LRUCache) Get(key int) int {
	fmt.Printf("get(%d) ; queue(%v)\n", key, lru.queue)
	if e, exists := lru.cache[key]; exists {
		lru.queue.delete(e)
		lru.queue.push(e)

		return e.value
	}

	return -1
}

// Put ...
func (lru *LRUCache) Put(key int, value int) {
	fmt.Printf("put(%d) ; queue(%v)\n", key, lru.queue)

	if e, exists := lru.cache[key]; exists {
		lru.queue.delete(e)
	}

	mru := entry{key: key, value: value}

	if evicted := lru.queue.push(&mru); evicted != nil {
		delete(lru.cache, evicted.key)
	}

	lru.cache[key] = &mru
}

type queue struct {
	head     *entry
	tail     *entry
	capacity int
	size     int
}

func (q queue) String() (inspect string) {
	inspect += fmt.Sprintf("s=%d,c=%d", q.size, q.capacity)

	if q.head == nil {
		inspect += " []"
		return
	}

	inspect += fmt.Sprintf("h(%d),t(%d)", q.head.key, q.tail.key)

	inspect += "--"

	e := q.head

	for e != nil {
		inspect += fmt.Sprintf("-->{%d,%d}", e.key, e.value)
		e = e.next
	}

	return
}

func (q *queue) delete(e *entry) {
	if q.tail == e {
		q.tail = e.prev
	} else {
		e.next.prev = e.prev
	}

	if q.head == e {
		q.head = e.next
	} else {
		e.prev.next = e.next
	}

	q.size--
}

func (q *queue) push(e *entry) (evicted *entry) {
	if q.size >= q.capacity {
		evicted = q.tail
		q.delete(q.tail)
	}

	q.size++

	if q.tail == nil {
		q.tail = e
	} else {
		q.head.prev = e
		e.next = q.head
	}

	q.head = e

	return
}

type entry struct {
	key   int
	value int
	prev  *entry
	next  *entry
}
