package dbtypes

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type UUID string

func (b UUID) String() string {
	return string(b)
}

func (b *UUID) UnmarshalJSON(data []byte) error {

	var uuid []string
	
	jsoniter.Unmarshal(data,&uuid)

	if len(uuid) != 2 {
		return fmt.Errorf("invalid UUID JSON.  Unable to marshal: \"%s\"",string(data))
	} else {
		(*b) = UUID(uuid[1])
	}

	return nil
}
