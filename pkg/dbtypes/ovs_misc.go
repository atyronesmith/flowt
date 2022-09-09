package dbtypes

import (
	jsoniter "github.com/json-iterator/go"
)

func (b *LogicalFlowSB) UnmarshalJSON(data []byte) error {

	var tableId int = 0
	var priority int = 0

	type tmpLogicalFlowSB LogicalFlowSB
	var tmp tmpLogicalFlowSB = tmpLogicalFlowSB {
		TableId: &tableId,
		Priority: &priority,
	}

	jsoniter.Unmarshal(data,&tmp)
	
	(*b) = LogicalFlowSB(tmp)
	
	return nil
}
