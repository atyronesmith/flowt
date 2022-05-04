package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/atyronesmith/flowt/pkg/analysis"
	"github.com/atyronesmith/flowt/pkg/dbparse"
	"github.com/atyronesmith/flowt/pkg/dbtypes"
)

func main() {
	var chartFile string

	isVerbose := flag.Bool("verbose", false, "Print extra runtime information.")
	isHelp := flag.Bool("help", false, "Print usage information.")
	flag.StringVar(&chartFile, "chart", "", "Name of chart to generate.")
	flag.StringVar(&chartFile, "c", "", "Name of chart to generate.")

	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Usage: %s [options] file_to_parse\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(CommandLine.Output(), "       file_to_parse  -- Path to flow rules.\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *isVerbose {
		fmt.Println("Verbose...")
	}

	if *isHelp {
		flag.Usage()

		os.Exit(0)
	}

	var filename string
	if filename = flag.Arg(0); filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	db, dbSchema, err := dbparse.DBRead(filename)
	if err != nil {
		fmt.Printf("Error reading %s: %v", filename, err)
		os.Exit(1)
	}

	switch dbSchema.Type {
	case dbtypes.NB:
		analysis.GenNBStats(db.(*dbtypes.OVNNorthbound))
	case dbtypes.SB:
	}

	if ( len(chartFile) > 0 ) {
		if err = genChart(db,chartFile); err != nil {
			fmt.Printf("%v\n",err)
			os.Exit(1)
		}
	}
}

func genChart(db dbtypes.OVNDbType, chartFile string) error {
	cFile, err := os.Create(chartFile)
	if err != nil {
		return fmt.Errorf("unable to create/open file: %s", chartFile)
	}

	funcMap := template.FuncMap{
		"add": func(a int, b int) int {
			fmt.Printf("%d %d\n", a, b)
			return a + b
		},
	}

	tplFilename := "templates/lsw.tpl"
	fBuf, err := os.ReadFile(tplFilename)
	if err != nil {
		return fmt.Errorf("unable to read template file: %s",tplFilename)
	}
	tpl, err := template.New("LogicalSwitch").Funcs(funcMap).Parse(string(fBuf))
	if err != nil {
		return fmt.Errorf("error parsing %s: %v", tplFilename, err)
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, db); err != nil {
		return fmt.Errorf("error executing template: %s, %v", tplFilename, err)
	}
	cFile.Write(buf.Bytes())

	return nil
}
