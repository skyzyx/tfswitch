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

package lib

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Command : type string
type Command struct {
	name string
}

// NewCommand : get command
func NewCommand(name string) *Command {
	return &Command{name: name}
}

// PathList : get bin path list
func (cmd *Command) PathList() []string {
	path := os.Getenv("PATH")
	return strings.Split(path, string(os.PathListSeparator))
}

func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	return fileInfo.IsDir()
}

func isExecutable(path string) bool {
	if isDir(path) {
		return false
	}

	fileInfo, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}

	if runtime.GOOS == "windows" {
		return true
	}

	if fileInfo.Mode()&0o111 != 0 {
		return true
	}

	return false
}

// Find : find all bin path
func (cmd *Command) Find() func() string {
	pathChan := make(chan string)
	go func() {
		for _, p := range cmd.PathList() {
			if !isDir(p) {
				continue
			}
			fileList, err := os.ReadDir(p)
			if err != nil {
				continue
			}

			for _, f := range fileList {
				path := filepath.Join(p, f.Name())
				if isExecutable(path) && f.Name() == cmd.name {
					pathChan <- path
				}
			}
		}
		pathChan <- ""
	}()

	return func() string {
		return <-pathChan
	}
}
