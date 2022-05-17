package remote

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/golang/glog"
	"golang.org/x/crypto/ssh"
)

type DBType string

const (
	NB = "ovn-nbctl"
	SB = "ovn-sbctl"
	L_VS = "ovs-vsctl"
	L_OF = "ovs-ofctl"
)

var DBTypeMap map[string]DBType = map[string]DBType {
	"nb": NB,
	"NB": NB,
	"northdb": NB,
	"north": NB,
	"sb": SB,
	"SB": SB,
	"southdb": SB,
	"south": SB,
	"vs" : L_VS,
	"of" : L_OF,
}

type ExternalIds map[string]string

// TODO
// export SB=$(sudo ovs-vsctl get open . external_ids:ovn-remote | sed -e 's/\"//g')
// export NB=$(sudo ovs-vsctl get open . external_ids:ovn-remote | sed -e 's/\"//g' | sed -e 's/6642/6641/g')
// alias ovn-sbctl='sudo docker exec ovn_controller ovn-sbctl --db=$SB'
// alias ovn-nbctl='sudo docker exec ovn_controller ovn-nbctl --db=$NB'
// alias ovn-trace='sudo docker exec ovn_controller ovn-trace --db=$SB'

var lineSplitRE = regexp.MustCompile(`([^,\"\"\n]+\".*?\")|([^,\n]+)`)

var fieldSplitRE = regexp.MustCompile(`([^=]+)=(.*)`)

var replaceRemoteRE = regexp.MustCompile(`6642`)

func GetSBRemote(externalIds map[string]string) string {
	return externalIds["ovn-remote"]
}

func GetNBRemote(externalIds map[string]string) string {
	return replaceRemoteRE.ReplaceAllString(externalIds["ovn-remote"],"6641")
}

func GetExternalIds(client *ssh.Client) (ExternalIds, error) {
	buf, _, err := SshCommand(client, "sudo ovs-vsctl get open . external_ids")
	if err != nil {
		return nil, fmt.Errorf("SshCommand: %v", err)
	}

	retStr := strings.Trim(buf.String(), "}{ \n")

	x := map[string]string{}

	fields := lineSplitRE.FindAllString(retStr, -1)
	for _, field := range fields {
		field = strings.Trim(field, " ")
		fieldSlice := fieldSplitRE.FindStringSubmatch(field)
		if fieldSlice == nil {
			glog.V(5).Infof("unable to split field: <%s> in external_ids: %s",field, buf.String())

			continue
		}

		x[fieldSlice[1]] = strings.Trim(fieldSlice[2], "\"")
	}

	return x, nil
}

func DumpFlows(client *ssh.Client, bridge string) (*bytes.Buffer, error) {
	cmd := fmt.Sprintf("sudo ovs-ofctl dump-flows %s", bridge)

	buf, _, err := SshCommand(client, cmd)
	if err != nil {
		return nil, fmt.Errorf("SshCommand: %v", err)
	}

	return buf, nil
}

func DumpPorts(client *ssh.Client, bridge string) (*bytes.Buffer, error) {
	cmd := fmt.Sprintf("sudo ovs-ofctl dump-ports %s", bridge)

	buf, _, err := SshCommand(client, cmd)
	if err != nil {
		return nil, fmt.Errorf("SshCommand: %v", err)
	}

	return buf, nil
}

func dbCmd(client *ssh.Client, externalIds ExternalIds, db DBType, cmd string) (*bytes.Buffer, error) {
//ovn-sbctl='sudo docker exec ovn_controller ovn-sbctl --db=$SB

	switch db {
	case NB:
		cmd = fmt.Sprintf("sudo podman exec ovn_controller ovn-nbctl --db=%s %s", GetNBRemote(externalIds), cmd)
	case SB:		
		cmd = fmt.Sprintf("sudo podman exec ovn_controller ovn-sbctl --db=%s %s", GetSBRemote(externalIds), cmd)
	case L_VS:
		cmd = fmt.Sprintf("sudo ovs-vsctl %s", cmd)
	case L_OF:
		cmd = fmt.Sprintf("sudo ovs-ofctl %s", cmd)
	}
	buf, _, err := SshCommand(client, cmd)
	if err != nil {
		return nil, fmt.Errorf("SshCommand: %v", err)
	}

	return buf, nil
}

func GetHelp(client *ssh.Client,externalIds ExternalIds, db DBType) (*bytes.Buffer,error) {
	return dbCmd(client,externalIds,db,"--help")
}

func DumpFlowsSB(client *ssh.Client,externalIds ExternalIds) (*bytes.Buffer, error) {
//ovn-sbctl='sudo docker exec ovn_controller ovn-sbctl --db=$SB
	return dbCmd(client,externalIds,SB,"dump-flows")
}

func DumpFlowsLS(client *ssh.Client,externalIds ExternalIds) (*bytes.Buffer, error) {
//ovn-sbctl='sudo docker exec ovn_controller ovn-sbctl --db=$SB
	return dbCmd(client,externalIds,NB,"ls-list")
}

func RunCmd(client *ssh.Client, externalIds ExternalIds, cmd string, db DBType) (*bytes.Buffer, error) {
	return dbCmd(client,externalIds,db,cmd)
}

func GetDBFile(client *ssh.Client, externalIds ExternalIds,db DBType) (buf *bytes.Buffer, dbFile string, err error) {
	switch db {
	case NB:
		dbFile = "ovnnb_db.db"
	case SB:		
		dbFile = "ovnsb_db.db"
	}

	cmd := fmt.Sprintf("sudo cat /var/lib/openvswitch/ovn/%s", dbFile)

	buf, _, err = SshCommand(client, cmd)
	if err != nil {
		return nil, dbFile, fmt.Errorf("SshCommand: %v", err)
	}

	return buf, dbFile, nil
}