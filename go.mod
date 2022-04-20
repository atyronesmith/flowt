module github.com/atyronesmith/flowt

go 1.18

require github.com/go-echarts/go-echarts/v2 v2.2.5-0.20211021024243-33ae1aa415d6

require (
	github.com/aybabtme/uniplot v0.0.0-20151203143629-039c559e5e7e // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/lpabon/godbc v0.1.1 // indirect
	github.com/ovn-org/libovsdb v0.6.1-0.20220328142833-2cbe2d093e12 // indirect
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
)

// dev mode
//replace (
//	github.com/go-echarts/go-echarts/v2 => ../go-echarts
//)

require (
	github.com/go-echarts/examples v0.0.0-20211021070855-c1dea2301ce9 // indirect
	github.com/heketi/heketi v10.3.0+incompatible
	github.com/aybabtme/uniplot v0.0.0-20151203143629-039c559e5e7e
)

// require (
// 	"github.com/ovn-org/libovsdb/cache" main
// 	"github.com/ovn-org/libovsdb/client" v0.6.0
// 	"github.com/ovn-org/libovsdb/example/vswitchd" v0.6.0
// 	"github.com/ovn-org/libovsdb/model" v0.6.0
// 	"github.com/ovn-org/libovsdb/ovsdb" v0.6.0
// )
