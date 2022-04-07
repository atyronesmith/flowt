package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/atyronesmith/flowt/pkg/remote"
)

// TODO
// export SB=$(sudo ovs-vsctl get open . external_ids:ovn-remote | sed -e 's/\"//g')
// export NB=$(sudo ovs-vsctl get open . external_ids:ovn-remote | sed -e 's/\"//g' | sed -e 's/6642/6641/g')
// alias ovn-sbctl='sudo docker exec ovn_controller ovn-sbctl --db=$SB'
// alias ovn-nbctl='sudo docker exec ovn_controller ovn-nbctl --db=$NB'
// alias ovn-trace='sudo docker exec ovn_controller ovn-trace --db=$SB'

func main() {
	isVerbose := flag.Bool("verbose", false, "Print extra runtime information.")
	isHelp := flag.Bool("help", false, "Print usage information.")

	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Usage: %s [options] host\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(CommandLine.Output(), "       host  -- Host to gather flow rules.\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *isVerbose {
		fmt.Println("Verbose...")
	}

	if *isHelp {
		flag.Usage()

		os.Exit(0)
	}

	host := flag.Arg(0)

	if host == "" {
		flag.Usage()
		os.Exit(1)
	}

	ssh := remote.NewSsh()

	client, err := ssh.ConnectSSH(host)
	if err != nil {
		fmt.Printf("Error connectiong to host: %s, %s\n",host,err)

		os.Exit(1)
	}
	defer client.Close()

	

	ret, err := remote.GetExternalIds(client)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	fmt.Printf("Results: %#v\n", ret)
}
