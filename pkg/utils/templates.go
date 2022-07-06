package utils

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/atyronesmith/flowt/pkg/dbtypes"
)

func ProcessTemplate(templateFile string, name string, funcMap template.FuncMap, data interface{}) (*bytes.Buffer, error) {

	fBuf, err := os.ReadFile(templateFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read template file: %s, %v", templateFile, err)
	}
	tpl, err := template.New(name).Funcs(funcMap).Parse(string(fBuf))
	if err != nil {
		return nil, fmt.Errorf("failed to create template: %s, %v", name, err)
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("error processing template: %s, %v", templateFile, err)
	}

	return &buf, nil
}

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"postfix": func(name string, t string) string {
			return fmt.Sprintf("%s%s", name, t)
		},
		"add": func(a int, b int) int {
			return a + b
		},
		"UUIDToString": func(uuid dbtypes.UUID) string {
			return string(uuid)
		},
		"LookupLogicalSwitchPort": func(lsp map[string]dbtypes.LogicalSwitchPortNB, uuid dbtypes.UUID) dbtypes.LogicalSwitchPortNB {
			return lsp[string(uuid)]
		},
		"LookupLogicalRouterPort": func(lsp map[string]dbtypes.LogicalRouterPortNB, uuid dbtypes.UUID) dbtypes.LogicalRouterPortNB {
			return lsp[string(uuid)]
		},
		"LookupLogicalRouterStaticRoute": func(lsp map[string]dbtypes.LogicalRouterStaticRouteNB, uuid dbtypes.UUID) dbtypes.LogicalRouterStaticRouteNB {
			return lsp[string(uuid)]
		},
		"LookupNAT": func(lsp map[string]dbtypes.NATNB, uuid dbtypes.UUID) dbtypes.NATNB {
			return lsp[string(uuid)]
		},
		"MapSetString": func(str1 string, str2 string) string {
			if len(str1) > 0 {
				return str1
			}
			return str2
		},
		"GenExternalIds": func(key string, value string) string {
			if len(value) == 0 {
				return fmt.Sprintf("external_ids:\\\"%s\\\"=\\\"\\\"",key)				
			}
			return fmt.Sprintf("external_ids:\\\"%s\\\"=\"%s\"",key,value)
		},
		"BuildAddresses": func(addresses []string) string {
			var addrs []string
			for _, v := range addresses {
				addrs = append(addrs, fmt.Sprintf("\"%s\"", v))
			}
			return strings.Join(addrs, " ")
		},
		"BuildMap": func(m map[string]string) string {
			var addrs []string
			for k, v := range m {
				addrs = append(addrs, fmt.Sprintf("\"%s\"=\"%s\"", k, v))
			}
			return strings.Join(addrs, " ")
		},
		"DashToUnder": func(m interface{}) string {
			switch v := m.(type) {
			case dbtypes.UUID:
				return strings.ReplaceAll(v.String(), "-", "_")
			case *dbtypes.UUID:
				return strings.ReplaceAll((*v).String(), "-", "_")
			case string:
				return strings.ReplaceAll(v, "-", "_")
			default:
				return reflect.TypeOf(m).Name()
			}
		},
		"AccessStringSlice": func(s dbtypes.OVSSet[dbtypes.UUID]) string {
			return strings.ReplaceAll(string(s[0]), "-", "_")
		},
	}
}

func WriteByteData(buf *bytes.Buffer, dir string, fileName string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("unable to mkdir: %s", dir)
		}
	}
	dFile, err := os.Create(dir + "/" + fileName)
	if err != nil {
		fmt.Printf("unable to create/open file: %s", fileName)
	}
	defer dFile.Close()

	fmt.Printf("Writing %s...\n", fileName)

	if _, err = dFile.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("unable to write instructions to: %s", dir+"/"+fileName)
	}
	return nil
}
