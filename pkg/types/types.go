package types

import (
	"fmt"

	"github.com/atyronesmith/flowt/pkg/utils"
)

type OFReg int

const (
	R10 OFReg = iota
	R11
	R12
	R13
	R14
	R15
	REND
)

type RuleNode struct {
	Line        uint
	Cookie      uint64   `json:"cookie"`
	Duration    float64  `json:"duration"`
	Table       uint32   `json:"table"`
	NPackets    uint64   `json:"n_packets"`
	NBytes      uint64   `json:"n_bytes"`
	IdleAge     uint64   `json:"idle_age,omitempty"`
	IdleTimeout uint64   `json:"idle_timeout,omitempty"`
	Priority    uint     `json:"priority"`
	Match       []string `json:"match,omitempty"`
	Actions     []string `json:"actions"`

	// REG 15 -- Logical Output Port
	// REG 14 -- Logical Input Port
	// REG 13 -- conntrack zone for logical ports
	// REG 12 -- conntrack zone, zone information for south to north traffic (for SNATing)
	// REG 11 -- zone information north to south traffic (for DNATting or
	//           ECMP symmetric replies) in Open vSwitch extension
	// REG 10 -- logical flow flags
	Registers [REND]int 

	// move:NXM_NX_TUN_ID[0..23]->OXM_OF_METADATA[0..23]
	// load:0x2->OXM_OF_METADATA[]
	// move:NXM_NX_TUN_METADATA0[16..30]->NXM_NX_REG14[0..14]
	OFMetadata int

}

func NewRuleNode() *RuleNode {
	rn := new(RuleNode)

	return rn
}

func (ruleNode RuleNode) Compare(rn *RuleNode) int {
	return int(ruleNode.Line) - int(rn.Line)
}

func (ruleNode RuleNode) GetZero() RuleNode {
	var result RuleNode

	return result
}

//  %#v format
func (ruleNode RuleNode) GoString() string {
	hdrColLen := 23
	matchColLen := 50
	actionColLen := 80

	hdr := fmt.Sprintf("%d/%d/0x%x/%d", ruleNode.Line, ruleNode.Table, ruleNode.Cookie, ruleNode.Priority)
	hdrCol := fmt.Sprintf("%*s", hdrColLen, hdr)

	match, _ := utils.ArrToString(ruleNode.Match, matchColLen)

	actions, _ := utils.ArrToString(ruleNode.Actions, actionColLen)

	return fmt.Sprintf("%s %*s|%s", hdrCol, matchColLen, match, actions)
}

// %v format
func (ruleNode RuleNode) String() string {
	hdrCol := fmt.Sprintf("%d/%d/0x%x/%d", ruleNode.Line, ruleNode.Table, ruleNode.Cookie, ruleNode.Priority)

	match, _ := utils.ArrToString(ruleNode.Match, -1)

	actions, _ := utils.ArrToString(ruleNode.Actions, -1)

	return fmt.Sprintf("[%s %s %s]", hdrCol, match, actions)
}
