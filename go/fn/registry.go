package fn

import (
	"reflect"
)

type registryProcessor struct {
	registry registry
}
func (r registryProcessor) Process(rl *ResourceList) (bool, error) {
	fnName := rl.FunctionConfig.GetKind()
	if processor, ok:= r.registry.fnRunners[fnName]; ok {
		return processor.Process(rl)
	}
	var processors []ResourceListProcessor
	for _, rProcessor := range r.registry.fnRunners {
		processors = append(processors, rProcessor)
	}
	return Chain(processors...).Process(rl)
}


func NewRegistry() registry {
	return registry{fnRunners: map[string]runnerProcessor{}}
}
type registry struct {
	fnRunners map[string]runnerProcessor
}

func (r *registry) Register(fn FunctionRunner) {
	fnName := reflect.ValueOf(fn).Elem().Type().Name()
	r.fnRunners[fnName] = runnerProcessor{fnRunner: fn}
}

