package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/atyronesmith/flowt/pkg/analysis"
	dstructs "github.com/atyronesmith/flowt/pkg/dstructs"
	types "github.com/atyronesmith/flowt/pkg/types"
)

func Parse(in io.Reader) {
	buf := bufio.NewScanner(in)

	buf.Split(bufio.ScanLines)

	lineSplitRE := regexp.MustCompile(`\b,*\s+`)
	fieldSplitRE := regexp.MustCompile(`([^=]+)=(.*)`)
	actionSplitRE := regexp.MustCompile(`([^,\(\)\n]+\(.*?\))|([^,\n]+)`)
	registerRE := regexp.MustCompile(`load:0x([[:xdigit:]])\-\>NXM_NX_REG(\d+)`)
	matchMetadataLoadRE := regexp.MustCompile(`load:0x([[:xdigit:]])\-\>OXM_OF_METADATA`)

	hf := func(rn *types.RuleNode) uint64 {
		return rn.Cookie
	}

	hl := dstructs.NewHashList("Parser", hf)

	lineNo := 1

	for buf.Scan() {
		line := buf.Text()

		fields := lineSplitRE.Split(line, -1)

		if fields[0] == "NXST_FLOW" {
			lineNo++
			continue
		}

		n := types.NewRuleNode()

		for _, v := range fields {
			v = strings.TrimLeft(v, " ")

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
				// OpenFlow  table  0 performs physical-to-logical translation.
				// OpenFlow  tables  8  through  31 execute the logical ingress
				//  pipeline from the Logical_Flow table in the  OVN  Southbound
				//  database.
				// OpenFlow tables 37 through 39 implement the output action in
				//  the logical ingress pipeline.
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
				matches := strings.Split(value, ",")
				priority, err := strconv.ParseInt(matches[0], 10, 32)
				if err != nil {
					fmt.Printf("(%d) Invalid priority: %s, %s\n", lineNo, value, err)
					os.Exit(1)
				}
				n.Priority = uint(priority)
				n.Match = append(n.Match, matches[1:]...)
			} else if key == "actions" {
				actions := actionSplitRE.FindAllString(value, -1)
				n.Actions = append(n.Actions, actions...)

				for _, v := range actions[1:] {
					
					if regMatch := registerRE.FindStringSubmatch(v); len(regMatch) > 0 {
						regVal, err := strconv.ParseInt(regMatch[1], 16, 0)
						if err != nil {
							fmt.Printf("(%d) Unable to parse Reg: %s, %s\n", lineNo, regMatch[1], err)
							os.Exit(1)
						}
						regNum, err := strconv.ParseInt(regMatch[2], 10, 0)
						if err != nil {
							fmt.Printf("(%d) Unable to parse Reg: %s, %s\n", lineNo, regMatch[1], err)
							os.Exit(1)
						}
						switch regNum {
						case 15:
							n.Registers[types.R15] = int(regVal)
						case 14:
							n.Registers[types.R14] = int(regVal)
						case 13:
							n.Registers[types.R13] = int(regVal)
						case 12:
							n.Registers[types.R12] = int(regVal)
						case 11:
							n.Registers[types.R11] = int(regVal)
						case 10:
							n.Registers[types.R10] = int(regVal)
						default:
							fmt.Printf("(%d) Unknown Reg: %d\n", lineNo, regNum)
							os.Exit(1)
						}
					} else if regMatch := matchMetadataLoadRE.FindStringSubmatch(v); len(regMatch) > 0 {
						regVal, err := strconv.ParseInt(regMatch[1], 16, 0)
						if err != nil {
							fmt.Printf("(%d) Unable to parse Reg: %s, %s\n", lineNo, regMatch[1], err)
							os.Exit(1)
						}
						fmt.Printf("metadata %d\n", regVal)
						n.OFMetadata = int(regVal)
					}
				}
			}
		}
		n.Line = uint(lineNo)

		hl.Add(n)

		lineNo++
	}

	//	fmt.Printf("%#v\n", hl)

	fmt.Printf("Lines: %d\n", lineNo)
	tables := analysis.GetTables(hl)
	sort.Slice(tables, func(i, j int) bool {
		return tables[i] < tables[j]
	})
	fmt.Printf("Tables: %v\n", tables)
}
