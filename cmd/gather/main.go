package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/atyronesmith/flowt/pkg/remote"
	"golang.org/x/exp/slices"
)

var commands = []string{"osp", "scale"}

func main() {
	var optDbase string
	var optVerbose, optHelp bool
	var host string

	flag.CommandLine.BoolVar(&optVerbose, "verbose", false, "Print extra runtime information.")
	flag.BoolVar(&optHelp, "help", false, "Print usage information.")
	flag.StringVar(&optDbase, "db", "NB", "Use NB (northbound) | SB (southbound)")

	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Usage: %s [options] osp <host> | scale <host>\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(CommandLine.Output(), "       osp   host -- Retrieve DB from OSP Controller host.\n")
		fmt.Fprintf(CommandLine.Output(), "       scale host -- Retrieve DB from Scale Controller host.\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if optVerbose {
		fmt.Println("Verbose...")
	}

	if optHelp || flag.NArg() < 2 {
		flag.Usage()

		os.Exit(0)
	}

	db, ok := remote.DBTypeMap[optDbase]
	if !ok {
		flag.Usage()
		os.Exit(1)
	}

	command := flag.Arg(0)
	if !slices.Contains(commands, command) {
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}

	host = flag.Arg(1)

	var buf *bytes.Buffer

	var ssh *remote.Ssh
	var err error

	if command == "osp" {
		ssh, err = remote.NewSsh("heat-admin")
		if err != nil {
			fmt.Printf("Error creating ssh config, %v", err)
			os.Exit(1)
		}
	} else if command == "scale" {
		ssh, err = remote.NewSsh("root")
		if err != nil {
			fmt.Printf("Error creating ssh config, %v", err)
			os.Exit(1)
		}
	}
	client, err := ssh.ConnectSSH(host)
	if err != nil {
		fmt.Printf("Error connectiong to host: %s, %s\n", host, err)

		os.Exit(1)
	}
	defer client.Close()

	if command == "osp" {
		buf, err = remote.GetDBOSP(client, db)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
	} else if command == "scale" {
		buf, err = remote.GetDBPod(client, db)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

	}

	if buf != nil {
		fmt.Printf("%s\n", buf.String())
	}
}
