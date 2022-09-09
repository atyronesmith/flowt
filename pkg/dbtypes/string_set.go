package dbtypes

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type StringSet string

func (b StringSet) String() string {
	return string(b)
}

func (b *StringSet) UnmarshalJSON(data []byte) error {

	var uuid []string

	jsoniter.Unmarshal(data,&uuid)

	if string(data) != "[\"set\",[]]" {
		(*b) = "NULL"
	} else if len(uuid) != 2 {
		return fmt.Errorf("invalid StringSet JSON.  Unable to marshal: \"%s\"",string(data))
	} else {
		(*b) = StringSet(uuid[1])
	}

	return nil
}
