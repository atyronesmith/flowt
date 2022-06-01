package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	"github.com/atyronesmith/flowt/pkg/analysis"
	"github.com/atyronesmith/flowt/pkg/dbparse"
	"github.com/atyronesmith/flowt/pkg/dbtypes"
	"github.com/atyronesmith/flowt/pkg/utils"
)

func mapMerge(data interface{}, delta interface{}) (bool, error) {

	deltaType := reflect.ValueOf(delta)
	dataType := reflect.ValueOf(data)

	assign := false

	if deltaType.Kind() != dataType.Kind() {
		return assign, fmt.Errorf("objects not the same")
	}

	switch deltaType.Kind() {
	case reflect.Map:
		for k, v := range delta.(map[string]interface{}) {
			d := data.(map[string]interface{})
			// Check to see if data already has an entry
			// If it does, continue down the layers
			// If it does not, add the new delta data to data
			if dk, ok := d[k]; ok {
				a, err := mapMerge(dk, v)
				if err != nil {
					return false, err
				}
				if a {
					d[k] = v
				}
			} else {
				d[k] = v
			}
		}
	case reflect.Slice:
		for _, v := range data.([]interface{}) {
			mapMerge(data, v)
		}
	default:
		// Only atomic values at this point
		assign = true
	}

	return assign, nil
}

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
		fmt.Fprintf(CommandLine.Output(), "Read an OVN NB or SB database, generate a json file of just the data,\n")
		fmt.Fprintf(CommandLine.Output(), "and generate a .dot file that represents the schema of the database.\n\n")

		fmt.Fprintf(CommandLine.Output(), "Usage: %s [options] db_to_parse\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(CommandLine.Output(), "       db_to_parse  -- Path to NB or SB database.\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *isHelp {
		flag.Usage()

		os.Exit(0)
	}

	var dbFilename string
	if dbFilename = flag.Arg(0); dbFilename == "" {
		flag.Usage()
		os.Exit(1)
	}

	db, dbSchema, err := dbparse.DBRead(dbFilename)
	if err != nil {
		fmt.Printf("Error reading %s: %v", dbFilename, err)
		os.Exit(1)
	}

	if *isVerbose {
		switch dbSchema.Type {
		case dbparse.NB:
			analysis.GenNBStats(db.(*dbtypes.OVNNorthbound))
		case dbparse.SB:
			analysis.GenSBStats(db.(*dbtypes.OVNSouthbound))
		}
	}
	
	if len(chartFile) > 0 && dbSchema.Type == dbparse.NB {
		if err = genChart(db, chartFile); err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}

	var dbDef dbparse.DbDef
	err = dbDef.ParseSchema(dbSchema)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	err = dbDef.AugmentSchema()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	type tStruct struct {
		Schema dbparse.OVSdbSchema
		Db     *dbparse.OVNDbType
		DbDef  *dbparse.DbDef
	}

	tPlate := tStruct{
		Schema: *dbSchema,
		Db:     &db,
		DbDef:  &dbDef,
	}

	buf, err := utils.ProcessTemplate("templates/schema_dot.tmpl", "dotschema", utils.GetFuncMap(), &tPlate)
	if err != nil {
		fmt.Printf("unable to process template file: %s, %v", "templates/schema_dot.tmpl", err)
		os.Exit(1)
	}

	outBaseName := "/" + strings.ToLower(dbSchema.Type.String())

	dotFile := outDir + outBaseName + ".dot"
	dFile, err := os.Create(dotFile)
	if err != nil {
		fmt.Printf("unable to create/open file: %s", dotFile)
	}
	defer dFile.Close()

	dFile.Write(buf.Bytes())

	pretty, _ := json.MarshalIndent(db, "", "    ")
	os.WriteFile(outDir+outBaseName+".json", pretty, 0644)
}

func genChart(db dbparse.OVNDbType, chartFile string) error {
	cFile, err := os.Create(chartFile)
	if err != nil {
		return fmt.Errorf("unable to create/open file: %s", chartFile)
	}

	funcMap := template.FuncMap{
		"add": func(a int, b int) int {
			return a + b
		},
	}

	tplFilename := "templates/lsw.tmpl"
	fBuf, err := os.ReadFile(tplFilename)
	if err != nil {
		return fmt.Errorf("unable to read template file: %s", tplFilename)
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
