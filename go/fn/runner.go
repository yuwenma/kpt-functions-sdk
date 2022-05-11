package fn

import (
	"fmt"
	"reflect"
	"strings"

	corev1 "k8s.io/api/core/v1"
)

type FunctionRunner interface {
	RequireTrackOrigin() bool
	SetDefaultConfigMap(configMap corev1.ConfigMap) bool
    SetFunctionConfig(o *KubeObject) bool
	Run(context *Context, items []*KubeObject)
}

type runnerProcessor struct {
	fnRunner FunctionRunner
}

func (r runnerProcessor) Process(rl *ResourceList) (bool, error) {
	ctx := &Context{results: &rl.Results}
	if r.fnRunner.RequireTrackOrigin() {
		fnName := reflect.ValueOf(r.fnRunner).Elem().Type().Name()
		resMap, foundOrigin := r.buildOriginFromAnnotations(rl.Items)
		if !foundOrigin {
			return false, fmt.Errorf("origin tracking is required by function %v, but not enabled", fnName)
		}
		ctx.ResMap = resMap
	}
	r.config(ctx, rl.FunctionConfig)
	r.fnRunner.Run(ctx, rl.Items)
	return true, nil
}

var getRunnerName = func(fr FunctionRunner) string {
	return reflect.ValueOf(fr).Elem().Type().Name()
}

func (r *runnerProcessor) config(ctx *Context, o *KubeObject) {
	fnName := getRunnerName(r.fnRunner)
	switch true {
    case o.IsEmpty():
		ctx.AddGeneralResult("`FunctionConfig` is not given", Info)
    case o.IsGVK("v1", "ConfigMap"):
		var cm corev1.ConfigMap
		o.AsOrDie(&cm)
		if ok := r.fnRunner.SetDefaultConfigMap(cm); !ok {
			ctx.AddErrResultAndDie(fmt.Sprintf("function %v does not accept `ConfigMap` as functionConfig", fnName), o)
		}
	case o.IsGVK("", fnName):
		o.As(r.fnRunner)
	default:
		ctx.AddErrResultAndDie(fmt.Sprintf("unknown FunctionConfig `%v`, expect %v", o.GetKind(), fnName), o)
	}
}

func (r *runnerProcessor) buildOriginFromAnnotations(items []*KubeObject) (map[Identifier][]*KubeObject, bool) {
	getfirst := func(val string) string {
		elems := strings.Split(val, ",")
		if len(elems) == 0 {
			return ""
		} else {
			return elems[0]
		}
	}
	foundPrevisOriginSet := false
	resMap := map[Identifier][]*KubeObject{}
	for _, o := range items {
		apiVersion := getfirst(o.GetAnnotation(OriginApiVersions))
		kind := getfirst(o.GetAnnotation(OriginKinds))
		name := getfirst(o.GetAnnotation(OriginNames))
		namespace := getfirst(o.GetAnnotation(OriginNamespaces))

		if apiVersion != "" || kind != "" || name != "" || namespace != "" {
			foundPrevisOriginSet = true
			id := Identifier{
				APIVersion: apiVersion,
				Kind:       kind,
				nameType: NameType{
					Name:      name,
					Namespace: namespace,
				},
			}
			resMap[id] = append(resMap[id], o)
		}
	}
	return resMap, foundPrevisOriginSet
}
