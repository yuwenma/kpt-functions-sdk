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
	"strings"
)

type ResourceIdentifier struct {
	Group     string
	Version   string
	Kind      string
	Name      string
	Namespace string
}

func (o *KubeObject) GetOriginId() ResourceIdentifier {
	upstreamId := o.GetAnnotation(UpstreamIdentifier)
	if upstreamId != "" {
		segments := strings.Split(upstreamId, "|")
		return ResourceIdentifier{
			Group:     segments[0],
			Kind:      segments[1],
			Namespace: segments[2],
			Name:      segments[3],
		}
	}
	group, _ := ParseGroupVersion(o.GetAPIVersion())
	return ResourceIdentifier{
		Group:     group,
		Kind:      o.GetKind(),
		Namespace: o.GetNamespace(),
		Name:      o.GetName(),
	}
}

func ParseGroupVersion(apiVersion string) (group, version string) {
	if i := strings.Index(apiVersion, "/"); i > -1 {
		return apiVersion[:i], apiVersion[i+1:]
	}
	return "", apiVersion
}
