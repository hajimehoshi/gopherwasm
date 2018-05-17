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

package js_test

import (
	"testing"

	"github.com/hajimehoshi/gopherwasm"
)

func TestNull(t *testing.T) {
	want := "null"
	if got := js.Null.String(); got != want {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestCallback(t *testing.T) {
	ch := make(chan int)
	c := js.NewCallback(func(args []js.Value) {
		go func() {
			ch <- args[0].Int() + args[1].Int()
		}()
	})
	defer c.Dispose()

	js.ValueOf(c).Invoke(1, 2)
	got := <-ch
	want := 3
	if got != want {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
