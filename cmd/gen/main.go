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

const (
	OVN_NB_GEN_SCRIPT   = "ovn_nb_net.sh"
	OVN_NB_INVENTORY    = "inventory_ovn.yaml"
	OVN_NB_VM_SCRIPT    = "ovn_nb_vm.sh"
	NORTHBOUND_NET_TMPL = "templates/gennb.tmpl"
	NORTHBOUND_VM_TMPL  = "templates/gen_vm.tmpl"
	CHASSIS_PARAM_TMPL  = "templates/chassis_params.tmpl"
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

type vmPort struct {
	Port      string
	PortName  string
	Namespace string
	Mac       string
	IP4       string
	Hostname  string
	Routes    []routeDef
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

	err = processNB(tPlate, outDir)
	if err != nil {
		fmt.Printf("Error generating NB create commands: %v", err)
		os.Exit(1)
	}

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

	err = processSB(tPlate, outDir)
	if err != nil {
		fmt.Printf("Error generating SB create commands: %v", err)
		os.Exit(1)
	}

	err = genVMs(nbDb, sbDb, outDir)
	if err != nil {
		fmt.Printf("Error generating VM port commands: %v\n", err)
		os.Exit(1)
	}
}

func genVMs(nbDB, sbDB dbparse.OVNDbType, outDir string) error {
	vmPorts := make(map[string][]vmPort, 1)

	namespaces := make(map[string]string)

	nb := nbDB.(*dbtypes.OVNNorthbound)
	sb := sbDB.(*dbtypes.OVNSouthbound)

	for _, lsp := range nb.LogicalSwitchPort {
		if lsp.Type == nil {
			// This is a VM port

			// Lookup a binding for this port in the SB db
			var portBinding *dbtypes.PortBindingSB
			for _, pb := range sb.PortBinding {
				if *pb.LogicalPort == *lsp.Name {
					pbCopy := pb
					portBinding = &pbCopy
					break
				}
			}
			if portBinding == nil {
				return fmt.Errorf("no port binding for LSP: %s", *lsp.Name)
			}
			chassis, ok := sb.Chassis[string(portBinding.Chassis[0])]
			if !ok {
				return fmt.Errorf("no chassis for port binding: %v", *portBinding)
			}

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

			vmp.Hostname = *chassis.Hostname + "-ovs"
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
			// if _, ok := vmPorts[vmp.Hostname]; !ok {
			// 	vmPorts[vmp.Hostname] = make([]vmPort, 1)
			// }
			vmPorts[vmp.Hostname] = append(vmPorts[vmp.Hostname], vmp)
		}
	}

	type tData struct {
		Ports     []vmPort
		Namespace map[string]string
	}

	for hostname, p := range vmPorts {

		data := &tData{
			Ports:     p,
			Namespace: namespaces,
		}

		buf, err := utils.ProcessTemplate(NORTHBOUND_VM_TMPL, "vm", utils.GetFuncMap(), &data)
		if err != nil {
			fmt.Printf("Unable to process template file: %s, %v", NORTHBOUND_VM_TMPL, err)
			os.Exit(1)
		}
//		fmt.Printf("%s\n", buf.String())

		writeData(outDir, buf, outDir, hostname)
	}
	return nil
}

func processSB(tPlate tStruct, outDir string) error {
	sb := tPlate.SBDb

	// Differentiate between controllers and computes
	for _, v := range sb.ChassisPrivate {
		if _, ok := v.ExternalIds["neutron:ovn-metadata-id"]; ok {
			tPlate.Computes = append(tPlate.Computes, v.Chassis[0].String())
		} else {
			tPlate.Controllers = append(tPlate.Controllers, v.Chassis[0].String())
		}
	}

	buf, err := utils.ProcessTemplate(CHASSIS_PARAM_TMPL, "chassis_params", utils.GetFuncMap(), &tPlate)
	if err != nil {
		fmt.Printf("unable to process template file: %s, %v", CHASSIS_PARAM_TMPL, err)
		os.Exit(1)
	}

	return writeData(outDir, buf, outDir, OVN_NB_INVENTORY)
}

func processNB(tPlate tStruct, outDir string) error {
	buf, err := utils.ProcessTemplate(NORTHBOUND_NET_TMPL, "nb_net", utils.GetFuncMap(), &tPlate)
	if err != nil {
		fmt.Printf("Unable to process template file: %s, %v", NORTHBOUND_NET_TMPL, err)
		os.Exit(1)
	}

	return writeData(outDir, buf, outDir, OVN_NB_GEN_SCRIPT)
}

func writeData(outDir string, buf *bytes.Buffer, dir string, fileName string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("unable to mkdir: %s", dir)
		}
	}
	dFile, err := os.Create(dir + "/" + fileName)
	if err != nil {
		fmt.Printf("unable to create/open file: %s", fileName)
	}
	defer dFile.Close()

	fmt.Printf("Writing %s...\n", fileName)

	if _, err = dFile.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("unable to write instructions to: %s", dir+"/"+fileName)
	}
	return nil
}
