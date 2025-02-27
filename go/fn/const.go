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

const (
	// internalPrefix is the prefix given to internal annotations that are used
	// internally by the orchestrator
	internalPrefix string = "internal.config.kubernetes.io/"

	// IndexAnnotation records the index of a specific resource in a file or input stream.
	IndexAnnotation string = internalPrefix + "index"

	// PathAnnotation records the path to the file the Resource was read from
	PathAnnotation string = internalPrefix + "path"

	// SeqIndentAnnotation records the sequence nodes indentation of the input resource
	SeqIndentAnnotation string = internalPrefix + "seqindent"

	// IdAnnotation records the id of the resource to map inputs to outputs
	IdAnnotation string = internalPrefix + "id"

	// InternalAnnotationsMigrationResourceIDAnnotation is used to uniquely identify
	// resources during round trip to and from a function execution. We will use it
	// to track the internal annotations and reconcile them if needed.
	InternalAnnotationsMigrationResourceIDAnnotation = internalPrefix + "annotations-migration-resource-id"

	// ConfigPrefix is the prefix given to the custom kubernetes annotations.
	ConfigPrefix string = "config.kubernetes.io/"

	// KptLocalConfig marks a KRM resource to be skipped from deploying to the cluster via `kpt live apply`.
	KptLocalConfig = ConfigPrefix + "local-config"

	KptUseOnlyPrefix   = "internal.kpt.dev/"
	UpstreamIdentifier = KptUseOnlyPrefix + "upstream-identifier"
)
