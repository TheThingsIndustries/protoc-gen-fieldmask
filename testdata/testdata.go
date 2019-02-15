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

// Package testdata contains data used by tests.
package testdata

import "errors"

type CustomType []byte

func (t CustomType) Marshal() ([]byte, error) {
	return t, nil
}
func (t *CustomType) MarshalTo(data []byte) (n int, err error) {
	data = append(data, *t...)
	return len(*t), nil
}
func (t *CustomType) Unmarshal(data []byte) error {
	*t = make([]byte, len(data))
	copy(*t, data)
	return nil
}
func (t *CustomType) Size() int {
	return 42
}

func (t CustomType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(t) + `"`), nil
}
func (t *CustomType) UnmarshalJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("Invalid data")
	}
	*t = make([]byte, len(data)-2)
	copy(*t, data[1:len(data)-1])
	return nil
}
