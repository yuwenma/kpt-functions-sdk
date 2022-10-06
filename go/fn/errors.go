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
	"strings"
)

const pathDelimitor = "."

// ErrMissingFnConfig raises error if a required functionConfig is missing.
type ErrMissingFnConfig struct{}

func (ErrMissingFnConfig) Error() string {
	return "unable to find the functionConfig in the resourceList"
}

func NewErrUnmatchedField(obj SubObject, fields []string, dt any) *ErrUnmatchedField {
	relativefields := strings.Join(fields, pathDelimitor)
	obj.fieldpath += pathDelimitor + relativefields
	return &ErrUnmatchedField{
		SubObject: &obj, DataType: fmt.Sprintf("%T", dt),
	}
}

// errSubObjectFields raises if the SubObject operation panics.
type ErrUnmatchedField struct {
	SubObject *SubObject
	DataType  string
}

func (e *ErrUnmatchedField) Error() string {
	return fmt.Sprintf("Resource(apiVersion=%v, kind=%v) has unmatched field type %q in fieldpath %v",
		e.SubObject.parentGVK.GroupVersion(), e.SubObject.parentGVK.Kind, e.DataType, e.SubObject.fieldpath)
}

type ErrAttemptToTouchUpstreamIdentifier struct{}

func (ErrAttemptToTouchUpstreamIdentifier) Error() string {
	return fmt.Sprintf("annotation %v is managed by kpt and should not be modified", UpstreamIdentifier)
}

type ErrInternalAnnotation struct {
	Message string
}

func (e *ErrInternalAnnotation) Error() string {
	return e.Message
}
