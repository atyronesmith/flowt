package analysis

import (
	"os"

	dstructs "github.com/atyronesmith/flowt/pkg/dstructs"
	"github.com/atyronesmith/flowt/pkg/types"
	uniplot "github.com/aybabtme/uniplot/histogram"
)

// podman ps -q | xargs podman stats --no-stream
// ID            NAME                                                          CPU %   MEM USAGE / LIMIT  MEM %   NET IO   BLOCK IO           PIDS
// b7aa2d4cb2be  nova_virtlogd                                                 7.94%   4.067MB / 134.8GB  0.00%   -- / --  -- / --            2
// 08dac5ffcdeb  nova_libvirt                                                  0.00%   28.16MB / 134.8GB  0.02%   -- / --  0B / 108.3MB       63
// 685c9e778a15  iscsid                                                        7.98%   3.146MB / 134.8GB  0.00%   -- / --  -- / --            2
// e66ae472f328  logrotate_crond                                               2.07%   34.81MB / 134.8GB  0.03%   -- / --  12.29kB / 830.5MB  2
// 649232d351cf  nova_migration_target                                         0.00%   2.63MB / 134.8GB   0.00%   -- / --  -- / --            2
// 77e6e9f4c2d5  ovn_controller                                                2.78%   6.623MB / 134.8GB  0.00%   -- / --  -- / --            5
// 649164df335b  ovn_metadata_agent                                            5.79%   290.6MB / 134.8GB  0.22%   -- / --  0B / 2.261MB       53
// c9668115d977  nova_compute                                                  0.00%   373.2MB / 134.8GB  0.28%   -- / --  29.7kB / 37.26MB   143
// 9b3d6dbe2f0a  neutron-haproxy-ovnmeta-95164d20-1b84-4bad-8763-88a6e3ca70cc  1.53%   2.699MB / 134.8GB  0.00%   -- / --  -- / --            3


func GetTables[K comparable](hl *dstructs.HashList[K, types.RuleNode]) []uint32 {

	tables := make(map[uint32]int)

	nodes := hl.GetNodes()
	for _, v := range nodes {
		tables[(*v).Table] = 1
	}
	t := make([]uint32, 0)

	for k := range tables {
		t = append(t, k)
	}

	return t
}

func GetCookieCounts[K comparable](hl *dstructs.HashList[K, types.RuleNode]) map[uint64]int {

	cookies := make(map[uint64]int)

	nodes := hl.GetNodes()
	for _, v := range nodes {
		if (*v).Cookie != 0 {
			cookies[(*v).Cookie] += 1
		}
	}

	return cookies
}

func CalcHist[K int|int32|int64|uint64](arr []K,bins int) uniplot.Histogram {
	arrLen := len(arr)

	histArr := make([]float64,arrLen)

	for index, v := range arr {
		histArr[index] = float64(v)
	}

	return uniplot.Hist(bins,histArr)
}

func PrintHist(hist uniplot.Histogram) {
	uniplot.Fprint(os.Stdout,hist,uniplot.Linear(50))
}