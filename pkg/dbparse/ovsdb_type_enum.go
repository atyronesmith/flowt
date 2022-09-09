package dbparse

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
)

// TaskState represents the state of task, moving through Created, Running then Finished or Errorred
type OVSDBType int

const (
	Unknown OVSDBType = iota
	// Created represents the task has been created but not started yet
	NB  
	//Running represents the task has started
	SB
)

func (s OVSDBType) String() string {
	return toStringOVSDBType[s]
}

func (s OVSDBType) Postfix() string {
	switch s {
		case NB:
			return "NB"
		case SB:
			return "SB"
		default:
			return "Unknown"
	}
}

func (s OVSDBType) Filename() string {
	switch s {
		case NB:
			return "ovnnb_db.db"
		case SB:
			return "ovnsb_db.db"
		default:
			return "Unknown"
	}
}

var toStringOVSDBType = map[OVSDBType]string{
	Unknown: "Unknown",
	NB: "OVN_Northbound",
	SB: "OVN_Southbound",
}

var toIDOVSDBType = map[string]OVSDBType{
	"Unknown": Unknown,
	"OVN_Northbound": NB,
	"OVN_Southbound": SB,
}

// MarshalJSON marshals the enum as a quoted json string
func (s OVSDBType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringOVSDBType[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *OVSDBType) UnmarshalJSON(b []byte) error {
	var j string
	err := jsoniter.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Unknown' in this case.
	*s = toIDOVSDBType[j]
	return nil
}
