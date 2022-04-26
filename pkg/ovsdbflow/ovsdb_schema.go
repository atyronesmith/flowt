package ovsdbflow

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
)

type UUID string

type OVSdbSchema struct {
	ChkSum  string                 `json:"chksum"`
	Type    OVSDBType              `json:"name"`
	Version string                 `json:"version"`
	Tables  map[string]interface{} `json:"tables"`
}

func (schema *OVSdbSchema) OvsHeader(filename string) error {
	var in io.Reader

	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %s, %v", filename, err)
	}
	defer f.Close()

	in = f

	stats, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("error occured on file: %s, %v", filename, err)
	}

	scanner := bufio.NewScanner(in)

	scanner.Buffer(make([]byte, 0), int(stats.Size()))
	scanner.Split(bufio.ScanLines)

	// db file must be unformatted
	jsonObj := regexp.MustCompile(`^{.*}\s*$`)

	for scanner.Scan() {
		line := scanner.Text()

		if jsonObj.MatchString(line) {
			if err := json.Unmarshal([]byte(line), schema); err != nil {
				return fmt.Errorf("unable to unmarshall json string %v", err)
			}
			if schema.Type == Unknown {
				return fmt.Errorf("unknown ovsdb type %s", schema.Type)
			}
			return nil
		}
	}

	return fmt.Errorf("error, could not find ovsdb header: %s", filename)
}
