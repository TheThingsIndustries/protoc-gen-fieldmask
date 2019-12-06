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

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/TheThingsIndustries/protoc-gen-fieldmask/testdata"
	"github.com/kr/pretty"
	"github.com/mohae/deepcopy"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

var regenerate = flag.Bool("regenerate", false, "regenerate golden files")

// When the environment variable IN_TEST is set, we skip running
// tests and instead act as protoc-gen-fieldmask. This allows the test binary to
// pass itself to protoc.
func init() {
	if os.Getenv("IN_TEST") != "" {
		main()
		os.Exit(0)
	}
}

func runProtoc(t *testing.T, args ...string) {
	protocBin := "protoc"
	if v := os.Getenv("PROTOC"); v != "" {
		protocBin = v
	}

	cmd := exec.Command(
		strings.Fields(protocBin)[0],
		append(append(strings.Fields(protocBin)[1:], fmt.Sprintf("--plugin=protoc-gen-fieldmask=%s", os.Args[0])), args...)...,
	)
	cmd.Env = append(os.Environ(), "IN_TEST=1", "DEBUG=1")
	t.Logf(`Running '%s'...`, strings.Join(cmd.Args, " "))

	out, err := cmd.CombinedOutput()
	t.Logf(`Output:
%s`,
		string(out),
	)

	if err != nil {
		t.Fatalf("Error: %s", err)
	}
}

func TestGolden(t *testing.T) {
	workDir, err := ioutil.TempDir("", "fieldmask-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(workDir)

	var paths []string
	if err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".proto") {
			return nil
		}
		paths = append(paths, path)
		return nil
	}); err != nil {
		t.Errorf("Failed to walk `testdata`: %s", err)
		t.FailNow()
	}

	runProtoc(t, append([]string{
		"-Ivendor",
		"-Itestdata",
		fmt.Sprintf("--fieldmask_out=lang=gogo,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:%s", workDir),
		fmt.Sprintf("--gogo_out=Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:%s", workDir),
	}, paths...)...)

	if err := filepath.Walk(workDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("Failed to read generated file at `%s`: %s", path, err)
			return nil
		}

		goldenPath := filepath.Join(".", strings.TrimPrefix(path, filepath.Join(workDir, "github.com", "TheThingsIndustries", "protoc-gen-fieldmask")))
		if *regenerate {
			if err := ioutil.WriteFile(goldenPath, b, 0666); err != nil {
				t.Errorf("Failed to write golden file at `%s`: %s", goldenPath, err)
			}
			return nil
		}

		golden, err := ioutil.ReadFile(goldenPath)
		if err != nil {
			t.Errorf("Failed to read golden file at `%s`: %s", goldenPath, err)
			return nil
		}

		if len(pretty.Diff(golden, b)) > 0 {
			pretty.Ldiff(t, golden, b)
			t.Fail()
		}
		return nil

	}); err != nil {
		t.Errorf("Failed to walk `%s`: %s", workDir, err)
		t.FailNow()
	}
}

func TestFieldMaskPaths(t *testing.T) {
	a := assertions.New(t)

	a.So(testdata.TestFieldPathsNested, should.Resemble, []string{
		"a",
		"a.a",
		"a.a.a",
		"a.a.b",
		"a.a.c",
		"a.a.d",
		"a.a.h",
		"a.a.h.nested_field",
		"a.a.i",
		"a.a.i.nested_field_2",
		"a.a.testNestedNestedOneOf",
		"a.a.testNestedNestedOneOf.e",
		"a.a.testNestedNestedOneOf.f",
		"a.a.testNestedNestedOneOf.g",
		"a.b",
		"a.c",
		"a.d",
		"a.e",
		"a.f",
		"a.g",
		"b",
		"b.a",
		"b.a.a",
		"b.a.b",
		"b.a.c",
		"b.a.d",
		"b.a.h",
		"b.a.h.nested_field",
		"b.a.i",
		"b.a.i.nested_field_2",
		"b.a.testNestedNestedOneOf",
		"b.a.testNestedNestedOneOf.e",
		"b.a.testNestedNestedOneOf.f",
		"b.a.testNestedNestedOneOf.g",
		"b.b",
		"b.c",
		"b.d",
		"b.e",
		"b.f",
		"b.g",
		"c",
		"c.a",
		"c.a.a",
		"c.a.b",
		"c.a.c",
		"c.a.d",
		"c.a.h",
		"c.a.h.nested_field",
		"c.a.i",
		"c.a.i.nested_field_2",
		"c.a.testNestedNestedOneOf",
		"c.a.testNestedNestedOneOf.e",
		"c.a.testNestedNestedOneOf.f",
		"c.a.testNestedNestedOneOf.g",
		"c.b",
		"c.c",
		"c.d",
		"c.e",
		"c.f",
		"c.g",
		"g",
		"h",
		"i",
		"testOneof",
		"testOneof.d",
		"testOneof.e",
		"testOneof.f",
		"testOneof.k",
		"testOneof.k.a",
		"testOneof.k.a.a",
		"testOneof.k.a.b",
		"testOneof.k.a.c",
		"testOneof.k.a.d",
		"testOneof.k.a.h",
		"testOneof.k.a.h.nested_field",
		"testOneof.k.a.i",
		"testOneof.k.a.i.nested_field_2",
		"testOneof.k.a.testNestedNestedOneOf",
		"testOneof.k.a.testNestedNestedOneOf.e",
		"testOneof.k.a.testNestedNestedOneOf.f",
		"testOneof.k.a.testNestedNestedOneOf.g",
		"testOneof.k.b",
		"testOneof.k.c",
		"testOneof.k.d",
		"testOneof.k.e",
		"testOneof.k.f",
		"testOneof.k.g",
	})
	a.So(testdata.TestFieldPathsTopLevel, should.Resemble, []string{
		"a",
		"b",
		"c",
		"g",
		"h",
		"i",
		"testOneof",
	})

	a.So(testdata.Test_TestNestedFieldPathsNested, should.Resemble, []string{
		"a",
		"a.a",
		"a.b",
		"a.c",
		"a.d",
		"a.h",
		"a.h.nested_field",
		"a.i",
		"a.i.nested_field_2",
		"a.testNestedNestedOneOf",
		"a.testNestedNestedOneOf.e",
		"a.testNestedNestedOneOf.f",
		"a.testNestedNestedOneOf.g",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
	})
	a.So(testdata.Test_TestNestedFieldPathsTopLevel, should.Resemble, []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
	})

	a.So(testdata.Test_TestNested_TestNestedNestedFieldPathsNested, should.Resemble, []string{
		"a",
		"b",
		"c",
		"d",
		"h",
		"h.nested_field",
		"i",
		"i.nested_field_2",
		"testNestedNestedOneOf",
		"testNestedNestedOneOf.e",
		"testNestedNestedOneOf.f",
		"testNestedNestedOneOf.g",
	})
	a.So(testdata.Test_TestNested_TestNestedNestedFieldPathsTopLevel, should.Resemble, []string{
		"a",
		"b",
		"c",
		"d",
		"h",
		"i",
		"testNestedNestedOneOf",
	})
}

func TestSetFields(t *testing.T) {
	for _, tc := range []struct {
		Name                        string
		Source, Destination, Result *testdata.Test
		Paths                       []string
		ErrorAssertion              func(t *testing.T, err error) bool
	}{
		{
			Name: "nil source",
			Destination: &testdata.Test{
				A: &testdata.Test_TestNested{},
				CustomName: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{},
				},
			},
			Source: nil,
			Paths:  []string{"a.b", "b.c"},
			Result: &testdata.Test{
				A: &testdata.Test_TestNested{},
				CustomName: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{},
				},
			},
		},
		{
			Name: "no paths",
			Destination: &testdata.Test{
				CustomName: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{},
				},
			},
			Source: &testdata.Test{
				A: &testdata.Test_TestNested{
					B: []byte{1, 2, 3},
				},
				CustomName: &testdata.Test_TestNested{
					B: []byte{1, 2, 4},
				},
			},
			Paths: nil,
			Result: &testdata.Test{
				CustomName: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{},
				},
			},
		},
		{
			Name: "a.b",
			Destination: &testdata.Test{
				CustomName: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{},
				},
			},
			Source: &testdata.Test{
				A: &testdata.Test_TestNested{
					B: []byte{1, 2, 3},
				},
				CustomName: &testdata.Test_TestNested{
					B: []byte{1, 2, 4},
				},
			},
			Paths: []string{"a.b"},
			Result: &testdata.Test{
				A: &testdata.Test_TestNested{
					B: []byte{1, 2, 3},
				},
				CustomName: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{},
				},
			},
		},
		{
			Name: "a.b,a.a.a,a.b,a.b,b,testOneof",
			Destination: &testdata.Test{
				TestOneof: &testdata.Test_CustomNameOneof{},
				G:         &testdata.Empty{},
			},
			Source: &testdata.Test{
				A: &testdata.Test_TestNested{
					B: []byte{1, 2, 3},
				},
				CustomName: &testdata.Test_TestNested{
					B: []byte{1, 2, 4},
				},
				TestOneof: &testdata.Test_D{
					D: 42,
				},
			},
			Paths: []string{"a.b", "a.a.a", "a.b", "a.b", "b", "testOneof"},
			Result: &testdata.Test{
				A: &testdata.Test_TestNested{
					B: []byte{1, 2, 3},
				},
				CustomName: &testdata.Test_TestNested{
					B: []byte{1, 2, 4},
				},
				TestOneof: &testdata.Test_D{
					D: 42,
				},
				G: &testdata.Empty{},
			},
		},
		{
			Name: "testOneof.d",
			Destination: &testdata.Test{
				TestOneof: &testdata.Test_D{
					D: 42,
				},
				G: &testdata.Empty{},
			},
			Source: &testdata.Test{
				TestOneof: &testdata.Test_CustomNameOneof{
					CustomNameOneof: 42,
				},
			},
			Paths: []string{"testOneof.d"},
			Result: &testdata.Test{
				TestOneof: &testdata.Test_D{},
				G:         &testdata.Empty{},
			},
		},
		{
			Name:        "testOneof.k.a.a",
			Destination: &testdata.Test{},
			Source: &testdata.Test{
				TestOneof: &testdata.Test_K{
					K: &testdata.Test_TestNested{
						A: &testdata.Test_TestNested_TestNestedNested{
							A: 42,
						},
					},
				},
			},
			Paths: []string{"testOneof.k.a.a"},
			Result: &testdata.Test{
				TestOneof: &testdata.Test_K{
					K: &testdata.Test_TestNested{
						A: &testdata.Test_TestNested_TestNestedNested{
							A: 42,
						},
					},
				},
			},
		},
		{
			Name: "testOneof.e",
			Destination: &testdata.Test{
				G: &testdata.Empty{},
			},
			Source: &testdata.Test{},
			Paths:  []string{"testOneof.e"},
			Result: &testdata.Test{
				TestOneof: &testdata.Test_CustomNameOneof{},
				G:         &testdata.Empty{},
			},
		},
		{
			Name: "non-nullable c.a",
			Destination: &testdata.Test{
				C: testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{},
				},
			},
			Source: &testdata.Test{
				C: testdata.Test_TestNested{
					B: []byte("42"),
				},
			},
			Paths: []string{"c.b"},
			Result: &testdata.Test{
				C: testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{},
					B: []byte("42"),
				},
			},
		},
		{
			Name:           "non-existent top-level field",
			Destination:    &testdata.Test{},
			Source:         &testdata.Test{},
			Paths:          []string{"42"},
			Result:         &testdata.Test{},
			ErrorAssertion: func(t *testing.T, err error) bool { return assertions.New(t).So(err, should.BeError) },
		},
		{
			Name:           "non-existent sub-field",
			Destination:    &testdata.Test{},
			Source:         &testdata.Test{},
			Paths:          []string{"41.42.43"},
			Result:         &testdata.Test{},
			ErrorAssertion: func(t *testing.T, err error) bool { return assertions.New(t).So(err, should.BeError) },
		},
		{
			Name:           "non-existent oneof",
			Destination:    &testdata.Test{},
			Source:         &testdata.Test{},
			Paths:          []string{"testOneof.42"},
			Result:         &testdata.Test{},
			ErrorAssertion: func(t *testing.T, err error) bool { return assertions.New(t).So(err, should.BeError) },
		},
		{
			Name:           "double oneofs",
			Destination:    &testdata.Test{},
			Source:         &testdata.Test{},
			Paths:          []string{"testOneof.e", "testOneof.d"},
			Result:         &testdata.Test{},
			ErrorAssertion: func(t *testing.T, err error) bool { return assertions.New(t).So(err, should.BeError) },
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			a := assertions.New(t)

			src := deepcopy.Copy(tc.Source).(*testdata.Test)
			dst := deepcopy.Copy(tc.Destination).(*testdata.Test)
			paths := deepcopy.Copy(tc.Paths).([]string)

			err := dst.SetFields(src, paths...)
			if tc.ErrorAssertion != nil {
				a.So(tc.ErrorAssertion(t, err), should.BeTrue)
			} else {
				a.So(err, should.BeNil)
			}
			a.So(src, should.Resemble, tc.Source)
			a.So(dst, should.Resemble, tc.Result)
			a.So(paths, should.Resemble, tc.Paths)
		})
	}
}

func TestValidateFields(t *testing.T) {
	for _, tc := range []struct {
		Name           string
		Message        *testdata.Test
		Paths          []string
		ErrorAssertion func(t *testing.T, err error) bool
	}{
		{
			Name:  "nil message",
			Paths: []string{"a.b", "b.c"},
		},
		{
			Name: "a.a.a",
			Message: &testdata.Test{
				A: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{
						A: 42,
					},
				},
				CustomName: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{},
				},
			},
			Paths: []string{"a.a.a"},
		},
		{
			Name:    "a.g",
			Message: &testdata.Test{},
			Paths:   []string{"a.g"},
		},
		{
			Name: "nil paths/valid",
			Message: &testdata.Test{
				A: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{
						A: 42,
						Test_TestNested_TestNestedNested_TestNestedNestedEmbed2: testdata.Test_TestNested_TestNestedNested_TestNestedNestedEmbed2{
							NestedField_2: 2,
						},
					},
					C: func(v time.Duration) *time.Duration { return &v }(43 * time.Second),
				},
				TestOneof: &testdata.Test_CustomNameOneof{
					CustomNameOneof: 6,
				},
			},
		},
		{
			Name: "nil paths/invalid testOneof unset",
			Message: &testdata.Test{
				A: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{
						A: 42,
						Test_TestNested_TestNestedNested_TestNestedNestedEmbed2: testdata.Test_TestNested_TestNestedNested_TestNestedNestedEmbed2{
							NestedField_2: 2,
						},
					},
					C: func(v time.Duration) *time.Duration { return &v }(43 * time.Second),
				},
			},
			ErrorAssertion: func(t *testing.T, err error) bool { return assertions.New(t).So(err, should.BeError) },
		},
		{
			Name: "nil paths/invalid testOneof.d",
			Message: &testdata.Test{
				A: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{
						A: 42,
						Test_TestNested_TestNestedNested_TestNestedNestedEmbed2: testdata.Test_TestNested_TestNestedNested_TestNestedNestedEmbed2{
							NestedField_2: 2,
						},
					},
					C: func(v time.Duration) *time.Duration { return &v }(43 * time.Second),
				},
				TestOneof: &testdata.Test_D{
					D: 3,
				},
			},
			ErrorAssertion: func(t *testing.T, err error) bool { return assertions.New(t).So(err, should.BeError) },
		},
		{
			Name: "nil paths/invalid a.a.a",
			Message: &testdata.Test{
				A: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{
						A: 43,
					},
				},
			},
			ErrorAssertion: func(t *testing.T, err error) bool { return assertions.New(t).So(err, should.BeError) },
		},
		{
			Name: "a.a.i/valid a.a.i.nested_field_2",
			Message: &testdata.Test{
				A: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{
						Test_TestNested_TestNestedNested_TestNestedNestedEmbed2: testdata.Test_TestNested_TestNestedNested_TestNestedNestedEmbed2{
							NestedField_2: 2,
						},
					},
				},
			},
			Paths: []string{"a.a.i"},
		},
		{
			Name: "a.a.i/invalid a.a.i.nested_field_2",
			Message: &testdata.Test{
				A: &testdata.Test_TestNested{
					A: &testdata.Test_TestNested_TestNestedNested{},
				},
			},
			Paths:          []string{"a.a.i"},
			ErrorAssertion: func(t *testing.T, err error) bool { return assertions.New(t).So(err, should.BeError) },
		},
		{
			Name:           "non-existent sub-field",
			Message:        &testdata.Test{},
			Paths:          []string{"41.42.43"},
			ErrorAssertion: func(t *testing.T, err error) bool { return assertions.New(t).So(err, should.BeError) },
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			a := assertions.New(t)

			msg := deepcopy.Copy(tc.Message).(*testdata.Test)
			paths := deepcopy.Copy(tc.Paths).([]string)

			err := msg.ValidateFields(paths...)
			if tc.ErrorAssertion != nil {
				a.So(tc.ErrorAssertion(t, err), should.BeTrue)
			} else {
				a.So(err, should.BeNil)
			}
			a.So(paths, should.Resemble, tc.Paths)
		})
	}
}
