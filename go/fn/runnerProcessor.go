// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fn

import (
	"fmt"
	"reflect"
	"strings"
)

type runnerProcessor struct {
	fnRunner Runner
}

func (r runnerProcessor) Process(rl *ResourceList) (bool, error) {
	ctx := &Context{results: &rl.Results}
	r.config(ctx, rl.FunctionConfig)
	r.fnRunner.Run(ctx, rl.FunctionConfig, rl.Items)
	return true, nil
}

func (r *runnerProcessor) config(ctx *Context, o *KubeObject) {
	fnName := reflect.ValueOf(r.fnRunner).Elem().Type().Name()
	switch true {
	case o.IsEmpty():
		ctx.Result("`FunctionConfig` is not given", Info)
	case o.IsGVK("", "v1", "ConfigMap"):
		data := o.NestedStringMapOrDie("data")
		fnRunnerElem := reflect.ValueOf(r.fnRunner).Elem()
		for k, v := range data {
			lowerKey := strings.ToLower(k)
			for i := 0; i < fnRunnerElem.NumField(); i++ {
				switch fnRunnerElem.Field(i).Kind() {
				case reflect.Map:
					if "data" == strings.ToLower(fnRunnerElem.Type().Field(i).Name) {
						fnRunnerElem.Field(i).Set(reflect.ValueOf(v))
					}
				case reflect.String:
					if lowerKey == strings.ToLower(fnRunnerElem.Type().Field(i).Name) {
						fnRunnerElem.Field(i).SetString(v)
						break
					}
				}
			}
		}
	case o.IsGVK("fn.kpt.dev", "v1alpha1", fnName):
		o.AsOrDie(r.fnRunner)
	default:
		ctx.ResultErrAndDie(fmt.Sprintf("unknown FunctionConfig `%v`, expect %v", o.GetKind(), fnName), o)
	}
}
