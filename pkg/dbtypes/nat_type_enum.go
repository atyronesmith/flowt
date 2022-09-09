package dbtypes

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
)

type NatType int

const (
	Dnat NatType = iota
	DnatAndSnat
	Snat
)

func (s NatType) String() string {
	return toStringNatType[s]
}

var toStringNatType = map[NatType]string{
	Dnat:        "dnat",
	DnatAndSnat: "dnat_and_snat",
	Snat:        "dnat_snat",
}

var toIDNatType = map[string]NatType{
	"dnat":          Dnat,
	"dnat_and_snat": DnatAndSnat,
	"snat":          Snat,
}

// MarshalJSON marshals the enum as a quoted json string
func (s NatType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringNatType[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *NatType) UnmarshalJSON(b []byte) error {
	var j string
	err := jsoniter.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*s = toIDNatType[j]
	return nil
}
