// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package testdata

import (
	fmt "fmt"
	time "time"
)

var _TestFieldPaths = [...]string{"a.a.a", "a.a.b", "a.a.c", "a.a.d", "a.b", "a.c", "a.d", "a.e", "b.a.a", "b.a.b", "b.a.c", "b.a.d", "b.b", "b.c", "b.d", "b.e"}

func (*Test) FieldMaskPaths() []string {
	ret := make([]string, len(_TestFieldPaths))
	copy(ret, _TestFieldPaths[:])
	return ret
}

func (dst *Test) SetFields(src *Test, paths ...string) {
	for _, path := range paths {
		switch path {
		case "a.a.a":
			var nilPath bool
			nilPath = nilPath || src.A == nil
			nilPath = nilPath || src.A.A == nil

			switch {
			case dst.A != nil && nilPath:
			case dst.A == nil && nilPath:
				continue
			case dst.A == nil:
				dst.A = &Test_TestNested{}
			}

			switch {
			case dst.A.A != nil && nilPath:
			case dst.A.A == nil && nilPath:
				continue
			case dst.A.A == nil:
				dst.A.A = &Test_TestNested_TestNestedNested{}
			}

			if nilPath {
				var v int32
				dst.A.A.A = v
				continue
			}
			dst.A.A.A = src.A.A.A
		case "a.a.b":
			var nilPath bool
			nilPath = nilPath || src.A == nil
			nilPath = nilPath || src.A.A == nil

			switch {
			case dst.A != nil && nilPath:
			case dst.A == nil && nilPath:
				continue
			case dst.A == nil:
				dst.A = &Test_TestNested{}
			}

			switch {
			case dst.A.A != nil && nilPath:
			case dst.A.A == nil && nilPath:
				continue
			case dst.A.A == nil:
				dst.A.A = &Test_TestNested_TestNestedNested{}
			}

			if nilPath {
				var v int64
				dst.A.A.B = v
				continue
			}
			dst.A.A.B = src.A.A.B
		case "a.a.c":
			var nilPath bool
			nilPath = nilPath || src.A == nil
			nilPath = nilPath || src.A.A == nil

			switch {
			case dst.A != nil && nilPath:
			case dst.A == nil && nilPath:
				continue
			case dst.A == nil:
				dst.A = &Test_TestNested{}
			}

			switch {
			case dst.A.A != nil && nilPath:
			case dst.A.A == nil && nilPath:
				continue
			case dst.A.A == nil:
				dst.A.A = &Test_TestNested_TestNestedNested{}
			}

			if nilPath || src.A.A.C == nil {
				dst.A.A.C = nil
				continue
			}
			dst.A.A.C = make([][]byte, len(src.A.A.C))
			for i, v := range src.A.A.C {
				dst.A.A.C[i] = make([]byte, len(v))
				copy(dst.A.A.C[i], v)
			}
		case "a.a.d":
			var nilPath bool
			nilPath = nilPath || src.A == nil
			nilPath = nilPath || src.A.A == nil

			switch {
			case dst.A != nil && nilPath:
			case dst.A == nil && nilPath:
				continue
			case dst.A == nil:
				dst.A = &Test_TestNested{}
			}

			switch {
			case dst.A.A != nil && nilPath:
			case dst.A.A == nil && nilPath:
				continue
			case dst.A.A == nil:
				dst.A.A = &Test_TestNested_TestNestedNested{}
			}

			if nilPath || src.A.A.D == nil {
				dst.A.A.D = nil
				continue
			}
			deepCopy(&dst.A.A.D, &src.A.A.D)
		case "a.b":
			var nilPath bool
			nilPath = nilPath || src.A == nil

			switch {
			case dst.A != nil && nilPath:
			case dst.A == nil && nilPath:
				continue
			case dst.A == nil:
				dst.A = &Test_TestNested{}
			}

			if nilPath {
				var v []byte
				dst.A.B = v
				continue
			}
			dst.A.B = make([]byte, len(src.A.B))
			copy(dst.A.B, src.A.B)
		case "a.c":
			var nilPath bool
			nilPath = nilPath || src.A == nil

			switch {
			case dst.A != nil && nilPath:
			case dst.A == nil && nilPath:
				continue
			case dst.A == nil:
				dst.A = &Test_TestNested{}
			}

			if nilPath || src.A.C == nil {
				dst.A.C = nil
				continue
			}
			var v time.Duration
			dst.A.C = &v
			(*dst.A.C) = (*src.A.C)
		case "a.d":
			var nilPath bool
			nilPath = nilPath || src.A == nil

			switch {
			case dst.A != nil && nilPath:
			case dst.A == nil && nilPath:
				continue
			case dst.A == nil:
				dst.A = &Test_TestNested{}
			}

			if nilPath || src.A.D == nil {
				dst.A.D = nil
				continue
			}
			var v time.Time
			dst.A.D = &v
			(*dst.A.D) = time.Unix(0, (*src.A.D).UnixNano()).UTC()
		case "a.e":
			var nilPath bool
			nilPath = nilPath || src.A == nil

			switch {
			case dst.A != nil && nilPath:
			case dst.A == nil && nilPath:
				continue
			case dst.A == nil:
				dst.A = &Test_TestNested{}
			}

			if nilPath || src.A.E == nil {
				dst.A.E = nil
				continue
			}
			var v CustomType
			dst.A.E = &v
			deepCopy(&(*dst.A.E), &(*src.A.E))
		case "b.a.a":
			var nilPath bool
			nilPath = nilPath || src.CustomName == nil
			nilPath = nilPath || src.CustomName.A == nil

			switch {
			case dst.CustomName != nil && nilPath:
			case dst.CustomName == nil && nilPath:
				continue
			case dst.CustomName == nil:
				dst.CustomName = &Test_TestNested{}
			}

			switch {
			case dst.CustomName.A != nil && nilPath:
			case dst.CustomName.A == nil && nilPath:
				continue
			case dst.CustomName.A == nil:
				dst.CustomName.A = &Test_TestNested_TestNestedNested{}
			}

			if nilPath {
				var v int32
				dst.CustomName.A.A = v
				continue
			}
			dst.CustomName.A.A = src.CustomName.A.A
		case "b.a.b":
			var nilPath bool
			nilPath = nilPath || src.CustomName == nil
			nilPath = nilPath || src.CustomName.A == nil

			switch {
			case dst.CustomName != nil && nilPath:
			case dst.CustomName == nil && nilPath:
				continue
			case dst.CustomName == nil:
				dst.CustomName = &Test_TestNested{}
			}

			switch {
			case dst.CustomName.A != nil && nilPath:
			case dst.CustomName.A == nil && nilPath:
				continue
			case dst.CustomName.A == nil:
				dst.CustomName.A = &Test_TestNested_TestNestedNested{}
			}

			if nilPath {
				var v int64
				dst.CustomName.A.B = v
				continue
			}
			dst.CustomName.A.B = src.CustomName.A.B
		case "b.a.c":
			var nilPath bool
			nilPath = nilPath || src.CustomName == nil
			nilPath = nilPath || src.CustomName.A == nil

			switch {
			case dst.CustomName != nil && nilPath:
			case dst.CustomName == nil && nilPath:
				continue
			case dst.CustomName == nil:
				dst.CustomName = &Test_TestNested{}
			}

			switch {
			case dst.CustomName.A != nil && nilPath:
			case dst.CustomName.A == nil && nilPath:
				continue
			case dst.CustomName.A == nil:
				dst.CustomName.A = &Test_TestNested_TestNestedNested{}
			}

			if nilPath || src.CustomName.A.C == nil {
				dst.CustomName.A.C = nil
				continue
			}
			dst.CustomName.A.C = make([][]byte, len(src.CustomName.A.C))
			for i, v := range src.CustomName.A.C {
				dst.CustomName.A.C[i] = make([]byte, len(v))
				copy(dst.CustomName.A.C[i], v)
			}
		case "b.a.d":
			var nilPath bool
			nilPath = nilPath || src.CustomName == nil
			nilPath = nilPath || src.CustomName.A == nil

			switch {
			case dst.CustomName != nil && nilPath:
			case dst.CustomName == nil && nilPath:
				continue
			case dst.CustomName == nil:
				dst.CustomName = &Test_TestNested{}
			}

			switch {
			case dst.CustomName.A != nil && nilPath:
			case dst.CustomName.A == nil && nilPath:
				continue
			case dst.CustomName.A == nil:
				dst.CustomName.A = &Test_TestNested_TestNestedNested{}
			}

			if nilPath || src.CustomName.A.D == nil {
				dst.CustomName.A.D = nil
				continue
			}
			deepCopy(&dst.CustomName.A.D, &src.CustomName.A.D)
		case "b.b":
			var nilPath bool
			nilPath = nilPath || src.CustomName == nil

			switch {
			case dst.CustomName != nil && nilPath:
			case dst.CustomName == nil && nilPath:
				continue
			case dst.CustomName == nil:
				dst.CustomName = &Test_TestNested{}
			}

			if nilPath {
				var v []byte
				dst.CustomName.B = v
				continue
			}
			dst.CustomName.B = make([]byte, len(src.CustomName.B))
			copy(dst.CustomName.B, src.CustomName.B)
		case "b.c":
			var nilPath bool
			nilPath = nilPath || src.CustomName == nil

			switch {
			case dst.CustomName != nil && nilPath:
			case dst.CustomName == nil && nilPath:
				continue
			case dst.CustomName == nil:
				dst.CustomName = &Test_TestNested{}
			}

			if nilPath || src.CustomName.C == nil {
				dst.CustomName.C = nil
				continue
			}
			var v time.Duration
			dst.CustomName.C = &v
			(*dst.CustomName.C) = (*src.CustomName.C)
		case "b.d":
			var nilPath bool
			nilPath = nilPath || src.CustomName == nil

			switch {
			case dst.CustomName != nil && nilPath:
			case dst.CustomName == nil && nilPath:
				continue
			case dst.CustomName == nil:
				dst.CustomName = &Test_TestNested{}
			}

			if nilPath || src.CustomName.D == nil {
				dst.CustomName.D = nil
				continue
			}
			var v time.Time
			dst.CustomName.D = &v
			(*dst.CustomName.D) = time.Unix(0, (*src.CustomName.D).UnixNano()).UTC()
		case "b.e":
			var nilPath bool
			nilPath = nilPath || src.CustomName == nil

			switch {
			case dst.CustomName != nil && nilPath:
			case dst.CustomName == nil && nilPath:
				continue
			case dst.CustomName == nil:
				dst.CustomName = &Test_TestNested{}
			}

			if nilPath || src.CustomName.E == nil {
				dst.CustomName.E = nil
				continue
			}
			var v CustomType
			dst.CustomName.E = &v
			deepCopy(&(*dst.CustomName.E), &(*src.CustomName.E))
		default:
			panic(fmt.Errorf("Invaild fieldpath: '%s'", path))
		}
	}
}
