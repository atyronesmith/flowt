package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atyronesmith/flowt/pkg/dbparse"
	"github.com/atyronesmith/flowt/pkg/dbtypes"
	"github.com/atyronesmith/flowt/pkg/utils"
)

type tStruct struct {
	NBDb        *dbtypes.OVNNorthbound
	SBDb        *dbtypes.OVNSouthbound
	Computes    []string
	Controllers []string
}

type routeDef struct {
	Src string
	Dst string
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
		NBDb: nbDb.(*dbtypes.OVNNorthbound),
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

	err = genVMs(nbDb)
	if err != nil {
		fmt.Printf("Fail: %s\n", err)
		os.Exit(1)
	}
}

func genVMs(db dbparse.OVNDbType) error {

	type vmPort struct {
		Port      string
		PortName  string
		Namespace string
		Mac       string
		IP4       string
		Routes    []routeDef
	}
	var vmPorts []vmPort

	namespaces := make(map[string]string)

	nb := db.(*dbtypes.OVNNorthbound)

	for _, lsp := range nb.LogicalSwitchPort {
		if lsp.Type == nil {
			// This is a VM port

			// Get the Neutron VM Id
			deviceId, ok := lsp.ExternalIds["neutron:device_id"]
			if !ok {
				return fmt.Errorf("missing device_id?")
			}
			// Use the full device id for the namespace
			namespaces[deviceId] = *lsp.Name

			if len(lsp.Addresses) != 1 {
				return fmt.Errorf("multiple or missing port address definitions not supported yet")
			}
			addresses := strings.Split(lsp.Addresses[0], " ")
			if len(addresses) > 2 {
				return fmt.Errorf("multiple port addresses not supported yet")
			}
			vmp := vmPort{}

			vmp.Namespace = deviceId
			vmp.Port = *lsp.Name
			// ovs ports can only be 15 characters
			vmp.PortName = (*lsp.Name)[:15]
			vmp.Mac = addresses[0]
			cidrs, ok := lsp.ExternalIds["neutron:cidrs"]
			if !ok {
				return fmt.Errorf("missing neutron:cidrs")
			}
			vmp.IP4 = cidrs

			if len(lsp.Dhcpv4Options) > 1 {
				return fmt.Errorf("multiple dhcp configuration not yet supported: %v", lsp.Dhcpv4Options)
			}
			if len(lsp.Dhcpv4Options) == 1 {
				dhcpOptions, ok := nb.DHCPOptions[string(lsp.Dhcpv4Options[0])]
				if !ok {
					return fmt.Errorf("missing dhcp4 option: %s", lsp.Dhcpv4Options[0])
				}
				routeString, ok := dhcpOptions.Options["classless_static_route"]
				if !ok {
					return fmt.Errorf("missing dhcp4 option, classless_static_route: %s", lsp.Dhcpv4Options[0])
				}
				routeString = strings.Trim(routeString, "{}")
				// "{169.254.169.254/32,192.168.33.100, 0.0.0.0/0,192.168.33.1}"
				routes := strings.Split(routeString, ", ")
				for _, r := range routes {
					srcDst := strings.Split(r, ",")
					if len(srcDst) != 2 {
						return fmt.Errorf("invalid route: %s", r)
					}
					routeDef := routeDef{
						Src: srcDst[0],
						Dst: srcDst[1],
					}
					vmp.Routes = append(vmp.Routes, routeDef)
				}
			}
			vmPorts = append(vmPorts, vmp)
		}
	}

	type tData struct {
		Ports     []vmPort
		Namespace map[string]string
	}

	data := &tData{
		Ports:     vmPorts,
		Namespace: namespaces,
	}

	tplFile := "templates/gen_vm.tmpl"
	buf, err := utils.ProcessTemplate(tplFile, "vm", utils.GetFuncMap(), &data)
	if err != nil {
		fmt.Printf("Unable to process template file: %s, %v", tplFile, err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", buf.String())

	return nil
}

func processSB(tPlate tStruct, outDir string) {
	sb := tPlate.SBDb

	// Differentiate between controllers and computes
	for _, v := range sb.ChassisPrivate {
		if _, ok := v.ExternalIds["neutron:ovn-metadata-id"]; ok {
			tPlate.Computes = append(tPlate.Computes, v.Chassis[0].String())
		} else {
			tPlate.Controllers = append(tPlate.Controllers, v.Chassis[0].String())
		}
	}

	tplFile := "templates/chassis_params.tmpl"
	buf, err := utils.ProcessTemplate(tplFile, "chassis_params", utils.GetFuncMap(), &tPlate)
	if err != nil {
		fmt.Printf("unable to process template file: %s, %v", tplFile, err)
		os.Exit(1)
	}

	writeData(outDir, buf, outDir+"/inventory_ovn.yaml")
}

func processNB(tPlate tStruct, outDir string) {
	tplFile := "templates/gennb.tmpl"
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
