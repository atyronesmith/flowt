package dbparse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/atyronesmith/flowt/pkg/dbtypes"
)

func DBRead(filename string) (dbtypes.OVNDbType, *dbtypes.OVSdbSchema, error) {
	ovsSchema := dbtypes.OVSdbSchema{}

	if err := ovsSchema.OvsHeader(filename); err != nil {
		return nil, nil, fmt.Errorf("%v", err)
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	in := f
	stats, err := os.Stat(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error occured on file: %s, %v", filename, err)
	}

	scanner := bufio.NewScanner(in)

	scanner.Buffer(make([]byte, 0), int(stats.Size()))
	scanner.Split(bufio.ScanLines)

	// db file must be unformatted
	jsonObj := regexp.MustCompile(`^{(.*)}\s*$`)

	lineNo := 1

	for scanner.Scan() {
		line := scanner.Text()

		if m := jsonObj.FindString(line); len(m) > 0 {
			dd, err := ovsSchema.NewDb()
			if err != nil {
				return nil, nil, err
			}
			if err := json.Unmarshal([]byte(m), &dd); err != nil {
				return nil, nil, fmt.Errorf("error while unmarshalling(%d): %v", lineNo, err)
			}
			if dd.IsValid() {
				return dd, &ovsSchema, nil
			}
		}
		lineNo += 1
	}

	return nil, nil, fmt.Errorf("no valid data in %s", filename)
}
