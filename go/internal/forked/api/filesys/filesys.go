// Copyright 2021 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

// Package filesys provides a file system abstraction,
// a subset of that provided by golang.org/pkg/os,
// with an on-disk and in-memory representation.
//
// Deprecated: use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys instead.
package filesys

import "github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys"

const (
	// Separator is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.Separator.
	Separator = filesys.Separator
	// SelfDir is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.SelfDir.
	SelfDir = filesys.SelfDir
	// ParentDir is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.ParentDir.
	ParentDir = filesys.ParentDir
)

type (
	// FileSystem is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.FileSystem.
	FileSystem = filesys.FileSystem
	// FileSystemOrOnDisk is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.FileSystemOrOnDisk.
	FileSystemOrOnDisk = filesys.FileSystemOrOnDisk
	// ConfirmedDir is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.ConfirmedDir.
	ConfirmedDir = filesys.ConfirmedDir
)

// MakeEmptyDirInMemory is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.MakeEmptyDirInMemory.
func MakeEmptyDirInMemory() FileSystem { return filesys.MakeEmptyDirInMemory() }

// MakeFsInMemory is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.MakeFsInMemory.
func MakeFsInMemory() FileSystem { return filesys.MakeFsInMemory() }

// MakeFsOnDisk is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.MakeFsOnDisk.
func MakeFsOnDisk() FileSystem { return filesys.MakeFsOnDisk() }

// NewTmpConfirmedDir is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.NewTmpConfirmedDir.
func NewTmpConfirmedDir() (filesys.ConfirmedDir, error) { return filesys.NewTmpConfirmedDir() }

// RootedPath is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.RootedPath.
func RootedPath(elem ...string) string { return filesys.RootedPath(elem...) }

// StripTrailingSeps is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.StripTrailingSeps.
func StripTrailingSeps(s string) string { return filesys.StripTrailingSeps(s) }

// StripLeadingSeps is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.StripLeadingSeps.
func StripLeadingSeps(s string) string { return filesys.StripLeadingSeps(s) }

// PathSplit is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.PathSplit.
func PathSplit(incoming string) []string { return filesys.PathSplit(incoming) }

// PathJoin is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.PathJoin.
func PathJoin(incoming []string) string { return filesys.PathJoin(incoming) }

// InsertPathPart is deprecated, use github.com/GoogleContainerTools/kpt-functions-sdk/krmfn/internal/forked/kyaml/filesys.InsertPathPart.
func InsertPathPart(path string, pos int, part string) string {
	return filesys.InsertPathPart(path, pos, part)
}
