package dbparse

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/InVisionApp/conjungo"
)

var jsonObj = regexp.MustCompile(`^{(.*)}\s*$`)

func DBRead(filename string) (OVNDbType, *OVSdbSchema, error) {
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

	ovsSchema := OVSdbSchema{}

	if err := ovsSchema.ReadOvsDbSchemaFile(filename); err != nil {
		return nil, nil, fmt.Errorf("%v", err)
	}
	return DBReadWithSchema(in, int(stats.Size()), &ovsSchema)
}

func DBReadBuf(buf bytes.Buffer,ovsSchema *OVSdbSchema) (OVNDbType,  error) {
	rdr := bytes.NewReader(buf.Bytes())

	dbType, _, err := DBReadWithSchema(rdr, rdr.Len(), ovsSchema)

	return dbType, err
}

func DBReadWithSchema(in io.Reader, maxTokenSize int, ovsSchema *OVSdbSchema) (OVNDbType, *OVSdbSchema, error) {
	scanner := bufio.NewScanner(in)

	scanner.Buffer(make([]byte, 0), maxTokenSize)
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
				fmt.Printf("Merge error: %v\n", err)
				os.Exit(1)
			}
			if dd.IsValid() {
				found = true
			}
		}
		lineNo += 1
	}

	if found {
		return dd, ovsSchema, nil
	} else {
		return nil, nil, fmt.Errorf("no valid data")
	}
}
