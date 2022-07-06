package analysis

import (
	"fmt"

	"github.com/atyronesmith/flowt/pkg/dbtypes"
	"github.com/atyronesmith/flowt/pkg/utils"
)


func GenSBStats(db *dbtypes.OVNSouthbound, outDir string, filename string) error {
	fName := "templates/ovnsbstats.tmpl"

	buf, err := utils.ProcessTemplate(fName,"chart",utils.GetFuncMap(),db)
	if err != nil {
		return fmt.Errorf("error processing template: %s, %v",fName,err)
	}

	utils.WriteByteData(buf, outDir, filename)

	return nil
}
