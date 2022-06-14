package fn

import (
	"strings"
)

type ResourceIdentifier struct {
	Group     string
	Version   string
	Kind      string
	Name      string
	Namespace string
}

func GetResId(o *KubeObject) ResourceIdentifier {
	group, version := parseGroupVersion(o.GetAPIVersion())
	return ResourceIdentifier{
		Group:     group,
		Version:   version,
		Kind:      o.GetKind(),
		Name:      o.GetName(),
		Namespace: o.GetNamespace(),
	}
}

func GetOriginResIds(o *KubeObject) []ResourceIdentifier {
	names := strings.Split(o.GetAnnotation(BuildAnnotationPreviousNames), ",")
	if len(names) == 0 {
		return []ResourceIdentifier{}
	}
	namespaces := strings.Split(o.GetAnnotation(BuildAnnotationPreviousNamespaces), ",")
	kinds := strings.Split(o.GetAnnotation(BuildAnnotationPreviousKinds), ",")
	var originIDs []ResourceIdentifier
	group, version := parseGroupVersion(o.GetAPIVersion())
	for i := range names {
		originIDs = append(originIDs,
			ResourceIdentifier{
				Group:     group,
				Version:   version,
				Kind:      kinds[i],
				Name:      names[i],
				Namespace: namespaces[i],
			})
	}
	return originIDs
}

func parseGroupVersion(apiVersion string) (group, version string) {
	if i := strings.Index(apiVersion, "/"); i > -1 {
		return apiVersion[:i], apiVersion[i+1:]
	}
	return "", apiVersion
}

func (o *KubeObject) UpdateOriginResId(kind, namespace, name string) {
	// Read origin kinds, namespaces and names from "Previous*" annotations
	var kinds, namespaces, names []string
	if o.GetAnnotation(BuildAnnotationPreviousNames) != "" {
		kinds = strings.Split(o.GetAnnotation(BuildAnnotationPreviousKinds), ",")
		namespaces = strings.Split(o.GetAnnotation(BuildAnnotationPreviousNamespaces), ",")
		names = strings.Split(o.GetAnnotation(BuildAnnotationPreviousNames), ",")
	}

	// Cleanup origin kinds, namespaces and names to exclude the current and pass-in kind, namespace and name.
	var newKinds, newNamespaces, newNames []string
	for i := range names {
		if kinds[i] == o.GetKind() && namespaces[i] == o.GetNamespace() && names[i] == o.GetName() {
			continue
		}
		if kinds[i] == kind && namespaces[i] == namespace && names[i] == name {
			continue
		}
		newKinds = append(newKinds, kinds[i])
		newNamespaces = append(newNamespaces, namespaces[i])
		newNames = append(newNames, names[i])
	}

	// Append pass-in kind, namespace and name if they do not match current.
	if kind != o.GetKind() || namespace != o.GetNamespace() || name != o.GetName() {
		newKinds = append(newKinds, kind)
		newNamespaces = append(newNamespaces, namespace)
		newNames = append(newNames, name)
	}

	// Update "Previous*" annotations
	if len(newNames) > 0 {
		o.SetAnnotation(BuildAnnotationPreviousKinds, strings.Join(newKinds, ","))
		o.SetAnnotation(BuildAnnotationPreviousNamespaces, strings.Join(newNamespaces, ","))
		o.SetAnnotation(BuildAnnotationPreviousNames, strings.Join(newNames, ","))
	}
}
