package dbtypes

import (
	"bytes"
	"encoding/json"
)

type AclAction int

const (
	AllowStateless AclAction = iota
	Allow
	AllowRelated
	Drop
	Reject
)

func (s AclAction) String() string {
	return toStringAclAction[s]
}

var toStringAclAction = map[AclAction]string{
	AllowStateless: "allow-stateless",
	Allow:          "allow",
	AllowRelated:   "allow-related",
	Drop:           "drop",
	Reject:         "reject",
}

var toIDAclAction = map[string]AclAction{
	"allow-stateless": AllowStateless,
	"allow":           Allow,
	"allow-related":   AllowRelated,
	"drop":            Drop,
	"reject":          Reject,
}

// MarshalJSON marshals the enum as a quoted json string
func (s AclAction) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringAclAction[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *AclAction) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*s = toIDAclAction[j]
	return nil
}

// action: string, one of allow-related, allow-stateless, allow, drop,  or
//  reject
//         The action to take when the ACL rule matches:
//         �      allow-stateless:  Always  forward the packet in stateless
//                manner, omitting connection tracking  mechanism,  regard?
//                less  of  other rules defined for the switch. May require
//                defining additional rules for inbound replies. For  exam?
//                ple,  if  you define a rule to allow outgoing TCP traffic
//                directed to an IP address, then you probably also want to
//                define  another rule to allow incoming TCP traffic coming
//                from this same IP address.
//         �      allow: Forward the packet. It will also send the  packets
//                through  connection  tracking  when  allow-related  rules
//                exist on the logical switch. Otherwise,  it?s  equivalent
//                to allow-stateless.
//         �      allow-related:  Forward  the  packet  and related traffic
//                (e.g. inbound replies to an outbound connection).
//         �      drop: Silently drop the packet.
//         �      reject: Drop the packet, replying with a RST for  TCP  or
//                ICMPv4/ICMPv6     unreachable     message    for    other
//                IPv4/IPv6-based protocols.
