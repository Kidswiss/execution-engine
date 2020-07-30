package logic

import (
	"fmt"
	"strconv"

	"github.com/Kidswiss/execution-engine/prio"

	"github.com/projectsyn/lieutenant-operator/pkg/apis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	prio.Register(addAnnotation, 12, []string{"common"})
	prio.Register(addDeletionProtection, 11, []string{"common"})
	prio.Register(timeWasting, 1337, []string{"wasting"})
}

const (
	DeleteProtectionAnnotation = "syn.tools/protected-delete"
)

func addAnnotation(obj metav1.Object, data prio.ExecutionData) prio.ExecutionReturn {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[apis.LabelNameTenant] = data.Tenant
	obj.SetAnnotations(annotations)
	fmt.Println(annotations)
	return prio.ExecutionReturn{}
}

func addDeletionProtection(obj metav1.Object, data prio.ExecutionData) prio.ExecutionReturn {
	protected, err := strconv.ParseBool("true")
	if err != nil {
		protected = true
	}

	if protected {
		annotations := obj.GetAnnotations()

		if annotations == nil {
			annotations = make(map[string]string)
		}

		if _, ok := annotations[DeleteProtectionAnnotation]; !ok {
			annotations[DeleteProtectionAnnotation] = "true"
		}

		obj.SetAnnotations(annotations)
		fmt.Println(annotations)
	}
	return prio.ExecutionReturn{}
}

func timeWasting(obj metav1.Object, data prio.ExecutionData) prio.ExecutionReturn {
	fmt.Println("just wasting time")
	return prio.ExecutionReturn{}
}
