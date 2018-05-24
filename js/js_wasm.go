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
	"syscall/js"
)

var (
	Undefined = js.Undefined
	Null      = js.Null
	Global    = js.Global
)

type Callback = js.Callback

func NewCallback(f func([]Value)) Callback {
	return js.NewCallback(f)
}

func NewEventCallback(preventDefault, stopPropagation, stopImmediatePropagation bool, fn func(event Value)) Callback {
	return js.NewEventCallback(preventDefault, stopPropagation, stopImmediatePropagation, fn)
}

type Error = js.Error

type Value = js.Value

func ValueOf(x interface{}) Value {
	return js.ValueOf(x)
}
