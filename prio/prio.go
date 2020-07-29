package prio

import (
	"container/heap"
)

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) add(item *Item) {
	heap.Push(pq, item)
}

func (pq *PriorityQueue) get() *Item {
	return heap.Pop(pq).(*Item)
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, function ItemFunc, priority int) {
	item.function = function
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq PriorityQueue) filterByTags(tags []string) *PriorityQueue {

	newPq := PriorityQueue{}

	for i := range pq {
		if sliceContainsStrings(pq[i].tags, tags) {
			newPq = append(newPq, pq[i])
		}
	}
	heap.Init(&newPq)
	return &newPq
}

func newPriorityQueue() *PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

// Checks if the slice of strings contains a specific string
func sliceContainsStrings(list []string, compare []string) bool {
	for _, s := range compare {
		for _, v := range list {
			if v == s {
				return true
			}
		}
	}
	return false
}
