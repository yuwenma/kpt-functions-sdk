package fn

type FunctionRunner interface {
	Run(fnConfig *KubeObject, items []*KubeObject, results *Results)
}

/*
type Registry struct {
	Ctx context.Context
	fnConfigMap map[string]interface{}
}

func (r *Registry) Register(fn FunctionRunner) {
	fnValue := reflect.ValueOf(fn).Elem()
	name := fnValue.Type().Name()
	r.fnRunners[name] = fn
}
 */
