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

type Itemizer interface {
	Priority() int
}

type PriorityQueue(type T Itemizer) []T

func (pq *PriorityQueue(T)) Init() {
	Init(T)(pq)
}

func (pq *PriorityQueue(T)) PushItem(x T) {
	Push(T)(pq, x)
}

func (pq *PriorityQueue(T)) PopItem() T {
		return Pop(T)(pq)
}

func (pq PriorityQueue(T)) Len() int { return len(pq) }

func (pq PriorityQueue(T)) Less(i, j int) bool {
	return pq[i].Priority() < pq[j].Priority()
}

func (pq PriorityQueue(T)) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue(T)) Push(x T) {
	item := x
	*pq = append(*pq, item)
}

func (pq *PriorityQueue(T)) Pop() T {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Fruit struct {
	Name string
	Val int
}

func(item Fruit) Priority() int {
	return item.Val
}

func main() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue(Fruit), len(items))
	i := 0
	for value, priority := range items {
		pq[i] = Fruit{
			Name:    value,
			Val: priority,
		}
		i++
	}
	pq.Init()

	item := Fruit{
		Name:    "orange",
		Val: 1,
	}
	pq.PushItem(item)

	for pq.Len() > 0 {
		item := pq.PopItem()
		fmt.Printf("%.2d:%s ", item.Val, item.Name)
	}
}
