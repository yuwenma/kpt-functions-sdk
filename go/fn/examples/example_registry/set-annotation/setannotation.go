package setannotation

import (
	"fmt"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	corev1 "k8s.io/api/core/v1"
)

var _ fn.FunctionRunner = &SetAnnotations{}

type SetAnnotations struct {
	Annotations map[string]string `json:"spec,omitempty"`
}

func (r *SetAnnotations) Run(ctx *fn.Context, items []*fn.KubeObject) {
	for _, o := range items {
		for k, newLabel := range r.Annotations {
			oldLabel := o.GetAnnotation(k)
			if oldLabel != newLabel {
				o.SetLabel(k, newLabel)}
		}
		ctx.AddGeneralResult(fmt.Sprintf("update %v labels %v", o, r.Annotations), fn.Info)
	}
}
