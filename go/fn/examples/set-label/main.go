package main

import (
	"fmt"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	corev1 "k8s.io/api/core/v1"
)

var _ fn.FunctionRunner = &SetLabel{}

type SetLabel struct {
	Labels map[string]string `json:"spec,omitempty"`
}

func (r *SetLabel) Run(fnConfig *fn.KubeObject, items []*fn.KubeObject, results *fn.Results, selector *fn.Identifier) {
	switch true {
	case fnConfig.IsEmpty():
		*results = append(*results, fn.ErrorResult(fmt.Errorf("FunctionConfig is missing, required `ConfigMap` or `SetLabels`")))
		return
	case fnConfig.IsGVK("v1", "ConfigMap"):
		var cm corev1.ConfigMap
		fnConfig.As(&cm)
		r.Labels = cm.Data
	case fnConfig.IsGVK("fn.kpt.dev/v1alpha1", "SetLabels"):
		var sl SetLabel
		fnConfig.As(&sl)
		r.Labels = sl.Labels
	}
	for _, o := range items {
		for k, newLabel := range r.Labels {
			o.SetLabel(k, newLabel)
		}
	}
}

func main(){
	fn.AsMain(&SetLabel{})
}

