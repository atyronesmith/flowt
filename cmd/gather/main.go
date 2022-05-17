package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atyronesmith/flowt/pkg/remote"
)

func main() {
	var dbase string

	isVerbose := flag.Bool("verbose", false, "Print extra runtime information.")
	isHelp := flag.Bool("help", false, "Print usage information.")
	flag.StringVar(&dbase,"db","NB", "Use NB (northbound) | SB (southbound) | L_OF (local ofctl) | L_VS (local vsctl).")

	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Usage: %s [options] host [cmd ..]\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(CommandLine.Output(), "       host  -- Host to gather flow rules.\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *isVerbose {
		fmt.Println("Verbose...")
	}

	if *isHelp || flag.NArg() < 1 {
		flag.Usage()

		os.Exit(0)
	}

	host := flag.Arg(0)

	if host == "" {
		flag.Usage()
		os.Exit(1)
	}

	ssh, err := remote.NewSsh()
	if err != nil {
		fmt.Printf("Error creating ssh config, %v", err)
		os.Exit(1)
	}

	client, err := ssh.ConnectSSH(host)
	if err != nil {
		fmt.Printf("Error connectiong to host: %s, %s\n", host, err)

		os.Exit(1)
	}
	defer client.Close()

	externalIds, err := remote.GetExternalIds(client)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	var buf *bytes.Buffer

	if flag.NArg() > 1 {
		db, ok := remote.DBTypeMap[dbase]
		if !ok {
			flag.Usage()
			os.Exit(1)
		}
		cmdStrings := flag.Args()
		cmd := strings.Join(cmdStrings[1:], " ")
		buf, err = remote.RunCmd(client, externalIds, cmd, db)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", buf)
	} else {
		db, ok := remote.DBTypeMap[dbase]
		if !ok {
			flag.Usage()
			os.Exit(1)
		}
		buf, dbFile, err := remote.GetDBFile(client,externalIds,db)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		os.WriteFile(dbFile, buf.Bytes(), 0644)
	}
}
