package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/pprof"
	"sort"
	"time"

	"github.com/atyronesmith/flowt/pkg/dbparse"
	"github.com/atyronesmith/flowt/pkg/dbtypes"
	"github.com/atyronesmith/flowt/pkg/utils"
)

// type tStruct struct {
// 	Schema dbparse.OVSdbSchema
// 	Db     *dbparse.OVNDbType
// 	DbDef  *dbparse.DbDef
// }

func main() {
	var outDir string

	doCpuprofile := flag.Bool("cpuprof", false, "write cpu profile to `file`")
	doMemprofile := flag.Bool("heapprof", false, "write memory profile to `file`")
	isVerbose := flag.Bool("v", false, "Print extra runtime information.")
	isHelp := flag.Bool("help", false, "Print usage information.")
	flag.StringVar(&outDir, "outDir", ".", "Directory to place the results (Defaults to local directory)")
	flag.StringVar(&outDir, "o", ".", "Directory to place the results (Defaults to local directory)")

	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Read an OVN NB and SB databases, generate a datapath flow diagram.\n")

		fmt.Fprintf(CommandLine.Output(), "Usage: %s [options] nb_db sb_db [datapath]\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(CommandLine.Output(), "       nb_db,sb_db  -- Paths to NB and SB databases.\n")
		fmt.Fprintf(CommandLine.Output(), "       datapath  -- name of network to analyze (Defaults to providing a list of datapath).\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *isHelp {
		flag.Usage()

		os.Exit(0)
	}

	argCount := len(flag.Args())
	if argCount < 2 || argCount > 3 {
		flag.Usage()
		os.Exit(1)
	}
	sbdbFilename := flag.Arg(1)

	startTime := time.Now()

	if *doCpuprofile  {
		f, _ := os.Create("cpu.pprof")
		defer f.Close()
		err := pprof.StartCPUProfile(f)
		if err != nil {
			fmt.Printf("Unable to start CPU profile: %s\n", err)
		}
	}

	if *isVerbose {
		fmt.Printf("Read SB DB..")
	}
	sbdb, sbdbSchema, err := dbparse.DBRead(sbdbFilename)
	if err != nil {
		fmt.Printf("\nError reading %s: %v", sbdbFilename, err)
		os.Exit(1)
	}
	endTime := time.Now()
	
	if *isVerbose {
		diffTime := endTime.Sub(startTime)
		fmt.Printf("Seconds: %f\n", diffTime.Seconds())
	}

	if *doCpuprofile {
		pprof.StopCPUProfile()
	}

	if *doMemprofile {
       f, _ := os.Create("mem.pprof")
        defer f.Close() // error handling omitted for example
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
	}

	if *doCpuprofile || *doMemprofile {
		os.Exit(0)
	}

	if sbdbSchema.Type != dbparse.SB {
		fmt.Printf("\nInvalid SB DB.  Not of type Southbound, %s", sbdbFilename)
		os.Exit(1)
	}

	if argCount == 2 {
		listDatapath(sbdb)

		os.Exit(0)
	}

	nbdbFilename := flag.Arg(0)

	if *isVerbose {
		fmt.Printf("Read NB DB..")
	}
	nbdb, nbdbSchema, err := dbparse.DBRead(nbdbFilename)
	if err != nil {
		fmt.Printf("Error reading %s: %v", nbdbFilename, err)
		os.Exit(1)
	}
	if *isVerbose {
		fmt.Printf("\n")
	}

	if nbdbSchema.Type != dbparse.NB {
		fmt.Printf("Invalid NB DB.  Not of type Northbound, %s", nbdbFilename)
		os.Exit(1)
	}

	dataPath := flag.Arg(2)

	genDatapath(nbdb, sbdb, dataPath, outDir, dataPath+".dot")
}

func listDatapath(db dbparse.OVNDbType) {
	sb := db.(*dbtypes.OVNSouthbound)

	for _, dataPathValue := range sb.DatapathBinding {
		dpType := "unknown"
		var uuid string
		uuid, ok := dataPathValue.ExternalIds["logical-router"]
		if ok {
			dpType = "logical-router"
		} else {
			uuid, ok = dataPathValue.ExternalIds["logical-switch"]
			if ok {
				dpType = "logical-switch"
			}
		}
		fmt.Printf("%s (%s %s)\n", dataPathValue.ExternalIds["name2"], dpType, uuid)
	}
}

type ByTableIdPriority []dbtypes.LogicalFlowSB

func (a ByTableIdPriority) Len() int { return len(a) }
func (a ByTableIdPriority) Less(i, j int) bool {
	if *a[i].TableId == *a[j].TableId {
		return *a[i].Priority > *a[j].Priority
	}
	return *a[i].TableId < *a[j].TableId
}
func (a ByTableIdPriority) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func genDatapath(nbdb dbparse.OVNDbType, sbdb dbparse.OVNDbType, datapath string, outDir string, filename string) error {

	sb := sbdb.(*dbtypes.OVNSouthbound)
	//nb := sbdb.(*dbtypes.OVNNorthbound)

	dpRE := regexp.MustCompile(datapath)

	// Search through DatapathBinding looking for net names that match the supplied Regex
	for dataPathKey, dataPathValue := range sb.DatapathBinding {
		if dpRE.MatchString(dataPathValue.ExternalIds["name2"]) {
			// Found the network

			// Search through the logical flows to find flows that reference the selected datapath
			// Flows can either have a reference to a single datapath (.LogicalDatapath)
			// Or a reference to a datapath group that contains multiple datapaths
			lfIngress, portsIngress := extractFlows(sb, dataPathKey, "ingress")

			flowTableIngress := buildFlowTable(lfIngress)

			// Now have a sorted list of flows belonging to the specified datapath(s)
			type tStruct struct {
				Ports     []dbtypes.PortBindingSB
				FlowTable [][]dbtypes.LogicalFlowSB
			}

			ts := tStruct{
				FlowTable: flowTableIngress,
				Ports:     portsIngress,
			}
			buf, err := utils.ProcessTemplate("templates/datapath.tmpl", "datapath", utils.GetFuncMap(), &ts)
			if err != nil {
				fmt.Printf("unable to process template file: %s, %v", "templates/schema_dot.tmpl", err)
				os.Exit(1)
			}
			utils.WriteByteData(buf, outDir, filename)

			// for _, a := range lf {
			// 	encoder := json.NewEncoder(os.Stdout)
			// 	encoder.SetEscapeHTML(false)
			// 	encoder.SetIndent("", "    ")
			// 	encoder.Encode(a)
			// }
			// for _, a := range ports {
			// 	encoder := json.NewEncoder(os.Stdout)
			// 	encoder.SetEscapeHTML(false)
			// 	encoder.SetIndent("", "    ")
			// 	encoder.Encode(a)
			// }
		}
	}

	return nil
}

func buildFlowTable(logicalFlow []dbtypes.LogicalFlowSB) [][]dbtypes.LogicalFlowSB {

	sort.Sort(ByTableIdPriority(logicalFlow))

	var flowTable [][]dbtypes.LogicalFlowSB

	currTableId := *logicalFlow[0].TableId
	currPriority := *logicalFlow[0].Priority

	var currArray []dbtypes.LogicalFlowSB

	for _, v := range logicalFlow {
		if *v.TableId != currTableId || *v.Priority != currPriority {
			flowTable = append(flowTable, currArray)

			currArray = nil

			currTableId = *v.TableId
			currPriority = *v.Priority
		}

		currArray = append(currArray, v)
	}
	return flowTable
}

func extractFlows(sb *dbtypes.OVNSouthbound, dataPathKey string, pipeline string) ([]dbtypes.LogicalFlowSB, []dbtypes.PortBindingSB) {
	var lf []dbtypes.LogicalFlowSB
	var ports []dbtypes.PortBindingSB
	//	portMap := make(map[string]string)

	for flowKey, flowValue := range sb.LogicalFlow {
		var appendFlow bool

		if flowValue.LogicalDatapath == nil {
			if flowValue.LogicalDpGroup == nil {
				fmt.Printf("Logical flow with no datapath binding: %s\n", flowKey)
			} else {
				lDPG, ok := sb.LogicalDPGroup[flowValue.LogicalDpGroup.String()]
				if !ok {
					fmt.Printf("Missing Logical Flow Group: %s\n", flowValue.LogicalDpGroup.String())
				} else {
					for _, lg := range lDPG.Datapaths {
						if lg.String() == dataPathKey {
							appendFlow = true
						}
					}
				}
			}
		} else {
			if flowValue.LogicalDatapath.String() == dataPathKey {
				appendFlow = true
			}
		}
		if appendFlow && *flowValue.Pipeline == pipeline {
			if flowValue.ExternalIds == nil {
				flowValue.ExternalIds = dbtypes.OVSMapString{}
			}
			flowValue.ExternalIds["_uuid"] = flowKey
			lf = append(lf, flowValue)
			// tag, ok := flowValue.Tags["in_out_port"]
			// if ok {
			// 	for pKey, pValue := range sb.PortBinding {
			// 		if pValue.LogicalPort != nil && *pValue.LogicalPort == tag {
			// 			_, ok := portMap[tag]
			// 			if !ok {
			// 				if pValue.ExternalIds == nil {
			// 					pValue.ExternalIds = dbtypes.OVSMap[string]{}
			// 				}
			// 				pValue.ExternalIds["_uuid"] = pKey
			// 				portMap[tag] = pKey
			// 				ports = append(ports, pValue)
			// 			}
			// 		}
			// 	}
			// }
		}
	}
	return lf, ports
}
