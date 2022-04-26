package ovsdbflow

import (
	"bytes"
	"encoding/json"
)

// TaskState represents the state of task, moving through Created, Running then Finished or Errorred
type NatType int

const (
	// Created represents the task has been created but not started yet
	Dnat NatType = iota
	//Running represents the task has started
	DnatAndSnat
	Snat
)

func (s NatType) String() string {
	return toStringNatType[s]
}

var toStringNatType = map[NatType]string{
	Dnat: "dnat",
	DnatAndSnat:   "dnat_and_snat",
	Snat:   "dnat_snat",
}

var toIDNatType = map[string]NatType{
	"dnat": Dnat,
	"dnat_and_snat":   DnatAndSnat,
	"snat": Snat,
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
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*s = toIDNatType[j]
	return nil
}
