package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atyronesmith/flowt/pkg/dbparse"
	"github.com/atyronesmith/flowt/pkg/utils"
)

func main() {
	var chartFile string
	var outDir string

	isVerbose := flag.Bool("v", false, "Print extra runtime information.")
	isHelp := flag.Bool("help", false, "Print usage information.")
	flag.StringVar(&chartFile, "chart", "", "Name of chart to generate.")
	flag.StringVar(&chartFile, "c", "", "Name of chart to generate.")
	flag.StringVar(&outDir, "o", ".", "Directory to place the results (Defaults to local directory)")

	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Generate commands to recreate a NB database.\n")
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

	var dbFile string
	if dbFile = flag.Arg(0); dbFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	db, dbSchema, err := dbparse.DBRead(dbFile)
	if err != nil {
		fmt.Printf("Error reading %s: %v", dbFile, err)
		os.Exit(1)
	}

	type tStruct struct {
		Schema dbparse.OVSdbSchema
		Db     *dbparse.OVNDbType
	}

	tPlate := tStruct{
		Schema: *dbSchema,
		Db:     &db,
	}
	tplFile := "templates/gennb.tpl"
	buf, err := utils.ProcessTemplate(tplFile, "generate", utils.GetFuncMap(), &tPlate)
	if err != nil {
		fmt.Printf("unable to process template file: %s, %v", tplFile, err)
		os.Exit(1)
	}

	outBaseName := "/" + strings.ToLower(dbSchema.Type.String())

	bashFle := outDir + outBaseName + ".sh"

	dFile, err := os.Create(bashFle)
	if err != nil {
		fmt.Printf("unable to create/open file: %s", bashFle)
	}
	defer dFile.Close()

	dFile.Write(buf.Bytes())

}

