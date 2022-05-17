package utils

import (
	"flag"
)

type FlagVar interface {
	bool | int | string
}

type FlagDef struct {
	Var     interface{}
	Name    string
	Default interface{}
	Help    string
}

func (fd *FlagDef) Populate() error {
	switch v := fd.Var.(type) {
	case bool:
		flag.BoolVar(&v, fd.Name, fd.Default.(bool), fd.Help)
		flag.BoolVar(&v, fd.Name[0:1], fd.Default.(bool), fd.Help)
	case int:
		flag.IntVar(&v, fd.Name, fd.Default.(int), fd.Help)
		flag.IntVar(&v, fd.Name[0:1], fd.Default.(int), fd.Help)
	}

	return nil
}
