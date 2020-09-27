package normal

import (
	"container/heap"
)

type Node struct {
	i, v int
}

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}
func (h NodeHeap) Less(i, j int) bool {
	return h[i].v<h[j].v
}
func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}
func (h *NodeHeap) Pop() interface{} {
	n := len(*h)
	v := (*h)[n-1]
	*h = (*h)[:n-1]
	return v
}
func (h NodeHeap) peek() Node {
	return h[0]
}

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	h := make(NodeHeap, 0)
	heap.Init(&h)

	for key, v := range m {
		if h.Len() < k {
			heap.Push(&h, Node{key, v})
		} else if h.peek().v < v {
			heap.Pop(&h)
			heap.Push(&h, Node{key, v})
		}
	}
	res := make([]int, 0)
	for h.Len() > 0 {
		res = append(res, h.Pop().(Node).i)
	}
	return res
}

func Hello() string {

	return "hello world!"
}

func P347(nums []int, k int) []int {

	return topKFrequent(nums, k)
}
