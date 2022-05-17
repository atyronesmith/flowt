package analysis

import (
	"fmt"
	"os"

	"github.com/atyronesmith/flowt/pkg/dbtypes"
	"github.com/atyronesmith/flowt/pkg/utils"
)


func GenSBStats(db *dbtypes.OVNSouthbound) error {
	fName := "templates/ovnsbstats.tpl"

	buf, err := utils.ProcessTemplate(fName,"chart",utils.GetFuncMap(),db)
	if err != nil {
		return fmt.Errorf("error processing template: %s, %v",fName,err)
	}

	os.Stdout.Write(buf.Bytes())

	return nil
}
