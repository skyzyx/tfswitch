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
	"log"
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/skyzyx/tfswitch/lib"
)

// TestCreateSymlink : check if symlink exist-remove if exist,
// create symlink, check if symlink exist, remove symlink
func TestCreateSymlink(t *testing.T) {
	testSymlinkSrc := "/test-tfswitcher-src"

	testSymlinkDest := "/test-tfswitcher-dest"

	usr, errCurr := user.Current()
	if errCurr != nil {
		log.Fatal(errCurr)
	}
	symlinkPathSrc := filepath.Join(usr.HomeDir, testSymlinkSrc)
	symlinkPathDest := filepath.Join(usr.HomeDir, testSymlinkDest)

	ln, _ := os.Readlink(symlinkPathSrc)

	if ln != symlinkPathDest {
		t.Logf("Symlink does not exist %v [expected]", ln)
	} else {
		t.Logf("Symlink exist %v [expected]", ln)
		os.Remove(symlinkPathSrc)
		t.Logf("Removed existing symlink for testing purposes")
	}

	lib.CreateSymlink(symlinkPathDest, symlinkPathSrc)

	lnCheck, _ := os.Readlink(symlinkPathSrc)
	if lnCheck == symlinkPathDest {
		t.Logf("Symlink exist %v [expected]", lnCheck)
	} else {
		t.Logf("Symlink does not exist %v [unexpected]", lnCheck)
		t.Error("Symlink was not created")
	}

	os.Remove(symlinkPathSrc)
}

// TestRemoveSymlink : check if symlink exist-create if does not exist,
// remove symlink, check if symlink exist
func TestRemoveSymlink(t *testing.T) {
	testSymlinkSrc := "/test-tfswitcher-src"

	testSymlinkDest := "/test-tfswitcher-dest"

	usr, errCurr := user.Current()
	if errCurr != nil {
		log.Fatal(errCurr)
	}
	symlinkPathSrc := filepath.Join(usr.HomeDir, testSymlinkSrc)
	symlinkPathDest := filepath.Join(usr.HomeDir, testSymlinkDest)

	ln, _ := os.Readlink(symlinkPathSrc)

	if ln != symlinkPathDest {
		t.Logf("Symlink does exist %v [expected]", ln)
		t.Log("Creating symlink")
		if err := os.Symlink(symlinkPathDest, symlinkPathSrc); err != nil {
			t.Error(err)
		}
	}

	lib.RemoveSymlink(symlinkPathSrc)

	lnCheck, _ := os.Readlink(symlinkPathSrc)
	if lnCheck == symlinkPathDest {
		t.Logf("Symlink should not exist %v [unexpected]", lnCheck)
		t.Error("Symlink was not removed")
	} else {
		t.Logf("Symlink was removed  %v [expected]", lnCheck)
	}
}

// TestCheckSymlink : Create symlink, test if file is symlink
func TestCheckSymlink(t *testing.T) {
	testSymlinkSrc := "/test-tgshifter-src"

	testSymlinkDest := "/test-tgshifter-dest"

	usr, errCurr := user.Current()
	if errCurr != nil {
		log.Fatal(errCurr)
	}
	symlinkPathSrc := filepath.Join(usr.HomeDir, testSymlinkSrc)
	symlinkPathDest := filepath.Join(usr.HomeDir, testSymlinkDest)

	ln, _ := os.Readlink(symlinkPathSrc)

	if ln != symlinkPathDest {
		t.Log("Creating symlink")
		if err := os.Symlink(symlinkPathDest, symlinkPathSrc); err != nil {
			t.Error(err)
		}
	}

	symlinkExist := lib.CheckSymlink(symlinkPathSrc)

	if symlinkExist {
		t.Logf("Symlink does exist %v [expected]", ln)
	} else {
		t.Logf("Symlink does not exist %v [unexpected]", ln)
	}

	os.Remove(symlinkPathSrc)
}
