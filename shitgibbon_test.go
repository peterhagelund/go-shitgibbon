package shitgibbon

import "testing"

func TestMake(t *testing.T) {
	sg := Make()
	if len(sg) == 0 {
		t.Error("Empty shitgibbon returned")
	}
}

func TestMakeAll(t *testing.T) {
	all := MakeAll()
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
