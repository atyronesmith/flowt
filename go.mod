module github.com/atyronesmith/flowt

go 1.18

require (
	github.com/aybabtme/uniplot v0.0.0-20151203143629-039c559e5e7e
	github.com/go-echarts/go-echarts/v2 v2.2.5-0.20211021024243-33ae1aa415d6
	github.com/golang/glog v1.0.0
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29
)

require (
	github.com/InVisionApp/conjungo v1.1.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/exp v0.0.0-20220613132600-b0d781184e0d // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
)

// dev mode
//replace (
//	github.com/go-echarts/go-echarts/v2 => ../go-echarts
//)

require github.com/go-echarts/examples v0.0.0-20211021070855-c1dea2301ce9
