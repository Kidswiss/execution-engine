package prio

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// A Step is a step we want to execute in order
type Step struct {
	tags     []string // Contains tags for filtering
	function StepFunc // The value of the item; arbitrary.
	priority int      // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// ExecutionReturn indicates wether the current execution should be aborted and
// if there was an error.
type ExecutionReturn struct {
	abort bool
	err   error
}

// StepFunc is the function that should be executed in a step.
type StepFunc func(object metav1.Object, data ExecutionData) ExecutionReturn
