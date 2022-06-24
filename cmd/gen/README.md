# gen

This command generates a file containing ovn-nbctl commands to recreate parts of an emulated OVN network running on the **ovn-scale-tests** test framework.  The command requires the OVN Northbound and Southbound databases from the network to be tested as inputs.  The databases are text files in the RFC7047 format.  Currently, only OpenStack network implementations are supported.  As an example, the ovnnb_db.db and ovnsb_db.db files can be found in the **/var/lib/openvswitch/ovn** directory on an OSP Controller.

The command generates two bash scripts and an Ansible inventory file that are used to recreate the database.  The bash scripts are run in the context of the*nortd*container that is created by the*ovn-scale-tests* ansible playbook.  The first script, **ovn_northbound.sh** recreates the logical network structures contained in the target OVN databases to be emulated.  The second script file generates ovn-vsctl and linux commands to emulate VM ports.

## Example

`go run cmd/gen/main.go example/ovnnb_db.db example/ovnsb_db.db`

A small section of generated commands is show below, with the complete file shown [here](../../example/ovn_nb_net.sh).  

```bash
# ovn_nb_net.sh
#!/bin/bash

set -e

CIDR_1d646b6e_fcce_4b0a_8895_5e173e6648f9=$(ovn-nbctl create dhcp_options "cidr"="192.168.10.0/24" \
  options='"lease_time"="43200" "mtu"="8942" "router"="192.168.10.1" "server_id"="192.168.10.1" "server_mac"="fa:16:3e:53:b0:2f" "classless_static_route"="{169.254.169.254/32,192.168.10.2, 0.0.0.0/0,192.168.10.1}" "dns_server"="{10.11.5.19, 10.10.160.2, 10.5.30.160}"')
CIDR_3aacef64_db3c_4f79_930d_6282a3e6b95a=$(ovn-nbctl create dhcp_options "cidr"="192.168.33.0/24" \
  options='"server_mac"="fa:16:3e:1f:5d:84" "classless_static_route"="{169.254.169.254/32,192.168.33.100, 0.0.0.0/0,192.168.33.1}" "dns_server"="{10.11.5.19, 10.10.160.2, 10.5.30.160}" "lease_time"="43200" "mtu"="8942" "router"="192.168.33.1" "server_id"="192.168.33.1"')

ovn-nbctl ls-add neutron-d8953248-ba41-4ef4-b7a3-471afed8fd8f
ovn-nbctl lsp-add neutron-d8953248-ba41-4ef4-b7a3-471afed8fd8f 008572d0-b96b-40ef-a3b0-6b20f3650390 
ovn-nbctl lsp-set-addresses 008572d0-b96b-40ef-a3b0-6b20f3650390 "fa:16:3e:27:83:73"
ovn-nbctl lsp-set-enabled 008572d0-b96b-40ef-a3b0-6b20f3650390 enabled
ovn-nbctl lsp-set-type 008572d0-b96b-40ef-a3b0-6b20f3650390 localport
ovn-nbctl lsp-set-options 008572d0-b96b-40ef-a3b0-6b20f3650390 requested-chassis=
ovn-nbctl lsp-add neutron-d8953248-ba41-4ef4-b7a3-471afed8fd8f provnet-27023655-efe2-4757-859b-2e0121b685a3 
ovn-nbctl lsp-set-addresses provnet-27023655-efe2-4757-859b-2e0121b685a3 "unknown"
ovn-nbctl lsp-set-type provnet-27023655-efe2-4757-859b-2e0121b685a3 localnet
ovn-nbctl lsp-set-options provnet-27023655-efe2-4757-859b-2e0121b685a3 mcast_flood=false
ovn-nbctl lsp-set-options provnet-27023655-efe2-4757-859b-2e0121b685a3 mcast_flood_reports=true
ovn-nbctl lsp-set-options provnet-27023655-efe2-4757-859b-2e0121b685a3 network_name=tenant
ovn-nbctl lsp-add neutron-d8953248-ba41-4ef4-b7a3-471afed8fd8f 8b332d14-b00a-4e75-af70-eac6e56afec4 
ovn-nbctl lsp-set-addresses 8b332d14-b00a-4e75-af70-eac6e56afec4 "fa:16:3e:a9:2d:c1 192.0.10.141"
ovn-nbctl lsp-set-port-security 8b332d14-b00a-4e75-af70-eac6e56afec4 "fa:16:3e:a9:2d:c1 192.0.10.141"
ovn-nbctl lsp-set-enabled 8b332d14-b00a-4e75-af70-eac6e56afec4 enabled
ovn-nbctl lsp-set-options 8b332d14-b00a-4e75-af70-eac6e56afec4 mcast_flood_reports=true
ovn-nbctl lsp-set-options 8b332d14-b00a-4e75-af70-eac6e56afec4 requested-chassis=sos-novacompute-0.localdomain
...
```

A partial command sample is shown below, with a full example [here](../../example/sos-novacompute-0.localdomain-ovs).  

```bash
#!/bin/bash

set -e

ip netns add 17df02cc-5577-4f07-a150-49e6499381c8

ovs-vsctl \
      -- add-port br-int 8b332d14-b00a-4 \
      -- set interface 8b332d14-b00a-4 type=internal \
      -- set Interface 8b332d14-b00a-4 external_ids:iface-id=8b332d14-b00a-4e75-af70-eac6e56afec4
ip link set 8b332d14-b00a-4 netns 17df02cc-5577-4f07-a150-49e6499381c8
ip netns exec 17df02cc-5577-4f07-a150-49e6499381c8 ip link set lo up
ip netns exec 17df02cc-5577-4f07-a150-49e6499381c8 ip link set 8b332d14-b00a-4 address fa:16:3e:a9:2d:c1
ip netns exec 17df02cc-5577-4f07-a150-49e6499381c8 ip addr add 192.0.10.141/24 dev 8b332d14-b00a-4
ip netns exec 17df02cc-5577-4f07-a150-49e6499381c8 ip link set 8b332d14-b00a-4 up

ovs-vsctl \
      -- add-port br-int 71d95a8d-9844-4 \
      -- set interface 71d95a8d-9844-4 type=internal \
      -- set Interface 71d95a8d-9844-4 external_ids:iface-id=71d95a8d-9844-409b-b113-d40250684a69
ip link set 71d95a8d-9844-4 netns 17df02cc-5577-4f07-a150-49e6499381c8
ip netns exec 17df02cc-5577-4f07-a150-49e6499381c8 ip link set lo up
ip netns exec 17df02cc-5577-4f07-a150-49e6499381c8 ip link set 71d95a8d-9844-4 address fa:16:3e:d7:da:20
ip netns exec 17df02cc-5577-4f07-a150-49e6499381c8 ip addr add 192.168.10.68/24 dev 71d95a8d-9844-4
ip netns exec 17df02cc-5577-4f07-a150-49e6499381c8 ip link set 71d95a8d-9844-4 up
ip netns exec 17df02cc-5577-4f07-a150-49e6499381c8 ip route add 169.254.169.254/32 via 192.168.10.2 dev 71d95a8d-9844-4
ip netns exec 17df02cc-5577-4f07-a150-49e6499381c8 ip route add 0.0.0.0/0 via 192.168.10.1 dev 71d95a8d-9844-4
...
```

## Usage

After the scripts and inventory have been generated. Copy the files to the server running **ovn-scale-tests**.  An example **podman ps** is shown below on a server running **ovn-scale-tests**.

`[root@rhos-nfv-02 ovn-scale-tests]# podman ps
CONTAINER ID  IMAGE                      COMMAND            CREATED     STATUS         PORTS                               NAMES
5a21c9ef88a8  k8s.gcr.io/pause:3.5                          7 days ago  Up 7 days ago  0.0.0.0:45535-46739->6641-6642/tcp  786a22bc4c60-infra
8b5909d3e299  localhost/ovn-base:latest  tail -f /dev/null  7 days ago  Up 7 days ago  0.0.0.0:45535-46739->6641-6642/tcp  northd
b9d74d81d305  k8s.gcr.io/pause:3.5                          7 days ago  Up 7 days ago                                      6490c98e6bac-infra
d8936f9477a6  localhost/ovn-base:latest  tail -f /dev/null  7 days ago  Up 7 days ago                                      sos-novacompute-0.localdomain-ovs
492acc43960c  localhost/ovn-base:latest  tail -f /dev/null  7 days ago  Up 7 days ago                                      sos-novacompute-0.localdomain-ovn-controller`

To recreate the network...

`podman cp ovn_nb_net.sh northd:/root`

`podman cp <chassis_name>.sh northd:/root`

`podman exec -ti northd /root/ovn_nb_net.sh`

## Limitations

Currently, one will need to run the approprate chassis-specific commands manually.
