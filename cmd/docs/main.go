package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/atyronesmith/flowt/pkg/utils"
)

func main() {
	var chartFile string
	var outDir string

	isHelp := flag.Bool("help", false, "Print usage information.")
	flag.StringVar(&chartFile, "chart", "", "Name of chart to generate.")
	flag.StringVar(&chartFile, "c", "", "Name of chart to generate.")
	flag.StringVar(&outDir, "o", ".", "Directory to place the results (Defaults to local directory)")

	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Read XML version of OVN db schema docs and generate tooltips.\n\n")
		fmt.Fprintf(CommandLine.Output(), "Usage: %s [options] db_to_parse > cmd/db/data/ovn_*.json\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(CommandLine.Output(), "       db_to_parse  -- Path to NB or SB database.\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *isHelp {
		flag.Usage()

		os.Exit(0)
	}

	var xmlFilename string
	if xmlFilename = flag.Arg(0); xmlFilename == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Open our xmlFile
	xmlFile, err := os.Open(xmlFilename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Printf("Failed to open: %s\n", xmlFilename)
		os.Exit(1)
	}
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	stringValue := string(byteValue)

	var reRef = regexp.MustCompile(`(?m)<ref (?:.*?)((?:db|table|column)=\"([^"]+)\")+/>`)

	var substitution = "&lt;B&gt;$2&lt;/B&gt;"

	stringValue = reRef.ReplaceAllString(stringValue, substitution)

	var reCode = regexp.MustCompile(`(?m)<code>/s*([^<]+)/s*</code>`)

	substitution = "&lt;B&gt;$1&lt;/B&gt;"
	stringValue = reCode.ReplaceAllString(stringValue, substitution)

	byteValue = []byte(stringValue)

	// we initialize our Users array
	var dbase utils.Database
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &dbase)

	var reSentenceSpace = regexp.MustCompile(`(?m)\.\s{2,}`)
	var reMoreSpace = regexp.MustCompile(`(?m)([\w,])\s{2,}([\w<])`)

	fmt.Printf("{\n")

	count := 0
	for _, tbl := range dbase.Table {
		if count > 0 {
			fmt.Printf(",\n")
		}
		fmt.Printf("    \"%s\": ", tbl.Name)
		var sb strings.Builder
		for _, p := range tbl.P {
			sb.WriteString(p.Text)
		}
		txt := reSentenceSpace.ReplaceAllString(sb.String(), ".  ")
		txt = reMoreSpace.ReplaceAllString(txt, "$1 $2")
		txt = strings.Replace(txt, "\n", "", -1)
		txt = strings.Replace(txt, "\"", "", -1)
		fmt.Printf(" \"%s\"", strings.TrimSpace(strings.Replace(txt, "\n", "", -1)))
		count++
	}
	fmt.Printf("\n}\n")

	// ind,err := xml.MarshalIndent(dbase,"","    ")
	// fmt.Printf("%s\n",string(ind))
}
