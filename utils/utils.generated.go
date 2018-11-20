// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package utils

import (
	"sort"
	"strings"

	"github.com/gogo/protobuf/types"
)

// CleanPaths cleans the given field mask paths. It returns a sorted slice of
// unique paths without any child paths that are already covered by parent paths.
func CleanPaths(paths ...string) []string {
	unique := make(map[string]struct{}, len(paths))
	for _, path := range paths {
		unique[path] = struct{}{}
	}
	for path := range unique {
		parts := strings.Split(path, ".")
		if len(parts) == 1 {
			continue
		}
		for i := 1; i < len(parts); i++ {
			if _, ok := unique[strings.Join(parts[:1], ".")]; ok {
				delete(unique, path)
			}
		}
	}
	out := make([]string, 0, len(unique))
	for path := range unique {
		out = append(out, path)
	}
	sort.Strings(out)
	return out
}

// CleanFieldMask returns a new FieldMask with cleaned paths (see CleanPaths).
func CleanFieldMask(mask *types.FieldMask) *types.FieldMask {
	return &types.FieldMask{Paths: CleanPaths(mask.Paths...)}
}

// TopLevelPaths returns the list of paths
func TopLevelPaths(paths ...string) []string {
	topLevel := make([]string, len(paths))
	for i, path := range paths {
		parts := strings.SplitN(path, ".", 2)
		topLevel[i] = parts[0]
	}
	return CleanPaths(topLevel...)
}

// PathsWithPrefix returns the list of paths, each with the given prefix prepended.
func PathsWithPrefix(prefix string, paths ...string) []string {
	if !strings.HasSuffix(prefix, ".") {
		prefix += "."
	}
	out := make([]string, len(paths))
	for i, path := range paths {
		out[i] = prefix + path
	}
	return out
}

// PathsWithoutPrefix returns the paths that contain the given prefix, but
// without that prefix.
func PathsWithoutPrefix(prefix string, paths ...string) []string {
	if !strings.HasSuffix(prefix, ".") {
		prefix += "."
	}
	out := make([]string, 0, len(paths))
	for _, path := range paths {
		if !strings.HasPrefix(path, prefix) {
			continue
		}
		out = append(out, strings.TrimPrefix(path, prefix))
	}
	return out
}

// FieldMaskWithoutPrefix returns a FieldMask with paths without the given prefix
// (see PathsWithoutPrefix).
func FieldMaskWithoutPrefix(mask *types.FieldMask, prefix string) *types.FieldMask {
	return &types.FieldMask{Paths: PathsWithoutPrefix(prefix, mask.Paths...)}
}
