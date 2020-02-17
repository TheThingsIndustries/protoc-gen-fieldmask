// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package testdata

import (
	"strings"
)

// _processPaths returns paths as a pathMap.
func _processPaths(paths []string) map[string][]string {
	if len(paths) == 0 {
		return nil
	}
	pathMap := make(map[string][]string, len(paths))
	for _, p := range paths {
		if !strings.Contains(p, ".") {
			pathMap[p] = nil
			continue
		}
		parts := strings.SplitN(p, ".", 2)
		h, t := parts[0], parts[1]
		if _, ok := pathMap[h]; ok {
			pathMap[h] = append(pathMap[h], t)
		} else {
			pathMap[h] = []string{t}
		}
	}

	return pathMap
}
