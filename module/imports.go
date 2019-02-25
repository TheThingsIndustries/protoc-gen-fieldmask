// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package module

import (
	"fmt"
	"sort"
	"strings"
)

type importMap map[string]string

func (m importMap) Add(name, pkg string) error {
	if name == "" && pkg == "" {
		return nil
	}
	if name == "" {
		return fmt.Errorf("import name must be specified")
	}
	if pkg == "" {
		return fmt.Errorf("package path must be specified")
	}

	if v, ok := m[name]; ok && v != pkg {
		return fmt.Errorf("import name clash at `%s`. Imported `%s` and `%s`", name, pkg, v)
	}
	m[name] = pkg
	return nil
}

func (m importMap) GoString() string {
	var importString string
	switch len(m) {
	case 0:
	case 1:
		for name, pkg := range m {
			importString = fmt.Sprintf(`
import %s "%s"`, name, pkg)
		}
	default:
		importLines := make([]string, 0, len(m))
		for name, pkg := range m {
			importLines = append(importLines, fmt.Sprintf(`	%s "%s"`, name, pkg))
		}
		sort.Slice(importLines, func(i, j int) bool {
			return strings.Fields(importLines[i])[1] < strings.Fields(importLines[j])[1]
		})
		importString = fmt.Sprintf(`
import (
%s
)`,
			strings.Join(importLines, "\n"))
	}
	return importString
}
