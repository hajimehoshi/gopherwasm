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
	"github.com/gopherjs/gopherjs/js"
)

var (
	Undefined = Value{v: js.Undefined}
	Null      = Value{v: nil}
	Global    = Value{v: js.Global}
)

type Callback func([]*js.Object)

func NewCallback(f func([]Value)) Callback {
	return func(arguments []*js.Object) {
		args := []Value{}
		for _, arg := range arguments {
			args = append(args, Value{v: arg})
		}
		f(args)
	}
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
}

func ValueOf(x interface{}) Value {
	return Value{v: js.InternalObject(x)}
}

func (v Value) Bool() bool {
	return v.v.Bool()
}

func convertArgs(args []interface{}) []interface{} {
	newArgs := []interface{}{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Value:
			newArgs = append(newArgs, arg.v)
		default:
			newArgs = append(newArgs, arg)
		}
	}
	return newArgs
}

func (v Value) Call(m string, args ...interface{}) Value {
	// TODO:
	return Value{v: v.v.Call(m, convertArgs(args)...)}
}

func (v Value) Float() float64 {
	return v.v.Float()
}

func (v Value) Get(p string) Value {
	return Value{v: v.v.Get(p)}
}

func (v Value) Index(i int) Value {
	return Value{v: v.v.Index(i)}
}

func (v Value) Int() int {
	return v.v.Int()
}

func (v Value) Invoke(args ...interface{}) Value {
	return Value{v: v.v.Invoke(convertArgs(args)...)}
}

func (v Value) Length() int {
	return v.v.Length()
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
	return v.v.String()
}

