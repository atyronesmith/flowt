package dstructs

import (
	"fmt"
	"testing"

	types "github.com/atyronesmith/flowt/pkg/types"
)

func TestAddNode(t *testing.T) {

	testList := NewList[uint64,types.RuleNode]()

	rn := types.RuleNode{Line: 0,
		Cookie:   10,
		Duration: 100}

	rn2 := types.RuleNode{Line: 1,
		Cookie:   100,
		Duration: 1000}

	testList.InsertHead(&rn)
	if testList.Len() != 1 {
		t.Error("Invalid List length after add != 1")
	}
	testList.InsertHead(&rn2)
	if testList.Len() != 2 {
		t.Error("Invalid List length after add != 2")
	}

	fp := func(n *types.RuleNode) int {
		if n.Cookie == 100 {
			return 0
		}
		return 1
	}
	nodes := testList.Filter(fp)
	if len(nodes) != 1 {
		t.Error("Filter() failed.")
	}

	fmt.Printf("%s\n", testList.String())
}

