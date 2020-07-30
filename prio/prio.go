package prio

import (
	"container/heap"
)

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Step

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
	step := x.(*Step)
	step.index = n
	*pq = append(*pq, step)
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

func (pq *PriorityQueue) add(step *Step) {
	heap.Push(pq, step)
}

func (pq *PriorityQueue) get() *Step {
	return heap.Pop(pq).(*Step)
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(step *Step, function StepFunc, priority int) {
	step.function = function
	step.priority = priority
	heap.Fix(pq, step.index)
}

// filterByTags returns a new queue only containing the steps with the tags.
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

// newPriorityQueue returns a new, empty queue.
func newPriorityQueue() *PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

// sliceContainsStrings checks if the slice of strings contains a specific string
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
