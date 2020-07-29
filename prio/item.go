package prio

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// An Item is something we manage in a priority queue.
type Item struct {
	tags     []string // Contains tags for filtering
	function ItemFunc // The value of the item; arbitrary.
	priority int      // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type ExecutionReturn struct {
	abort bool
	err   error
}

type ItemFunc func(object metav1.Object, config ExecutionConfig) ExecutionReturn
