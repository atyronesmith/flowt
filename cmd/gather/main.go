package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/atyronesmith/flowt/pkg/dbparse"
	"github.com/atyronesmith/flowt/pkg/remote"
	"github.com/atyronesmith/flowt/pkg/utils"
	"golang.org/x/exp/slices"
)

var commands = []string{"osp", "scale"}

func main() {
	var getNB, getSB bool
	var optVerbose, optHelp bool
	var host string
	var outDir string
	var sshKey string

	flag.CommandLine.BoolVar(&optVerbose, "verbose", false, "Print extra runtime information.")
	flag.BoolVar(&optHelp, "help", false, "Print usage information.")
	flag.BoolVar(&getNB, "nb", true, "Gather the NB (northbound) database")
	flag.BoolVar(&getSB, "sb", true, "Gather the SB (southbound) database")
	flag.StringVar(&outDir, "outDir", ".", "Directory to place the results (Defaults to local directory)")
	flag.StringVar(&outDir, "o", ".", "Directory to place the results (Defaults to local directory)")
	flag.StringVar(&sshKey, "sshKey", "", "Location of ssh private key to use (Defaults to ~/.ssh/id_rsa)")
	flag.StringVar(&sshKey, "s", "", "Location of ssh private key to use (Defaults to ~/.ssh/id_rsa)")

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

	command := flag.Arg(0)
	if !slices.Contains(commands, command) {
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}

	host = flag.Arg(1)

	var nbBuf, sbBuf *bytes.Buffer

	var ssh *remote.Ssh
	var err error

	if command == "osp" {
		ssh, err = remote.NewSsh("heat-admin",sshKey)
		if err != nil {
			fmt.Printf("Error creating ssh config, %v", err)
			os.Exit(1)
		}
	} else if command == "scale" {
		ssh, err = remote.NewSsh("root",sshKey)
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
		if getNB {
			nbBuf, err = remote.GetDBOSP(client, dbparse.NB)
			if err != nil {
				fmt.Printf("%v", err)
				os.Exit(1)
			}
		}
		if getSB {
			sbBuf, err = remote.GetDBOSP(client, dbparse.SB)
			if err != nil {
				fmt.Printf("%v", err)
				os.Exit(1)
			}
		}
	} else if command == "scale" {
		if getNB {
			nbBuf, err = remote.GetDBPod(client, dbparse.NB)
			if err != nil {
				fmt.Printf("%v", err)
				os.Exit(1)
			}
		}
		if getNB {
			sbBuf, err = remote.GetDBPod(client, dbparse.SB)
			if err != nil {
				fmt.Printf("%v", err)
				os.Exit(1)
			}
		}
	}

	if nbBuf != nil {
		utils.WriteByteData(nbBuf, outDir, "ovnnb_db.db")
	}
	if sbBuf != nil {
		utils.WriteByteData(sbBuf, outDir, "ovnsb_db.db")
	}
}
