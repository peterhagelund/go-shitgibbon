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
	if strings.IndexRune(sg, '-') >= 0 {
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
	if strings.IndexRune(sg, '-') < 0 {
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
