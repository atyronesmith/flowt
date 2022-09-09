package dbparse

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/InVisionApp/conjungo"
	jsoniter "github.com/json-iterator/go"
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
	jsonObj := regexp.MustCompile(`^{`)
    isSchema := regexp.MustCompile(`"cksum":`)

	lineNo := 1

	dd, err := ovsSchema.NewDb()
	if err != nil {
		return nil, nil, err
	}

	found := false
	for scanner.Scan() {
		line := scanner.Text()

		if jsonObj.MatchString(line) {
			if isSchema.MatchString(line) {
				continue
			}
			ddDelta, err := ovsSchema.NewDb()
			if err != nil {
				return nil, nil, err
			}
			if err := jsoniter.Unmarshal([]byte(line), &ddDelta); err != nil {
				return nil, nil, fmt.Errorf("error while unmarshalling, lineno %d: %v\n\n%100s", lineNo, err,line)
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
