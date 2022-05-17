package main

import (
	"flag"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"

	dbparse "github.com/atyronesmith/flowt/pkg/dbparse"
	"github.com/atyronesmith/flowt/pkg/utils"
)

func main() {
	var isVerbose, isHelp bool
	var pkg, outFile string

	flag.BoolVar(&isVerbose, "v", false, "Print extra runtime information.")
	flag.BoolVar(&isHelp, "h", false, "Print usage information.")
	flag.StringVar(&outFile, "o", "", "Go Schema file output file (Defaults to name in schema in current directory)")
	flag.StringVar(&pkg, "pkg", "dbtypes", "Target package for schema.")

	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Generate an ./ovn_nortbound.go or ./ovn_southbound.go file depending on\n")
		fmt.Fprintf(CommandLine.Output(), "the type of database contained in the file_to_parse.  The ***.go file\n")
		fmt.Fprintf(CommandLine.Output(), "contains go structures that map the database data contents.\n\n")
		fmt.Fprintf(CommandLine.Output(), "Usage: %s [options] file_to_parse\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(CommandLine.Output(), "       file_to_parse  -- Path to flow rules.\n")

		flag.PrintDefaults()
	}

	flag.Parse()

	if isVerbose {
		fmt.Println("Verbose...")
	}

	if isHelp || flag.Arg(0) == "" {
		flag.Usage()

		os.Exit(0)
	}

	ovsSchema := dbparse.OVSdbSchema{}

	if err := ovsSchema.ReadOvsSchema(flag.Arg(0)); err != nil {
		fmt.Printf("Error while reading db info: %s, %v\n", flag.Arg(0), err)
		os.Exit(1)
	}

	schemaGoName := strings.ToLower(ovsSchema.Type.String()) + ".go"
	if len(outFile) > 0 {
		schemaGoName = outFile
	}

	oFile, err := os.Create(schemaGoName)
	if err != nil {
		fmt.Printf("unable to create/open file: %s", schemaGoName)
		os.Exit(1)
	}
	defer oFile.Close()

	var dbDef dbparse.DbDef
	err = dbDef.ParseSchema(&ovsSchema)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	type tStruct struct {
		PkgName string
		Schema  dbparse.OVSdbSchema
		DBDef   dbparse.DbDef
	}

	tPlate := tStruct{
		PkgName: pkg,
		Schema:  ovsSchema,
		DBDef:   dbDef,
	}

	buf, err := utils.ProcessTemplate("templates/dbschema.tpl", "dbschema", utils.GetFuncMap(), &tPlate)
	if err != nil {
		fmt.Printf("unable to process template file: %s, %v", "templates/dbschema.tpl", err)
		os.Exit(1)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("Error in gofmt: %s", err)
		os.Exit(1)
	}
	oFile.Write(p)
}
