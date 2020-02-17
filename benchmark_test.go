package main

import (
	"testing"

	"github.com/TheThingsIndustries/protoc-gen-fieldmask/testdata"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func BenchmarkSetFields(t *testing.B) {
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
			Name: "destination testOneof mismatch",
			Destination: &testdata.Test{
				TestOneof: &testdata.Test_CustomNameOneof{
					CustomNameOneof: 42,
				},
				G: &testdata.Empty{},
			},
			Source: &testdata.Test{
				TestOneof: &testdata.Test_D{
					D: 42,
				},
			},
			Paths: []string{"testOneof.d"},
			Result: &testdata.Test{
				TestOneof: &testdata.Test_CustomNameOneof{
					CustomNameOneof: 42,
				},
				G: &testdata.Empty{},
			},
			ErrorAssertion: func(t *testing.T, err error) bool { return assertions.New(t).So(err, should.BeError) },
		},
		{
			Name: "source testOneof mismatch",
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
				TestOneof: &testdata.Test_D{
					D: 42,
				},
				G: &testdata.Empty{},
			},
			ErrorAssertion: func(t *testing.T, err error) bool { return assertions.New(t).So(err, should.BeError) },
		},
		{
			Name: "unset testOneof",
			Destination: &testdata.Test{
				TestOneof: &testdata.Test_D{
					D: 42,
				},
				G: &testdata.Empty{},
			},
			Source: &testdata.Test{},
			Paths:  []string{"testOneof.d"},
			Result: &testdata.Test{
				G: &testdata.Empty{},
			},
		},
		{
			Name: "set non-existing testOneof",
			Destination: &testdata.Test{
				G: &testdata.Empty{},
			},
			Source: &testdata.Test{},
			Paths:  []string{"testOneof.e"},
			Result: &testdata.Test{
				G: &testdata.Empty{},
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
		t.Run(tc.Name, func(t *testing.B) {
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				tc.Destination.ValidateFields(tc.Paths...)
				tc.Destination.SetFields(tc.Source, tc.Paths...)
			}
		})
	}
}
