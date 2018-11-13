// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package testdata

import (
	"reflect"
	"time"
)

func deepCopy(dst, src interface{}) {
	copyRecursive(reflect.ValueOf(dst), reflect.ValueOf(src))
}

// NOTE: The following block is sligthly modified https://github.com/mohae/deepcopy/tree/c48cc78d482608239f6c4c92a4abd87eb8761c90

// The MIT License (MIT)
//
// Copyright (c) 2014 Joel
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// deepCopier for delegating copy process to type
type deepCopier interface {
	DeepCopy() interface{}
}

// copyRecursive does the actual copying of the interface. It currently has
// limited support for what it can handle. Add as needed.
func copyRecursive(cpy, original reflect.Value) {
	// check for implement deepcopy.deepCopier
	if original.CanInterface() {
		if copier, ok := original.Interface().(deepCopier); ok {
			cpy.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	// handle according to original's Kind
	switch original.Kind() {
	case reflect.Ptr:
		// Get the actual value being pointed to.
		originalValue := original.Elem()

		// if  it isn't valid, return.
		if !originalValue.IsValid() {
			return
		}
		cpy.Set(reflect.New(originalValue.Type()))
		copyRecursive(cpy.Elem(), originalValue)

	case reflect.Interface:
		// If this is a nil, don't do anything
		if original.IsNil() {
			return
		}
		// Get the value for the interface, not the pointer.
		originalValue := original.Elem()

		// Get the value by calling Elem().
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(copyValue, originalValue)
		cpy.Set(copyValue)

	case reflect.Struct:
		t, ok := original.Interface().(time.Time)
		if ok {
			cpy.Set(reflect.ValueOf(t))
			return
		}
		// Go through each field of the struct and copy it.
		for i := 0; i < original.NumField(); i++ {
			// The Type's StructField for a given field is checked to see if StructField.PkgPath
			// is set to determine if the field is exported or not because CanSet() returns false
			// for settable fields.  I'm not sure why.  -mohae
			if original.Type().Field(i).PkgPath != "" {
				continue
			}
			copyRecursive(cpy.Field(i), original.Field(i))
		}

	case reflect.Slice:
		if original.IsNil() {
			return
		}
		// Make a new slice and copy each element.
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(cpy.Index(i), original.Index(i))
		}

	case reflect.Map:
		if original.IsNil() {
			return
		}
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, originalKey := range original.MapKeys() {
			originalValue := original.MapIndex(originalKey)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(copyValue, originalValue)
			copyKey := reflect.New(originalKey.Type()).Elem()
			copyRecursive(copyKey, originalKey)
			cpy.SetMapIndex(copyKey, copyValue)
		}

	default:
		cpy.Set(original)
	}
}
