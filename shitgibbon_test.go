// Copyright (c) 2017-2024 Peter Hagelund
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

package shitgibbon

import (
	"strings"
	"testing"
)

func TestMakeNoDash(t *testing.T) {
	generator := NewGenerator()
	if generator.UseDash() {
		t.Error("Generator using dash")
	}
	sg := generator.Make()
	if len(sg) == 0 {
		t.Error("Empty shitgibbon returned")
	}
	if strings.ContainsRune(sg, '-') {
		t.Error("shitgibbon returned with dash")
	}
}

func TestMakeWithDash(t *testing.T) {
	generator := NewGenerator()
	generator.SetUseDash(true)
	if !generator.UseDash() {
		t.Error("Generator not using dash")
	}
	sg := generator.Make()
	if len(sg) == 0 {
		t.Error("Empty shitgibbon returned")
	}
	if !strings.ContainsRune(sg, '-') {
		t.Error("shitgibbon returned without dash")
	}
}

func TestMakeAllNoDash(t *testing.T) {
	generator := NewGenerator()
	if generator.UseDash() {
		t.Error("Generator using dash")
	}
	all := generator.MakeAll()
	if len(all) == 0 {
		t.Error("Empty slice returned")
	}
	sgs := make(map[string]bool)
	for _, sg := range all {
		if _, ok := sgs[sg]; ok {
			t.Errorf("Shitgibbon '%v' returned more than once", sg)
		}
		sgs[sg] = true
	}
}

func TestMakeAllWithDash(t *testing.T) {
	generator := NewGenerator()
	generator.SetUseDash(true)
	if !generator.UseDash() {
		t.Error("Generator not using dash")
	}
	all := generator.MakeAll()
	if len(all) == 0 {
		t.Error("Empty slice returned")
	}
	sgs := make(map[string]bool)
	for _, sg := range all {
		if _, ok := sgs[sg]; ok {
			t.Errorf("Shitgibbon '%v' returned more than once", sg)
		}
		sgs[sg] = true
	}
}
