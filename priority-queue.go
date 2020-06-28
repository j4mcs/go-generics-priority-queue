package main

import (
	"fmt"
	"sort"
)

///HEAP

type Interface(type T) interface {
	sort.Interface
	Push(T)
	Pop() T
}

func Init(type T)(h Interface(T)) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func Push(type T)(h Interface(T), x T) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop(type T)(h Interface(T)) T {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func Remove(type T)(h Interface(T), i int) T {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

func Fix(type T)(h Interface(T), i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func up(type T)(h Interface(T), j int) {
	for {
		i := (j - 1) / 2
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down(type T)(h Interface(T), i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

///PQ

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x *Item) {
	item := x
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() *Item {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func main() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
		}
		i++
	}
	Init(*Item)(&pq)

	item := &Item{
		value:    "orange",
		priority: 1,
	}
	Push(*Item)(&pq, item)

	for pq.Len() > 0 {
		item := Pop(*Item)(&pq)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}

