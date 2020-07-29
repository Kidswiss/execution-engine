package prio

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	// pq holds all steps
	pq = newPriorityQueue()
)

type ExecutionConfig struct {
	Annotation    string
	FinalizerName string
	Tenant        string
}

func Register(function ItemFunc, prio int, tags []string) {
	pq.add(&Item{
		priority: prio,
		function: function,
		tags:     tags,
	})
}

func Execute(object metav1.Object, config ExecutionConfig, tags ...string) error {
	taskList := pq.filterByTags(tags)
	for taskList.Len() > 0 {
		result := taskList.get().function(object, config)
		if result.err != nil || result.abort {
			return result.err
		}
	}
	return nil
}
