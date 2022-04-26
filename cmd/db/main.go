package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	types "github.com/atyronesmith/flowt/pkg/ovsdbflow"
	"github.com/atyronesmith/flowt/pkg/ovsdbflow/nb"
	utils "github.com/atyronesmith/flowt/pkg/utils"
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

	var in io.Reader

	var filename string
	if filename = flag.Arg(0); filename == "" {
		flag.Usage()
		os.Exit(1)

	}
	ovsSchema := types.OVSdbSchema{}

	if err := ovsSchema.OvsHeader(filename); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1)
	}
	defer f.Close()

	in = f
	stats, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("error occured on file: %s, %v", filename, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(in)

	scanner.Buffer(make([]byte, 0), int(stats.Size()))
	scanner.Split(bufio.ScanLines)

	// OVSDB JSON 867056 d03a2cb121de101e6f9a1906751f54f05cd80606
	// {
	// ...
	// }

	// db file must be unformatted
	jsonObj := regexp.MustCompile(`^{(.*)}\s*$`)

	lineNo := 1
	//	data := types.OVSdb{}
	data := nb.OVNNorthbound{}

	for scanner.Scan() {
		line := scanner.Text()

		if m := jsonObj.FindString(line); len(m) > 0 {
			fmt.Printf("Unmarshal line %d\n", lineNo)
			if err := json.Unmarshal([]byte(m), &data); err != nil {
				fmt.Printf("Error while Unmarshalling: %v", err)
				break
			}
			if len(data.LogicalSwitchPort) > 0 {
				fmt.Println(utils.PrettyStruct(data))

				break
			}
		}
		lineNo += 1
	}

}
