package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"text/template"

	dbparse "github.com/atyronesmith/flowt/pkg/dbparse"
	dbtypes "github.com/atyronesmith/flowt/pkg/dbtypes"
)

func main() {
	isVerbose := flag.Bool("verbose", false, "Print extra runtime information.")
	isHelp := flag.Bool("help", false, "Print usage information.")
	hasPackage := flag.String("pkg", "dbtypes", "Target package for schema.")

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

	if *isHelp || flag.Arg(0) == "" {
		flag.Usage()

		os.Exit(0)
	}

	ovsSchema := dbtypes.OVSdbSchema{}

	if err := ovsSchema.OvsHeader(flag.Arg(0)); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	dbDef, err := dbparse.ParseSchema(ovsSchema, *hasPackage)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	type tStruct struct {
		PkgName string
		Schema  dbtypes.OVSdbSchema
		DBDef   dbtypes.DbDef
	}

	tPlate := tStruct{
		PkgName: *hasPackage,
		Schema:  ovsSchema,
		DBDef:   *dbDef,
	}

	funcMap := template.FuncMap{
		"postfix": func(name string, t string) string {
			return fmt.Sprintf("%s%s", name, t)
		},
	}

	fBuf, err := os.ReadFile("templates/dbschema.tpl")
	if err != nil {
		fmt.Printf("Unable to read template file: %s", "templates/dbschema.tpl")
		os.Exit(1)
	}
	tpl, err := template.New("dbschema").Funcs(funcMap).Parse(string(fBuf))
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, tPlate); err != nil {
		fmt.Printf("Error processing template: %s", err)
		os.Exit(1)
	}
//	os.Stdout.Write(buf.Bytes())
	p, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("Error in gofmt: %s", err)
		os.Exit(1)
	}
	os.Stdout.Write(p)
}
