package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/atyronesmith/flowt/pkg/dbparse"
	"github.com/atyronesmith/flowt/pkg/dbtypes"
	"github.com/atyronesmith/flowt/pkg/utils"
)

type tStruct struct {
	Db          *dbparse.OVNDbType
	SBDb        *dbtypes.OVNSouthbound
	Computes    []string
	Controllers []string
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
		fmt.Fprintf(CommandLine.Output(), "Generate commands to recreate a NB database.\n")
		fmt.Fprintf(CommandLine.Output(), "Usage: %s [options] northbound_db [southbound_db]\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(CommandLine.Output(), "       northbound_db    -- Path to northbound database.\n")
		fmt.Fprintf(CommandLine.Output(), "       [southbound_db]  -- Optional path to SB database.\n")
		fmt.Fprintf(CommandLine.Output(), "                           Providing a SB database enables generation of chassis information.\n")
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

	var nbDbFile string
	if nbDbFile = flag.Arg(0); nbDbFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	nbDb, nbDbSchema, err := dbparse.DBRead(nbDbFile)
	if err != nil {
		fmt.Printf("Error reading %s: %v", nbDbFile, err)
		os.Exit(1)
	}

	if nbDbSchema.Type != dbparse.NB {
		fmt.Printf("Not a NB database: %s", nbDbFile)
		os.Exit(1)
	}

	tPlate := tStruct{
		Db: &nbDb,
	}

	processNB(tPlate, outDir)

	var sbDbFile string
	if sbDbFile = flag.Arg(1); sbDbFile == "" {
		os.Exit(1)
	}

	sbDb, sbDbSchema, err := dbparse.DBRead(sbDbFile)
	if err != nil {
		fmt.Printf("Error reading %s: %v", sbDbFile, err)
		os.Exit(1)
	}

	if sbDbSchema.Type != dbparse.SB {
		fmt.Printf("Not a SB database: %s", sbDbFile)
		os.Exit(1)
	}

	tPlate.SBDb = sbDb.(*dbtypes.OVNSouthbound)

	processSB(tPlate, outDir)

}

func processSB(tPlate tStruct, outDir string) {
	sb := tPlate.SBDb

	// Differentiate between controllers and computes
	for _, v := range sb.ChassisPrivate {
		if _, ok := v.ExternalIds["neutron:ovn-metadata-id"];ok {
				tPlate.Computes = append(tPlate.Computes, v.Chassis[0].String())
		} else {
				tPlate.Controllers = append(tPlate.Controllers, v.Chassis[0].String())
		}
	}

	tplFile := "templates/chassis-params.tpl"
	buf, err := utils.ProcessTemplate(tplFile, "chassis-params", utils.GetFuncMap(), &tPlate)
	if err != nil {
		fmt.Printf("unable to process template file: %s, %v", tplFile, err)
		os.Exit(1)
	}

	writeData(outDir, buf, outDir+"/ovn.yaml")
}

func processNB(tPlate tStruct, outDir string) {
	tplFile := "templates/gennb.tpl"
	buf, err := utils.ProcessTemplate(tplFile, "generate", utils.GetFuncMap(), &tPlate)
	if err != nil {
		fmt.Printf("Unable to process template file: %s, %v", tplFile, err)
		os.Exit(1)
	}

	writeData(outDir, buf, outDir+"/ovn_northbound.sh")
}

func writeData(outDir string, buf *bytes.Buffer, fileName string) {
	dFile, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("unable to create/open file: %s", fileName)
	}
	defer dFile.Close()

	fmt.Printf("Writing %s...\n", fileName)

	dFile.Write(buf.Bytes())
}
