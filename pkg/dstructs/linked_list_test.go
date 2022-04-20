package dstructs

import (
	"fmt"
	"math/rand"
	"testing"

	types "github.com/atyronesmith/flowt/pkg/types"
)

func TestAddNode(t *testing.T) {

	of := func(rn1 *types.RuleNode, rn2 *types.RuleNode) int {
		if rn1.Table < rn2.Table {
			return -1
		}

		if rn1.Table == rn2.Table {
			return -(int(rn1.Priority) - int(rn2.Priority))
		}

		return 1
	}

	testList := NewList("Test List", of)

	rn := types.RuleNode{Line: 0,
		Cookie:   0,
		Duration: rand.Float64(),
		Table:    0}

	testList.InsertOrdered(&rn)

	nodeCount := rand.Intn(100) + 1

	for i := 0; i < nodeCount-1; i++ {

		rn := types.RuleNode{Line: uint(i),
			Cookie:   uint64(rand.Intn(5)),
			Duration: rand.Float64(),
			Table:    uint32(i % 10)}

		testList.InsertOrdered(&rn)
	}
	if testList.Len() != nodeCount {
		t.Errorf("Invalid List length after add added %d, len == %d", nodeCount, testList.Len())
	}

	fp := func(n *types.RuleNode) int {
		if n.Cookie == 0 {
			return 0
		}
		return 1
	}
	nodes := testList.Filter(fp)
	if len(nodes) == 2 {
		t.Error("Filter() failed.")
	}

	fmt.Printf("%s\n", testList.String())
}
