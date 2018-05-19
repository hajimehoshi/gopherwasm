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
	"fmt"
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

func (c Callback) Close() {
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

var (
	id = js.Global.Call("eval", "(function(x) { return x; })")
)

func ValueOf(x interface{}) Value {
	switch x := x.(type) {
	case Value:
		return x
	case Callback:
		return Value{
			v: js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
				// Call the function asyncly to emulate Wasm's Callback more
				// precisely.
				go func() {
					args := []Value{}
					for _, arg := range arguments {
						args = append(args, ValueOf(arg.Interface()))
					}
					x.f(args)
				}()
				return nil
			}),
		}
	case nil:
		return Null
	case bool, int8, int16, int32, uint8, uint16, uint32, float32, float64, string, []byte:
		return Value{v: id.Invoke(x)}
	case int, int64, uint, uint64, unsafe.Pointer:
		return Value{x: x}
	default:
		panic("invalid arg")
	}
}

func (v Value) Bool() bool {
	return v.v.Bool()
}

func convertArgs(args []interface{}) []interface{} {
	newArgs := []interface{}{}
	for _, arg := range args {
		v := ValueOf(arg)
		if v.x != nil {
			newArgs = append(newArgs, v.x)
		} else {
			newArgs = append(newArgs, v.v)
		}
	}
	return newArgs
}

func (v Value) Call(m string, args ...interface{}) Value {
	if v.x != nil {
		panic("invalid receiver")
	}
	return Value{v: v.v.Call(m, convertArgs(args)...)}
}

func (v Value) Float() float64 {
	if v.x != nil {
		switch x := v.x.(type) {
		case int:
			return float64(x)
		case int64:
			return float64(x)
		case uint:
			return float64(x)
		case uint64:
			return float64(x)
		case unsafe.Pointer:
			return float64(uintptr(x))
		}
		panic("not reached")
	}
	return v.v.Float()
}

func (v Value) Get(p string) Value {
	if v.x != nil {
		panic("invalid receiver")
	}
	return Value{v: v.v.Get(p)}
}

func (v Value) Index(i int) Value {
	if v.x != nil {
		panic("invalid receiver")
	}
	return Value{v: v.v.Index(i)}
}

func (v Value) Int() int {
	if v.x != nil {
		switch x := v.x.(type) {
		case int:
			return x
		case int64:
			return int(x)
		case uint:
			return int(x)
		case uint64:
			return int(x)
		case unsafe.Pointer:
			return int(uintptr(x))
		}
		panic("not reached")
	}
	return v.v.Int()
}

func (v Value) Invoke(args ...interface{}) Value {
	if v.x != nil {
		panic("invalid receiver")
	}
	return Value{v: v.v.Invoke(convertArgs(args)...)}
}

func (v Value) Length() int {
	return v.v.Length()
}

func (v Value) New(args ...interface{}) Value {
	if v.x != nil {
		panic("invalid receiver")
	}
	return Value{v: v.v.New(convertArgs(args)...)}
}

func (v Value) Set(p string, x interface{}) {
	if v.x != nil {
		panic("invalid receiver")
	}
	v.v.Set(p, x)
}

func (v Value) SetIndex(i int, x interface{}) {
	if v.x != nil {
		panic("invalid receiver")
	}
	v.v.SetIndex(i, x)
}

func (v Value) String() string {
	if v.x != nil {
		return fmt.Sprintf("%v", v.x)
	}
	return v.v.String()
}
