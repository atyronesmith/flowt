package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/atyronesmith/flowt/pkg/analysis"
	"github.com/atyronesmith/flowt/pkg/dbparse"
	"github.com/atyronesmith/flowt/pkg/dbtypes"
	"github.com/atyronesmith/flowt/pkg/utils"
)

type tStruct struct {
	Schema dbparse.OVSdbSchema
	Db     *dbparse.OVNDbType
	DbDef  *dbparse.DbDef
}

func main() {
	var chartFile string
	var outDir string
	var jsonData bool
	var dotSchema bool
	var dataPath string

	isVerbose := flag.Bool("v", false, "Print extra runtime information.")
	isHelp := flag.Bool("help", false, "Print usage information.")
	flag.StringVar(&chartFile, "chart", "", "Name of chart to generate.")
	flag.StringVar(&chartFile, "c", "", "Name of chart to generate.")
	flag.StringVar(&outDir, "outDir", ".", "Directory to place the results (Defaults to local directory)")
	flag.StringVar(&outDir, "o", ".", "Directory to place the results (Defaults to local directory)")
	flag.BoolVar(&jsonData, "jsonData", false, "Generage json file with DB data.")
	flag.BoolVar(&jsonData, "jd", false, "Generage json file with DB data.")
	flag.BoolVar(&dotSchema, "dotSchema", false, "Generage dot (GraphViz) file with DB schema.")
	flag.BoolVar(&dotSchema, "ds", false, "Generage dot (GraphViz) file with DB schema.")
	flag.StringVar(&dataPath, "datapath", "", "Generate datapath info for --datapath <neutron_name>|*")
	flag.StringVar(&dataPath, "dp", "", "Generate datapath info for --dp <neutron_name>|*")

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
		fmt.Printf("chart = <%s>\n", chartFile)
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

	tPlate := tStruct{
		Schema: *dbSchema,
		Db:     &db,
		DbDef:  &dbDef,
	}

	outBaseName := strings.ToLower(dbSchema.Type.String())

	if dotSchema {
		genDotSchema(tPlate, dbSchema, outDir, outBaseName)
	}

	if jsonData {
		genJsonData(db, outDir, outBaseName)
	}

	if len(dataPath) > 0 {
		genDatapath(dbSchema, db, dataPath)
	}
}

type ByTableId []dbtypes.LogicalFlowSB

func (a ByTableId) Len() int { return len(a) }
func (a ByTableId) Less(i, j int) bool {
	if a[i].TableId == a[j].TableId {
		return a[i].Priority > a[j].Priority
	}
	return a[i].TableId < a[j].TableId
}
func (a ByTableId) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func genDatapath(dbSchema *dbparse.OVSdbSchema, db dbparse.OVNDbType, datapath string) error {

	if dbSchema.Type != dbparse.SB {
		return fmt.Errorf("input database must be of type SouthBound to generate datapath information")
	}

	sb := db.(*dbtypes.OVNSouthbound)

	dpRE := regexp.MustCompile(datapath)

	for dataPathKey, dataPathValue := range sb.DatapathBinding {
		if dpRE.MatchString(dataPathValue.ExternalIds["name2"]) {
			var lf []dbtypes.LogicalFlowSB

			for _, flowValue := range sb.LogicalFlow {
				if flowValue.LogicalDatapath.String() == dataPathKey {
					lf = append(lf, flowValue)
				}
			}
			sort.Sort(ByTableId(lf))
			for _, a := range lf {
				if a.Pipeline == "ingress" {
					buf, _ := json.MarshalIndent(a, "", "    ")
					fmt.Printf("%s\n", buf)
				}
			}

		}
	}

	return nil
}

func genJsonData(db dbparse.OVNDbType, outDir string, outBaseName string) {
	pretty, _ := json.MarshalIndent(db, "", "    ")

	utils.WriteByteData(bytes.NewBuffer(pretty), outDir, outBaseName+".json")
}

func genDotSchema(tPlate tStruct, dbSchema *dbparse.OVSdbSchema, outDir string, outBaseName string) {
	buf, err := utils.ProcessTemplate("templates/schema_dot.tmpl", "dotschema", utils.GetFuncMap(), &tPlate)
	if err != nil {
		fmt.Printf("unable to process template file: %s, %v", "templates/schema_dot.tmpl", err)
		os.Exit(1)
	}

	utils.WriteByteData(buf, outDir, outBaseName+".dot")
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
