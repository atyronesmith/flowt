package analysis

import (
	"fmt"
	"os"

	"github.com/atyronesmith/flowt/pkg/dbtypes"
	"github.com/atyronesmith/flowt/pkg/utils"
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

func GenNBStats(db *dbtypes.OVNNorthbound) error {
	fName := "templates/ovnnbstats.tmpl"

	type tStruct struct {
		Stats NBStats
		Db    *dbtypes.OVNNorthbound
	}

	nbs := nbStats(db)
	tPlate := tStruct{
		Stats: nbs,
		Db:    db,
	}

	buf, err := utils.ProcessTemplate(fName, "chart", utils.GetFuncMap(), tPlate)
	if err != nil {
		return fmt.Errorf("error processing template: %s, %v", fName, err)
	}

	os.Stdout.Write(buf.Bytes())

	return nil
}
