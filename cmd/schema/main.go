package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	ovsdb "github.com/atyronesmith/flowt/pkg/ovsdbflow"
	"github.com/atyronesmith/flowt/pkg/schema"
)

func main() {
	isVerbose := flag.Bool("verbose", false, "Print extra runtime information.")
	isHelp := flag.Bool("help", false, "Print usage information.")

	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Usage: %s [options] file_to_parse\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(CommandLine.Output(), "       file_to_parse  -- Path to flow rules.\n")
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

	ovsSchema := ovsdb.OVSdbSchema{}

	if err := ovsSchema.OvsHeader(flag.Arg(0)); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("// Type: %s\n", ovsSchema.Type)
	fmt.Printf("// Version: %s\n", ovsSchema.Version)
	fmt.Printf("// Table Definitions: %d\n", len(ovsSchema.Tables))
	for k := range ovsSchema.Tables {
		fmt.Printf("//\t%s\n", k)
	}
	fmt.Println()
	if err := schema.ParseSchema(ovsSchema); err != nil {
		fmt.Printf("Error:")
	}

}
