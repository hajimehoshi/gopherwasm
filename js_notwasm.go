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

// +build !wasm

package js

import (
	"math"
	"unsafe"

	"github.com/gopherjs/gopherjs/js"
)

var (
	Undefined = Value{v: js.Undefined}
	Null      = Value{v: nil}
	Global    = Value{v: js.Global}
)

type Callback struct {
	f func([]Value)
}

func NewCallback(f func([]Value)) Callback {
	return Callback{f: f}
}

func (c Callback) Dispose() {
}

type Error struct {
	e *js.Error
}

func (e Error) Error() string {
	return e.e.Error()
}

type Value struct {
	v *js.Object
	x interface{}
}

func ValueOf(x interface{}) Value {
	switch x := x.(type) {
	case Value:
		return x
	case Callback:
		return Value{
			v: js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
				args := []Value{}
				for _, arg := range arguments {
					args = append(args, ValueOf(arg.Interface()))
				}
				x.f(args)
				return nil
			}),
		}
	case nil:
		return Null
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, unsafe.Pointer, float32, float64, string, []byte:
		return Value{x: x}
	default:
		panic("invalid arg")
	}
}

func (v Value) Bool() bool {
	if v.v != nil {
		return v.v.Bool()
	}
	switch x := v.x.(type) {
	case nil:
		return false
	case bool:
		return x
	case int:
		return x != 0
	case int8:
		return x != 0
	case int16:
		return x != 0
	case int32:
		return x != 0
	case int64:
		return x != 0
	case uint:
		return x != 0
	case uint8:
		return x != 0
	case uint16:
		return x != 0
	case uint32:
		return x != 0
	case uint64:
		return x != 0
	case unsafe.Pointer:
		return uintptr(x) != 0
	case float32:
		return x != 0 && math.IsNaN(float64(x))
	case float64:
		return x != 0 && math.IsNaN(x)
	}
	return true // TODO: Is this OK?
}

func convertArgs(args []interface{}) []interface{} {
	newArgs := []interface{}{}
	for _, arg := range args {
		v := ValueOf(arg)
		if v.v != nil {
			newArgs = append(newArgs, v.v)
		} else {
			newArgs = append(newArgs, v.x)
		}
	}
	return newArgs
}

func (v Value) Call(m string, args ...interface{}) Value {
	return Value{v: v.v.Call(m, convertArgs(args)...)}
}

func (v Value) Float() float64 {
	if v.v != nil {
		return v.v.Float()
	}
	switch x := v.x.(type) {
	case nil:
		return 0
	case bool:
		if !x {
			return 0
		}
		return 1
	case int:
		return float64(x)
	case int8:
		return float64(x)
	case int16:
		return float64(x)
	case int32:
		return float64(x)
	case int64:
		return float64(x)
	case uint:
		return float64(x)
	case uint8:
		return float64(x)
	case uint16:
		return float64(x)
	case uint32:
		return float64(x)
	case uint64:
		return float64(x)
	case unsafe.Pointer:
		return float64(uintptr(x))
	case float32:
		return float64(x)
	case float64:
		return x
	}
	return math.NaN()
}

func (v Value) Get(p string) Value {
	return Value{v: v.v.Get(p)}
}

func (v Value) Index(i int) Value {
	return Value{v: v.v.Index(i)}
}

func (v Value) Int() int {
	if v.v != nil {
		return v.v.Int()
	}
	switch x := v.x.(type) {
	case nil:
		return 0
	case bool:
		if !x {
			return 0
		}
		return 1
	case int:
		return x
	case int8:
		return int(x)
	case int16:
		return int(x)
	case int32:
		return int(x)
	case int64:
		return int(x)
	case uint:
		return int(x)
	case uint8:
		return int(x)
	case uint16:
		return int(x)
	case uint32:
		return int(x)
	case uint64:
		return int(x)
	case unsafe.Pointer:
		return int(uintptr(x))
	case float32:
		return int(x)
	case float64:
		return int(x)
	}
	return 0
}

func (v Value) Invoke(args ...interface{}) Value {
	return Value{v: v.v.Invoke(convertArgs(args)...)}
}

func (v Value) Length() int {
	if v.v != nil {
		return v.v.Length()
	}
	switch x := v.x.(type) {
	case string:
		return len(x)
	case []byte:
		return len(x)
	}
	return 0
}

func (v Value) New(args ...interface{}) Value {
	return Value{v: v.v.New(convertArgs(args)...)}
}

func (v Value) Set(p string, x interface{}) {
	v.v.Set(p, x)
}

func (v Value) SetIndex(i int, x interface{}) {
	v.v.SetIndex(i, x)
}

func (v Value) String() string {
	if v.v != nil {
		return v.v.String()
	}
	switch x := v.x.(type) {
	case nil:
		return "null"
	case bool:
		if x {
			return "true"
		}
		return "false"
	case string:
		return x
	}
	return "" // TODO
}
