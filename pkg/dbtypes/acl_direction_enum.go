package dbtypes

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
)

type AclDirection int

const (
	FromLPort AclDirection = iota
	ToLPort
)

func (s AclDirection) String() string {
	return toStringAclDirection[s]
}

var toStringAclDirection = map[AclDirection]string{
	FromLPort: "from-lport",
	ToLPort:   "to-lport",
}

var toIDAclDirection = map[string]AclDirection{
	"from-lport": FromLPort,
	"to-lport":   ToLPort,
}

// MarshalJSON marshals the enum as a quoted json string
func (s AclDirection) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringAclDirection[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *AclDirection) UnmarshalJSON(b []byte) error {
	var j string
	err := jsoniter.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*s = toIDAclDirection[j]
	return nil
}
