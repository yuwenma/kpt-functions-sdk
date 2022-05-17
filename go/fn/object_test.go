package fn

import (
	"reflect"
	"testing"
)

func TestIsNamespaceScoped(t *testing.T) {
	testdata := map[string]struct{
		input []byte
		expected bool
 	}{
		"k8s resource, namespace scoped but unset": {
			input: []byte(`
apiVersion: v1
kind: ResourceQuota
metadata:
  name: example
spec:
  hard:
    limits.cpu: '10'
`),
			expected: true,
		},
		"k8s resource, cluster scoped": {
			input: []byte(`
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata: 
  name: example
subjects:
- kind: ServiceAccount
  name: example
  apiGroup: rbac.authorization.k8s.io
`),
			expected: false,
		},
		"custom resource, namespace set": {
			input: []byte(`
apiVersion: kpt-test
kind: KptTestResource
metadata: 
  name: example
  namespace: example
`),
			expected: true,
		},
		"custom resource, namespace unset": {
			input: []byte(`
apiVersion: kpt-test
kind: KptTestResource
metadata: 
  name: example
`),
			expected: false,
		},
	}
	for description, data := range testdata {
		o, _ := ParseKubeObject(data.input)
		if o.IsNamespaceScoped() != data.expected {
			t.Errorf("%v failed, resource namespace scope: got %v, want  %v", description, o.IsNamespaceScoped(), data.expected)
		}
	}
}

var noFnConfigResourceList = []byte(`apiVersion: config.kubernetes.io/v1
kind: ResourceList
`)
func TestNilFnConfigResourceList(t *testing.T) {
	rl, _ := ParseResourceList(noFnConfigResourceList)
	if rl.FunctionConfig == nil	{
		t.Errorf("Empty functionConfig in ResourceList should still be initialized to avoid nil pointer error")
	}
	if !rl.FunctionConfig.IsEmpty() {
		t.Errorf("The dummy fnConfig should be surfaced and checkable.")
	}
	// Check that FunctionConfig should be able to call KRM methods even if its Nil"
	{
		if rl.FunctionConfig.GetKind() != "" {
			t.Errorf("Nil KubeObject cannot call GetKind()")
		}
		if rl.FunctionConfig.GetAPIVersion() != "" {
			t.Errorf("Nil KubeObject cannot call GetAPIVersion()")
		}
		if rl.FunctionConfig.GetName() != "" {
			t.Errorf("Nil KubeObject cannot call GetName()")
		}
		if rl.FunctionConfig.GetNamespace() != "" {
			t.Errorf("Nil KubeObject cannot call GetNamespace()")
		}
		if rl.FunctionConfig.GetAnnotation("X") != "" {
			t.Errorf("Nil KubeObject cannot call GetAnnotation()")
		}
		if rl.FunctionConfig.GetLabel("Y") != "" {
			t.Errorf("Nil KubeObject cannot call GetLabel()")
		}
	}
	// Check that nil FunctionConfig can use SubObject methods.
	{
		_, found, err := rl.FunctionConfig.NestedString("not-exist")
		if found || err != nil {
			t.Errorf("Nil KubeObject shall not have the field path `not-exist` exist, and not expect errors")
		}
	}
	// Check that nil FunctionConfig should be editable.
	{
		rl.FunctionConfig.SetKind("CustomFn")
		if rl.FunctionConfig.GetKind() != "CustomFn" {
			t.Errorf("Nil KubeObject cannot call SetKind()")
		}
		rl.FunctionConfig.SetAPIVersion("kpt.fn/v1")
		if rl.FunctionConfig.GetAPIVersion() != "kpt.fn/v1" {
			t.Errorf("Nil KubeObject cannot call SetAPIVersion()")
		}
		rl.FunctionConfig.SetName("test")
		if rl.FunctionConfig.GetName() != "test" {
			t.Errorf("Nil KubeObject cannot call SetName()")
		}
		rl.FunctionConfig.SetNamespace("test-ns")
		if rl.FunctionConfig.GetNamespace() != "test-ns" {
			t.Errorf("Nil KubeObject cannot call SetNamespace()")
		}
		rl.FunctionConfig.SetAnnotation("k", "v")
		if rl.FunctionConfig.GetAnnotation("k") != "v" {
			t.Errorf("Nil KubeObject cannot call SetAnnotation()")
		}
		rl.FunctionConfig.SetLabel("k", "v")
		if rl.FunctionConfig.GetLabel("k") != "v" {
			t.Errorf("Nil KubeObject cannot call SetLabel()")
		}
		if rl.FunctionConfig.IsEmpty() {
			t.Errorf("The modified fnConfig is not nil.")
		}
	}
}

var deploymentResourceList =  []byte(`apiVersion: config.kubernetes.io/v1
kind: ResourceList
items:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: nginx-deployment
    labels:
      app: nginx
    annotations:
      config.kubernetes.io/index: '0'
      config.kubernetes.io/path: 'resources.yaml'
      internal.config.kubernetes.io/index: '0'
      internal.config.kubernetes.io/path: 'resources.yaml'
      internal.config.kubernetes.io/seqindent: 'compact'
  spec:
    replicas: 3
    selector:
      matchLabels:
        app: nginx
    paused: true
    strategy:
      type: Recreate
    template:
      metadata:
        labels:
          app: nginx
      spec:
        nodeSelector:
          disktype: ssd
        containers:
        - name: nginx
          image: nginx:1.14.2
          ports:
          - containerPort: 80
    fakeStringSlice:
    - test1
    - test2
`)
func TestGetNestedFields(t *testing.T) {
	rl, _ := ParseResourceList(deploymentResourceList)
	deployment := rl.Items[0]
	// Style 1, using concatenated fields in  NestedType function.
	if intVal := deployment.NestedInt64OrDie("spec", "replicas"); intVal != 3 {
		t.Errorf("deployment .spec.replicas expected to be 3, got %v", intVal)
	}
	if boolVal := deployment.NestedBoolOrDie("spec", "paused"); boolVal != true {
		t.Errorf("deployment .spec.paused expected to be true, got %v", boolVal)
	}
	if stringVal := deployment.NestedStringOrDie("spec", "strategy", "type"); stringVal != "Recreate" {
		t.Errorf("deployment .spec.strategy.type expected to be `Recreate`, got %v", stringVal)
	}
	if stringMapVal := deployment.NestedStringMapOrDie("spec", "template", "spec", "nodeSelector"); !reflect.DeepEqual(stringMapVal, map[string]string{"disktype": "ssd"}){
		t.Errorf("deployment .spec.template.spec.nodeSelector expected to get {`disktype`: `ssd`}, got %v", stringMapVal)
	}
	if sliceVal := deployment.NestedSliceOrDie("spec", "template", "spec", "containers"); sliceVal[0].NestedStringOrDie("name") != "nginx" {
		t.Errorf("deployment .spec.template.spec.containers[0].name expected to get `nginx`, got %v", sliceVal[0].NestedStringOrDie("name") )
	}
	if stringSliceVal := deployment.NestedStringSliceOrDie("spec", "fakeStringSlice"); !reflect.DeepEqual(stringSliceVal, []string{"test1", "test2"}) {
		t.Errorf("deployment .spec.fakeStringSlice expected to get [`test1`, `test2`], got %v", stringSliceVal)
	}
	// Style 2, get each struct layer by type.
	spec := deployment.GetMap("spec")
	if intVal := spec.GetInt("replicas"); intVal != 3 {
		t.Errorf("deployment .spec.replicas expected to be 3, got %v", intVal)
	}
	if boolVal := spec.GetBool("paused"); boolVal != true {
		t.Errorf("deployment .spec.paused expected to be true, got %v", boolVal)
	}
	if stringVal := spec.GetMap("strategy").GetString("type"); stringVal != "Recreate" {
		t.Errorf("deployment .spec.strategy.type expected to be `Recreate`, got %v", stringVal)
	}
	tmplSpec := spec.GetMap("template").GetMap("spec")
	if stringMapVal := tmplSpec.GetMap("nodeSelector"); stringMapVal.GetString("disktype") != "ssd" {
		t.Errorf("deployment .spec.template.spec.nodeSelector expected to get {`disktype`: `ssd`}, got %v", stringMapVal.GetString("disktype"))
	}
	if sliceVal := tmplSpec.GetSlice("containers"); sliceVal[0].GetString("name") != "nginx" {
		t.Errorf("deployment .spec.template.spec.containers[0].name expected to get `nginx`, got %v", sliceVal[0].NestedStringOrDie("name") )
	}
}

func TestSetNestedFields(t *testing.T) {
	o := NewEmptyKubeObject()
	o.SetNestedStringOrDie("some starlark script...", "source")
	if stringVal := o.NestedStringOrDie("source"); stringVal != "some starlark script..." {
		t.Errorf("KubeObject .source expected to get \"some starlark script...\", got %v", stringVal)
	}
	o.SetNestedStringMapOrDie(map[string]string{"tag1": "abc", "tag2": "test1"}, "tags")
	if stringMapVal := o.NestedStringOrDie("tags", "tag2") ; stringMapVal != "test1" {
		t.Errorf("KubeObject .tags.tag2 expected to get `test1`, got %v", stringMapVal)
	}
	o.SetNestedStringSliceOrDie([]string{"lable1", "lable2"}, "labels")
	if stringSliceVal := o.NestedStringSliceOrDie("labels"); !reflect.DeepEqual(stringSliceVal, []string{"lable1", "lable2"}) {
		t.Errorf("KubeObject .labels expected to get [`lable1`, `lable2`], got %v", stringSliceVal)
	}
}