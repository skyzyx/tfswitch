// MIT License
//
// Copyright (c) 2018 warrensbox
// Copyright (c) 2024 Ryan Parman <https://ryanparman.com>
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

package lib_test

import (
	"reflect"
	"testing"

	"github.com/skyzyx/tfswitch/lib"
)

// TestNewCommand : pass value and check if returned value is a pointer
func TestNewCommand(t *testing.T) {
	testCmd := "terraform"
	cmd := lib.NewCommand(testCmd)

	if reflect.ValueOf(cmd).Kind() == reflect.Ptr {
		t.Logf("Value returned is a pointer %v [expected]", cmd)
	} else {
		t.Errorf("Value returned is not a pointer %v [expected", cmd)
	}
}

// TestPathList : check if bin path exist
func TestPathList(t *testing.T) {
	testCmd := ""
	cmd := lib.NewCommand(testCmd)
	listBin := cmd.PathList()

	if listBin == nil {
		t.Error("No bin path found [unexpected]")
	} else {
		t.Logf("Found bin path [expected]")
	}
}

type Command struct {
	name string
}

// TestFind : check common "cd" command exist
// This is assuming that Windows and linux has the "cd" command
func TestFind(t *testing.T) {
	testCmd := "cd"
	cmd := lib.NewCommand(testCmd)

	next := cmd.Find()
	for path := next(); len(path) > 0; path = next() {
		if path != "" {
			t.Logf("Found installation path: %v [expected]\n", path)
		} else {
			t.Errorf("Unable to find '%v' command in this operating system [unexpected]", testCmd)
		}
	}
}
