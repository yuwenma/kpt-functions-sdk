package fn

type Identifier struct {
	APIVersion string `json:"apiVersion,omitempty"`
	Kind string `json:"kind,omitempty"`
	nameType NameType `json:"metadata,omitempty"`
}

type NameType struct {
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type FieldSpec struct {
	Identifier Identifier
	Path []string `json:"fieldSpec,omitempty"`
}