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

type Parser struct {
	Cookies *dstructs.HashList[uint64,types.RuleNode]
	Tables  *dstructs.HashList[int,types.RuleNode]

	LineCount int
}

func Parse(in io.Reader) (*Parser, error) {
	buf := bufio.NewScanner(in)

	buf.Split(bufio.ScanLines)

	lineSplitRE := regexp.MustCompile(`\b,*\s+`)
	fieldSplitRE := regexp.MustCompile(`([^=]+)=(.*)`)
	actionSplitRE := regexp.MustCompile(`([^,\(\)\n]+\(.*?\))|([^,\n]+)`)
	registerRE := regexp.MustCompile(`load:0x([[:xdigit:]])\-\>NXM_NX_REG(\d+)`)
	matchMetadataLoadRE := regexp.MustCompile(`load:0x([[:xdigit:]])\-\>OXM_OF_METADATA`)
	matchResubmit := regexp.MustCompile(`resubmit\(,(\d+)\)`)
	
	hf := func(rn *types.RuleNode) uint64 {
		return rn.Cookie
	}

	of := func(rn1 *types.RuleNode, rn2 *types.RuleNode) int {
		if rn1.Table < rn2.Table {
			return -1
		}

		if rn1.Table == rn2.Table {
			return -(int(rn1.Priority) - int(rn2.Priority))
		}

		return 1
	}
	hl := dstructs.NewHashList("Parser", hf, of)

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
				return nil, fmt.Errorf("(%d) unable to split field: <%s> in fields: %s", lineNo, v, fields[3])
			}
			key := fieldSlice[1]
			value := fieldSlice[2]

			if key == "cookie" {
				cookie, err := strconv.ParseUint(value, 0, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid cookie: %s, %v", value, err)
				}
				n.Cookie = cookie
			} else if key == "duration" {
				val := strings.Replace(value, "s", "", 1)
				duration, err := strconv.ParseFloat(val, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid duration: %s, %v", val, err)
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
					return nil, fmt.Errorf("invalid table: %s, %v", value, err)
				}
				n.Table = uint32(table)
			} else if key == "n_packets" {
				nPkts, err := strconv.ParseInt(fieldSlice[2], 10, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid n_packets: %s, %v", value, err)
				}
				n.NPackets = uint64(nPkts)
			} else if key == "n_bytes" {
				nBytes, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("nnvalid n_bytes: %s, %v", value, err)
				}
				n.NBytes = uint64(nBytes)
			} else if key == "priority" {
				matches := strings.Split(value, ",")
				priority, err := strconv.ParseInt(matches[0], 10, 32)
				if err != nil {
					return nil, fmt.Errorf("(%d) Invalid priority: %s, %v", lineNo, value, err)
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
							return nil, fmt.Errorf("(%d) unable to parse Reg: %s, %v", lineNo, regMatch[1], err)
						}
						regNum, err := strconv.ParseInt(regMatch[2], 10, 0)
						if err != nil {
							return nil, fmt.Errorf("(%d) nnable to parse Reg: %s, %v", lineNo, regMatch[1], err)
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
							return nil, fmt.Errorf("(%d) unable to parse Reg: %s, %v", lineNo, regMatch[1], err)
						}
						n.OFMetadata = int(regVal)
					} else if regMatch := matchResubmit.FindStringSubmatch(v); len(regMatch) > 0 {
						regVal, err := strconv.ParseInt(regMatch[1], 16, 0)
						if err != nil {
							return nil, fmt.Errorf("(%d) unable to parse Resubmit(,N): %s, %v", lineNo, regMatch[1], err)
						}
						n.NextTables = append(n.NextTables,int(regVal))
					} else if v == "drop" {
						n.Drop = true
					}
				}
			}
		}
		n.Line = uint(lineNo)

		if err := hl.Add(n); err != nil {
			return nil, fmt.Errorf("node insert failed: %v", err)
		}

		lineNo++
	}

	fmt.Printf("Lines: %d\n", lineNo)
	tables := analysis.GetTables(hl)
	sort.Slice(tables, func(i, j int) bool {
		return tables[i] < tables[j]
	})
	fmt.Printf("Tables: %v\n", tables)

	cookies := analysis.GetCookieCounts(hl)
	fmt.Printf("Cookies: %d\n",len(cookies))

	cookieCount := make([]int,len(cookies))
	var index int
	for _, v := range cookies {
//		fmt.Printf("%d\n",v)
		cookieCount[index] = v
		index++
	}
	hist := analysis.CalcHist(cookieCount,5)

	fmt.Printf("Max: %d, Min: %d\n", hist.Max, hist.Min)
	analysis.PrintHist(hist)

	parser := Parser {
		LineCount: lineNo,
		Cookies: hl,
	}
	return &parser, nil
}
