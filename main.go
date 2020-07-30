package main

import (
	"fmt"

	_ "github.com/Kidswiss/execution-engine/logic"
	"github.com/Kidswiss/execution-engine/prio"
	synv1alpha1 "github.com/projectsyn/lieutenant-operator/pkg/apis/syn/v1alpha1"
)

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {

	data := prio.ExecutionData{
		Tenant: "test",
	}

	fmt.Println("First run")
	prio.Execute(&synv1alpha1.GitRepo{}, data, "common")

	fmt.Println("Second run")
	prio.Execute(&synv1alpha1.GitRepo{}, data, "common", "wasting")

	fmt.Println("Third run")
	prio.Execute(&synv1alpha1.GitRepo{}, data, "wasting")
}
