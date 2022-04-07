package analysis

import (
	dstructs "github.com/atyronesmith/flowt/pkg/dstructs"
	"github.com/atyronesmith/flowt/pkg/types"
)

func GetTables[K comparable](hl *dstructs.HashList[K, types.RuleNode]) []uint32 {

	tables := make(map[uint32]int)

	nodes := hl.GetNodes()
	for _, v := range nodes {
		tables[(*v).Table] = 1
	}
	t := make([]uint32, 0)

	for k := range tables {
		t = append(t, k)
	}

	return t
}
