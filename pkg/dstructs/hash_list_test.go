package dstructs

import (
	"testing"

	types "github.com/atyronesmith/flowt/pkg/types"
)

func TestHashLink(t *testing.T) {
	hf := func(rn *types.RuleNode) uint64 {
		return rn.Cookie
	}

	hl := NewHashList("Test HashList", hf)

	rn := types.RuleNode{Line: 0,
		Cookie:   10,
		Duration: 100}

	rn2 := types.RuleNode{Line: 1,
		Cookie:   100,
		Duration: 1000}

	hl.Add(&rn)
	if hl.Count() != 1 {
		t.Error("Invalid HastList count after add != 1", hl.Count())
	}
	hl.Add(&rn2)
	if hl.Count() != 2 {
		t.Error("Invalid HastList count after add != 1", hl.Count())
	}
}
