package types

import (
	"fmt"

	"github.com/atyronesmith/flowt/pkg/utils"
)

type OFReg int

const (
    //   logical flow flags
    //          The  logical flags are intended to handle keeping context
    //          between tables in order to decide which rules  in  subse?
    //          quent  tables  are  matched. These values only have local
    //          significance and are not meaningful between chassis.  OVN
    //          stores the logical flags in Open vSwitch extension regis?
    //          ter number 10.
	R10 OFReg = iota
    //   conntrack zone fields for routers
    //          Fields that denote  the  connection  tracking  zones  for
    //          routers.  These  values  only have local significance and
    //          are not meaningful between chassis. OVN stores  the  zone
    //          information  for  north to south traffic (for DNATting or
    //          ECMP symmetric replies) in Open vSwitch extension  regis?
    //          ter  number  11  and  zone information for south to north
    //          traffic (for SNATing) in Open vSwitch extension  register
    //          number 12.
	R11
	R12
	// conntrack zone field for logical ports
    //          A field that denotes the  connection  tracking  zone  for
    //          logical  ports. The value only has local significance and
    //          is not meaningful between chassis. This is initialized to
    //          0  at  the beginning of the logical ingress pipeline. OVN
    //          stores this in Open vSwitch extension register number 13.
	R13
    //   logical input port field
    //          A  field  that  denotes  the  logical port from which the
    //          packet entered the logical datapath. OVN stores  this  in
    //          Open vSwitch extension register number 14.
	R14
    //   logical output port field
    //          A field that denotes the  logical  port  from  which  the
    //          packet  will leave the logical datapath. This is initial?
    //          ized to 0 at the beginning of the logical  ingress  pipe?
    //          line.  OVN stores this in Open vSwitch extension register
    //          number 15.
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
	Registers [REND] int 

	// move:NXM_NX_TUN_ID[0..23]->OXM_OF_METADATA[0..23]
	// load:0x2->OXM_OF_METADATA[]
	// move:NXM_NX_TUN_METADATA0[16..30]->NXM_NX_REG14[0..14]
	OFMetadata int

	// Arrary of tables that are next
	NextTables []int

	// True if an action drops in this rule
	Drop bool
}

func NewRuleNode() *RuleNode {
	rn := new(RuleNode)

	rn.NextTables = make([]int, 0)

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

	hdr := fmt.Sprintf("%d/%d/0x%x/%d/%v", ruleNode.Line, ruleNode.Table, ruleNode.Cookie, ruleNode.Priority, ruleNode.NextTables)
	hdrCol := fmt.Sprintf("%*s", hdrColLen, hdr)

	match, _ := utils.ArrToString(ruleNode.Match, matchColLen)

	actions, _ := utils.ArrToString(ruleNode.Actions, actionColLen)

	return fmt.Sprintf("%s %*s|%s", hdrCol, matchColLen, match, actions)
}

// %v format
func (ruleNode RuleNode) String() string {
	hdrCol := fmt.Sprintf("%d/%d/0x%x/%d/%v", ruleNode.Line, ruleNode.Table, ruleNode.Cookie, ruleNode.Priority, ruleNode.NextTables)

	match, _ := utils.ArrToString(ruleNode.Match, -1)

	actions, _ := utils.ArrToString(ruleNode.Actions, -1)

	return fmt.Sprintf("[%s %s %s]", hdrCol, match, actions)
}

func (ruleNode RuleNode) HdrString() string {
	return fmt.Sprintf("[ %s/%s/%s/%s %s %s", "Line", "Table", "Cookie", "Priority", "Match", "Action")
}