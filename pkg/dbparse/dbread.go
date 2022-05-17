package dbparse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/InVisionApp/conjungo"
)

var jsonObj = regexp.MustCompile(`^{(.*)}\s*$`)

func DBReadData(filename string) (map[string]interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	in := f
	stats, err := os.Stat(filename)
	if err != nil {
		return nil, fmt.Errorf("error occured on file: %s, %v", filename, err)
	}

	scanner := bufio.NewScanner(in)

	scanner.Buffer(make([]byte, 0), int(stats.Size()))
	scanner.Split(bufio.ScanLines)

	// db file must be unformatted

	lineNo := 1
	jsonObjCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if m := jsonObj.FindString(line); len(m) > 0 {
			if jsonObjCount > 0 {
				dataMap := make(map[string]interface{})
				if err := json.Unmarshal([]byte(m), &dataMap); err != nil {
					return nil, fmt.Errorf("error while unmarshalling(%d): %v", lineNo, err)
				}
				return dataMap, nil
			}
			jsonObjCount++
		}
		lineNo += 1
	}

	return nil, fmt.Errorf("could not find data in: %s", filename)
}

func DBRead(filename string) (OVNDbType, *OVSdbSchema, error) {
	ovsSchema := OVSdbSchema{}

	if err := ovsSchema.ReadOvsSchema(filename); err != nil {
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

	
	dd, err := ovsSchema.NewDb()
	if err != nil {
		return nil, nil, err
	}

	found := false
	for scanner.Scan() {
		line := scanner.Text()

		if m := jsonObj.FindString(line); len(m) > 0 {
			ddDelta, err := ovsSchema.NewDb()
			if err != nil {
				return nil, nil, err
			}

			if err := json.Unmarshal([]byte(m), &ddDelta); err != nil {
				return nil, nil, fmt.Errorf("error while unmarshalling(%d): %v", lineNo, err)
			}
			if err := conjungo.Merge(dd, ddDelta, nil); err != nil {
				fmt.Printf("Merge error: %v\n",err)
				os.Exit(1)
			}
			if dd.IsValid() {
				found = true
			}
		}
		lineNo += 1
	}

	if found {
		return dd, &ovsSchema, nil
	} else {
		return nil, nil, fmt.Errorf("no valid data in %s", filename)
	}
}
