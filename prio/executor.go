package prio

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	// pq holds all steps
	pq = newPriorityQueue()
)

type ExecutionData struct {
	Annotation    string
	FinalizerName string
	Tenant        string
}

func Register(function StepFunc, prio int, tags []string) {
	pq.add(&Step{
		priority: prio,
		function: function,
		tags:     tags,
	})
}

func Execute(object metav1.Object, data ExecutionData, tags ...string) error {
	taskList := pq.filterByTags(tags)
	for taskList.Len() > 0 {
		result := taskList.get().function(object, data)
		if result.err != nil || result.abort {
			return result.err
		}
	}
	return nil
}
