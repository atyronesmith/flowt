package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	types "github.com/atyronesmith/flowt/pkg/types"
)

func Parse(in io.Reader) {
	buf := bufio.NewScanner(in)

	buf.Split(bufio.ScanLines)

	lineSplitRE := regexp.MustCompile(`\b,*\s+`)
	fieldSplitRE := regexp.MustCompile(`([^=]+)=(.*)`)
	actionSplitRE := regexp.MustCompile(`([^,\(\)\n]+\(.*?\))|([^,\n]+)`)

	nodes := make([]types.RuleNode, 0, 100)

	lineNo := 1

	for buf.Scan() {
		line := buf.Text()

		fields := lineSplitRE.Split(line, -1)

		if fields[0] == "NXST_FLOW" {
			lineNo++
			continue;
		}

		n := types.RuleNode{}

		for _, v := range fields {
			fieldSlice := fieldSplitRE.FindStringSubmatch(v)
			if fieldSlice == nil {
				fmt.Printf("(%d) Unable to split field: <%s> in fields: %s\n\t\n", lineNo, v, fields[3])
				os.Exit(1)
			}
			key := fieldSlice[1]
			value := fieldSlice[2]

			if key == "cookie" {
				cookie, err := strconv.ParseUint(value, 0, 64)
				if err != nil {
					fmt.Printf("Invalid cookie: %s, %s\n", value, err)
					os.Exit(1)
				}
				n.Cookie = cookie
			} else if key == "duration" {
				val := strings.Replace(value, "s", "", 1)
				duration, err := strconv.ParseFloat(val, 64)
				if err != nil {
					fmt.Printf("Invalid duration: %s, %s\n", val, err)
					os.Exit(1)
				}
				n.Duration = duration
			} else if key == "table" {
				table, err := strconv.ParseInt(value, 10, 32)
				if err != nil {
					fmt.Printf("Invalid table: %s, %s\n", value, err)
					os.Exit(1)
				}
				n.Table = uint32(table)
			} else if key == "n_packets" {
				nPkts, err := strconv.ParseInt(fieldSlice[2], 10, 64)
				if err != nil {
					fmt.Printf("Invalid n_packets: %s, %s\n", value, err)
					os.Exit(1)
				}
				n.NPackets = uint64(nPkts)
			} else if key == "n_bytes" {
				nBytes, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					fmt.Printf("Invalid n_bytes: %s, %s\n", value, err)
					os.Exit(1)
				}
				n.NBytes = uint64(nBytes)
			} else if key == "priority" {
				matches := strings.Split(value,",")
				priority, err := strconv.ParseInt(matches[0], 10, 32)
				if err != nil {
					fmt.Printf("(%d) Invalid priority: %s, %s\n", lineNo, value, err)
					os.Exit(1)
				}
				n.Priority = uint(priority)
				n.Match = append(n.Match, matches[1:]...)
			} else if key == "actions" {
				actions := actionSplitRE.FindAllString(value,-1)
				n.Actions = append(n.Actions, actions...)
			}
		}
		n.Line = uint(lineNo)

		nodes = append(nodes, n)

		lineNo++
	}

	for _, n := range nodes {
		fmt.Printf("%+v\n",n)
	}

}
