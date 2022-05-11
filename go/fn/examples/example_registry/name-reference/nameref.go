package nameref

import (
	"fmt"
	"strings"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	corev1 "k8s.io/api/core/v1"
)

var _ fn.FunctionRunner = &NameReference{}


type NameReference struct {
	NameRef []SourceToTarget `json:"nameRef,omitempty"`
}

type SourceToTarget struct {
	source fn.Identifier `json:"source,omitempty"`
	targets []fn.FieldSpec `json:"targets,omitempty"`
}

func (r *NameReference) RequireTrackOrigin() bool {
	return true
}

func (r *NameReference) SetDefaultConfigMap(_ corev1.ConfigMap) bool {
	return false
}

func (r *NameReference) SetFunctionConfig(o *fn.KubeObject) bool {
	o.As(r)
	return true
}

func (r *NameReference) Run(ctx *fn.Context, items []*fn.KubeObject) {
 	for _, nameRef := range r.NameRef {
 		sObject := ctx.ResMap[nameRef.source][0]
		for _, target := range nameRef.targets {
			tObjects := ctx.ResMap[target.Identifier]
			for _, tObject := range tObjects {
				for _, path := range target.Path {
					fieldspec := strings.Split(path, ".")
					_, found, _ := tObject.NestedString(fieldspec...)
					if found {
						tObject.SetNestedField(sObject.GetName() ,fieldspec...)
						ctx.AddGeneralResult(fmt.Sprintf("update target %v path %v to %v, source %v", tObject, path, sObject.GetName(),sObject), fn.Info)
					}
				}
			}
		}
	}
}
