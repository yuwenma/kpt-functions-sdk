package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	nameref "github.com/GoogleContainerTools/kpt-functions-sdk/go/fn/examples/example_registry/name-reference"
	setannotation "github.com/GoogleContainerTools/kpt-functions-sdk/go/fn/examples/example_registry/set-annotation"
	setlabels "github.com/GoogleContainerTools/kpt-functions-sdk/go/fn/examples/example_registry/set-label"
)

// Run function with flexibility according to the functionConfig type.
func main(){
	registry := fn.NewRegistry()
	registry.Register(&nameref.NameReference{})
	registry.Register(&setlabels.SetLabels{})
	registry.Register(&setannotation.SetAnnotations{})

	fn.AsMain(registry, &fn.OriginAnnotationAdder{})
}


// Run a single function
func main2(){
	fn.AsMain(&setlabels.SetLabels{}, &fn.OriginAnnotationAdder{})
}
