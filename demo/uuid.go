package dbtypes

import (
	"encoding/json"
	"fmt"
)

type UUID string

func (b UUID) String() string {
	return string(b)
}

func (b *UUID) UnmarshalJSON(data []byte) error {

	var uuid []string

	json.Unmarshal(data,&uuid)

	if len(uuid) != 2 || uuid[0] != "uuid" {
		return fmt.Errorf("wrong number of args to UUID marshal")
	}

	(*b) = UUID(uuid[1])

	return nil
}
