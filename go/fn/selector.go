package fn

type Selector interface {
	Select(items []*KubeObject) []*KubeObject
	SelectOriginal(items []*KubeObject) []*KubeObject
}


var _ Selector = &IdentifierSelector{}
type IdentifierSelector struct {
	Identifiers []*Identifier
}
func (s *IdentifierSelector) SelectOriginal(items []*KubeObject) []*KubeObject {
	return nil
}
func (s *IdentifierSelector) Select(items []*KubeObject) []*KubeObject {
	var newItems []*KubeObject
	for _, o := range items {
		for _, identifier := range s.Identifiers {
			if identifier.APIVersion != "" {
				if o.GetAPIVersion() !=  identifier.APIVersion {
					continue
				}
			}
			if identifier.Kind != "" {
				if o.GetKind() !=  identifier.Kind {
					continue
				}
			}
			if identifier.nameType.Name != "" {
				if o.GetName() !=  identifier.nameType.Name {
					continue
				}
			}
			if identifier.nameType.Namespace != "" {
				if o.GetNamespace() !=  identifier.nameType.Namespace {
					continue
				}
			}
		}
		newItems = append(newItems, o)
	}
	return newItems
}


var _ Selector = &OriginAnnotationAdder{}
type OriginAnnotationAdder struct {
	Identifiers []*Identifier
}
func (s *OriginAnnotationAdder) SelectOriginal(items []*KubeObject) []*KubeObject {
	return nil
}
func (s *OriginAnnotationAdder) Select(items []*KubeObject) []*KubeObject {
	for _, o := range items {
		o.SetAnnotation(OriginApiVersions, o.GetAPIVersion())
		o.SetAnnotation(OriginKinds, o.GetKind())
		o.SetAnnotation(OriginNames, o.GetName())
		o.SetAnnotation(OriginNamespaces, o.GetNamespace())
	}
	return items
}
