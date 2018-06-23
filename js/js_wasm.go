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

var (
	int8Array    = js.Global.Get("Int8Array")
	int16Array   = js.Global.Get("Int16Array")
	int32Array   = js.Global.Get("Int32Array")
	int64Array   = js.Global.Get("Int64Array")
	uint16Array  = js.Global.Get("Uint16Array")
	uint32Array  = js.Global.Get("Uint32Array")
	uint64Array  = js.Global.Get("Uint64Array")
	float32Array = js.Global.Get("Float32Array")
	float64Array = js.Global.Get("Float64Array")
)

func ValueOf(x interface{}) Value {
	var xh *reflect.SliceHeader
	var class js.Value
	size := 0
	switch x := x.(type) {
	case []int8:
		size = 1
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
		class = int8Array
	case []int16:
		size = 2
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
		class = int16Array
	case []int32:
		size = 4
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
		class = int32Array
	case []int64:
		size = 8
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
		class = int64Array
	case []uint16:
		size = 2
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
		class = uint16Array
	case []uint32:
		size = 4
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
		class = uint32Array
	case []uint64:
		size = 8
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
		class = uint64Array
	case []float32:
		size = 4
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
		class = float32Array
	case []float64:
		size = 8
		xh = (*reflect.SliceHeader)(unsafe.Pointer(&x))
		class = float64Array
	default:
		return js.ValueOf(x)
	}

	var b []byte
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = xh.Data
	bh.Len = xh.Len * size
	bh.Cap = xh.Cap * size

	u8 := js.ValueOf(b)
	return class.New(u8.Get("buffer"), u8.Get("byteOffset"), xh.Len)
}

func GetInternalObject(v Value) interface{} {
	return v
}
