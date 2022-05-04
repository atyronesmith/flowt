package analysis

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/atyronesmith/flowt/pkg/dbtypes"
)

type NBStats struct {
	NumVMPorts        int
	NumRouterPorts    int
	NumLocalPorts     int
	NumLocalNetPorts  int
	NumL2GatewayPorts int
	NumVTEPPorts      int
	NumExternalPorts  int
	NumVirtualPorts   int
	NumRemotePorts    int
	NAT               map[string]int
}

func nbStats(nb *dbtypes.OVNNorthbound) NBStats {
	stats := NBStats{}
	stats.NAT = make(map[string]int)

	for _, v := range nb.LogicalSwitchPort {
		switch v.Type {
		case "":
			stats.NumVMPorts++
		case "router":
			stats.NumRouterPorts++
		case "localport":
			stats.NumLocalPorts++
		case "localnet":
			stats.NumLocalNetPorts++
		case "l2gateway":
			stats.NumL2GatewayPorts++
		case "vtep":
			stats.NumVTEPPorts++
		case "external":
			stats.NumExternalPorts++
		case "virtual":
			stats.NumVirtualPorts++
		case "remote":
			stats.NumRemotePorts++
		}
	}

	for _, v := range nb.NAT {
		stats.NAT[v.Type]++
	}

	return stats
}

func GenNBStats(db *dbtypes.OVNNorthbound) {
	fName := "templates/ovnnbstats.tpl"

	fBuf, err := os.ReadFile(fName)
	if err != nil {
		fmt.Printf("Unable to read template file: %s", fName)
		os.Exit(1)
	}
	funcMap := template.FuncMap{
		"add": func(a int, b int) int {
			fmt.Printf("%d %d\n", a, b)
			return a + b
		},
	}

	tpl, err := template.New("dbschema").Funcs(funcMap).Parse(string(fBuf))
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	type tStruct struct {
		Stats NBStats
		Db    dbtypes.OVNDbType
	}

	nbs := nbStats(db)
	tPlate := tStruct{
		Stats: nbs,
		Db:    db,
	}
	var buf bytes.Buffer

	if err := tpl.Execute(&buf, tPlate); err != nil {
		fmt.Printf("Error processing template: %s", err)
		os.Exit(1)
	}
	os.Stdout.Write(buf.Bytes())

}
