package remote

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/ssh"
)

func GetExternalIds(client *ssh.Client) (*map[string]string, error) {
	buf, err := SshCommand(client, []string{"sudo ovs-vsctl -d json get open . external_ids"})
	if err != nil {
		fmt.Printf("%v", err)

		return nil, err
	}
	// Not JSON

	retStr := strings.Trim(buf.String(), "{}")

	fmt.Printf("%s\n", retStr)

	x := map[string]string{}

	lineSplitRE := regexp.MustCompile(`\b,*\s+`)
	fieldSplitRE := regexp.MustCompile(`([^=]+)=(.*)`)

	fields := lineSplitRE.Split(retStr, -1)
	for _, field := range fields {
		fieldSlice := fieldSplitRE.FindStringSubmatch(field)
		if fieldSlice == nil {
			return &x, fmt.Errorf("unable to split field: <%s> in external_ids: %s", field, buf.String())
		}
		x[fieldSlice[1]] = fieldSlice[2]
	}

	return &x, nil
}
