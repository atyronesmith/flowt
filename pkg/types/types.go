package types

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
}

func (ruleNode RuleNode) Compare(rn *RuleNode) int {
	return int(ruleNode.Line)-int(rn.Line)
}

func (ruleNode RuleNode) GetZero() RuleNode {
    var result RuleNode

    return result
}

func (ruleNode RuleNode) Key()  uint64 {
	return ruleNode.Cookie
}