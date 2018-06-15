// Copyright 2018 Hajime Hoshi
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

// +build wasm

package js

import (
	"reflect"
	"syscall/js"
	"unsafe"
)

var (
	Undefined = js.Undefined
	Null      = js.Null
	Global    = js.Global
)

type Callback = js.Callback

type EventCallbackFlag = js.EventCallbackFlag

const (
	PreventDefault           = js.PreventDefault
	StopPropagation          = js.StopPropagation
	StopImmediatePropagation = js.StopImmediatePropagation
)

func NewCallback(f func([]Value)) Callback {
	return js.NewCallback(f)
}

func NewEventCallback(flags EventCallbackFlag, fn func(event Value)) Callback {
	return js.NewEventCallback(flags, fn)
}

type Error = js.Error

type Value = js.Value

func ValueOf(x interface{}) Value {
	var xh *reflect.SliceHeader
	size := 0
	switch x := x.(type) {
	case []int8:
		size = 1
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
	case []int16:
		size = 2
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
	case []int32:
		size = 4
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
	case []int64:
		size = 8
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
	case []uint16:
		size = 2
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
	case []uint32:
		size = 4
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
	case []uint64:
		size = 8
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
	case []float32:
		size = 4
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
	case []float64:
		size = 8
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
	default:
		return js.ValueOf(x)
	}

	var b []byte
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = xh.Data
	bh.Len = xh.Len * size
	bh.Cap = xh.Cap * size
	return js.ValueOf(b)
}
