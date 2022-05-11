package setlabels

import (
	"fmt"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	corev1 "k8s.io/api/core/v1"
)

var _ fn.FunctionRunner = &SetLabels{}

type SetLabels struct {
	Labels map[string]string `json:"labels,omitempty"`
}
func (r *SetLabels) RequireTrackOrigin() bool {
	return false
}

func (r *SetLabels) SetDefaultConfigMap(configMap corev1.ConfigMap) bool {
	r.Labels = configMap.Data
	return true
}

func (r *SetLabels) SetFunctionConfig(o *fn.KubeObject) bool {
	o.As(r)
	return true
}

func (r *SetLabels) Run(ctx *fn.Context, items []*fn.KubeObject) {
	for _, o := range items {
		for k, newLabel := range r.Labels {
			o.SetLabel(k, newLabel)
		}
		ctx.AddGeneralResult(fmt.Sprintf("update %v labels %v", o, r.Labels), fn.Info)
	}
}
