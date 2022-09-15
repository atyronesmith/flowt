#!/bin/bash

set -e

CIDR_1d646b6e_fcce_4b0a_8895_5e173e6648f9=$(ovn-nbctl create dhcp_options "cidr"="192.168.10.0/24" \
  options='"classless_static_route"="{169.254.169.254/32,192.168.10.2, 0.0.0.0/0,192.168.10.1}" "dns_server"="{10.11.5.19, 10.10.160.2, 10.5.30.160}" "lease_time"="43200" "mtu"="8942" "router"="192.168.10.1" "server_id"="192.168.10.1" "server_mac"="fa:16:3e:53:b0:2f"')
CIDR_3aacef64_db3c_4f79_930d_6282a3e6b95a=$(ovn-nbctl create dhcp_options "cidr"="192.168.33.0/24" \
  options='"router"="192.168.33.1" "server_id"="192.168.33.1" "server_mac"="fa:16:3e:1f:5d:84" "classless_static_route"="{169.254.169.254/32,192.168.33.100, 0.0.0.0/0,192.168.33.1}" "dns_server"="{8.8.8.8}" "lease_time"="43200" "mtu"="8942"')

ovn-nbctl ls-add neutron-d8953248-ba41-4ef4-b7a3-471afed8fd8f
LS_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch name=neutron-d8953248-ba41-4ef4-b7a3-471afed8fd8f | awk '{ print $3 }')
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:mtu\"="9000"
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:network_name\"="uplink1"
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:revision_number\"="1"
ovn-nbctl set logical_switch "$LS_UUID" other_config:mcast_flood_unregistered="false"
ovn-nbctl set logical_switch "$LS_UUID" other_config:mcast_snoop="false"
ovn-nbctl set logical_switch "$LS_UUID" other_config:vlan-passthru="false"
ovn-nbctl lsp-add neutron-d8953248-ba41-4ef4-b7a3-471afed8fd8f 008572d0-b96b-40ef-a3b0-6b20f3650390 
ovn-nbctl lsp-set-addresses 008572d0-b96b-40ef-a3b0-6b20f3650390 "fa:16:3e:27:83:73"
ovn-nbctl lsp-set-enabled 008572d0-b96b-40ef-a3b0-6b20f3650390 enabled
ovn-nbctl lsp-set-type 008572d0-b96b-40ef-a3b0-6b20f3650390 localport
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=008572d0-b96b-40ef-a3b0-6b20f3650390 | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="ovnmeta-d8953248-ba41-4ef4-b7a3-471afed8fd8f"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="network:dhcp"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-d8953248-ba41-4ef4-b7a3-471afed8fd8f"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="1"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"=\"\"
ovn-nbctl lsp-set-options 008572d0-b96b-40ef-a3b0-6b20f3650390 requested-chassis=
ovn-nbctl lsp-add neutron-d8953248-ba41-4ef4-b7a3-471afed8fd8f provnet-27023655-efe2-4757-859b-2e0121b685a3 
ovn-nbctl lsp-set-addresses provnet-27023655-efe2-4757-859b-2e0121b685a3 "unknown"
ovn-nbctl lsp-set-type provnet-27023655-efe2-4757-859b-2e0121b685a3 localnet
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=provnet-27023655-efe2-4757-859b-2e0121b685a3 | awk '{ print $3 }')
ovn-nbctl lsp-set-options provnet-27023655-efe2-4757-859b-2e0121b685a3 mcast_flood=false
ovn-nbctl lsp-set-options provnet-27023655-efe2-4757-859b-2e0121b685a3 mcast_flood_reports=true
ovn-nbctl lsp-set-options provnet-27023655-efe2-4757-859b-2e0121b685a3 network_name=tenant
ovn-nbctl lsp-add neutron-d8953248-ba41-4ef4-b7a3-471afed8fd8f 8b332d14-b00a-4e75-af70-eac6e56afec4 
ovn-nbctl lsp-set-addresses 8b332d14-b00a-4e75-af70-eac6e56afec4 "fa:16:3e:a9:2d:c1 192.0.10.141"
ovn-nbctl lsp-set-port-security 8b332d14-b00a-4e75-af70-eac6e56afec4 "fa:16:3e:a9:2d:c1 192.0.10.141"
ovn-nbctl lsp-set-enabled 8b332d14-b00a-4e75-af70-eac6e56afec4 enabled
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=8b332d14-b00a-4e75-af70-eac6e56afec4 | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"="192.0.10.141/24"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="17df02cc-5577-4f07-a150-49e6499381c8"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="compute:nova"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-d8953248-ba41-4ef4-b7a3-471afed8fd8f"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="4"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"="6558094b-5c0d-4258-93ab-6efc734c80bc"
ovn-nbctl lsp-set-options 8b332d14-b00a-4e75-af70-eac6e56afec4 mcast_flood_reports=true
ovn-nbctl lsp-set-options 8b332d14-b00a-4e75-af70-eac6e56afec4 requested-chassis=sos-novacompute-0.localdomain

ovn-nbctl ls-add neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c
LS_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch name=neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c | awk '{ print $3 }')
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:mtu\"="9000"
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:network_name\"="uplink2"
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:revision_number\"="1"
ovn-nbctl set logical_switch "$LS_UUID" other_config:mcast_flood_unregistered="false"
ovn-nbctl set logical_switch "$LS_UUID" other_config:mcast_snoop="false"
ovn-nbctl set logical_switch "$LS_UUID" other_config:vlan-passthru="false"
ovn-nbctl lsp-add neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c 6189098b-df23-4ca1-9062-2b756cad6acc 
ovn-nbctl lsp-set-addresses 6189098b-df23-4ca1-9062-2b756cad6acc "fa:16:3e:d7:13:a4 192.0.11.116"
ovn-nbctl lsp-set-port-security 6189098b-df23-4ca1-9062-2b756cad6acc "fa:16:3e:d7:13:a4 192.0.11.116"
ovn-nbctl lsp-set-enabled 6189098b-df23-4ca1-9062-2b756cad6acc enabled
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=6189098b-df23-4ca1-9062-2b756cad6acc | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"="192.0.11.116/24"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="4082929f-7ab4-4b76-b06f-fd5dea91c609"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="compute:nova"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="4"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"="6558094b-5c0d-4258-93ab-6efc734c80bc"
ovn-nbctl lsp-set-options 6189098b-df23-4ca1-9062-2b756cad6acc mcast_flood_reports=true
ovn-nbctl lsp-set-options 6189098b-df23-4ca1-9062-2b756cad6acc requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-add neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 
ovn-nbctl lsp-set-addresses provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 "unknown"
ovn-nbctl lsp-set-type provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 localnet
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 | awk '{ print $3 }')
ovn-nbctl lsp-set-options provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 mcast_flood=false
ovn-nbctl lsp-set-options provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 mcast_flood_reports=true
ovn-nbctl lsp-set-options provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 network_name=tenant
ovn-nbctl lsp-add neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c 6cbd2589-3c49-4f6a-a139-8cd6fa93522e 
ovn-nbctl lsp-set-addresses 6cbd2589-3c49-4f6a-a139-8cd6fa93522e "fa:16:3e:a9:84:91 192.0.11.219"
ovn-nbctl lsp-set-port-security 6cbd2589-3c49-4f6a-a139-8cd6fa93522e "fa:16:3e:a9:84:91 192.0.11.219"
ovn-nbctl lsp-set-enabled 6cbd2589-3c49-4f6a-a139-8cd6fa93522e enabled
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=6cbd2589-3c49-4f6a-a139-8cd6fa93522e | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"="192.0.11.219/24"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="17df02cc-5577-4f07-a150-49e6499381c8"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="compute:nova"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="4"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"="6558094b-5c0d-4258-93ab-6efc734c80bc"
ovn-nbctl lsp-set-options 6cbd2589-3c49-4f6a-a139-8cd6fa93522e mcast_flood_reports=true
ovn-nbctl lsp-set-options 6cbd2589-3c49-4f6a-a139-8cd6fa93522e requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-add neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c bd253cba-647a-4264-97c2-b16c6196e414 
ovn-nbctl lsp-set-addresses bd253cba-647a-4264-97c2-b16c6196e414 "fa:16:3e:57:37:78"
ovn-nbctl lsp-set-enabled bd253cba-647a-4264-97c2-b16c6196e414 enabled
ovn-nbctl lsp-set-type bd253cba-647a-4264-97c2-b16c6196e414 localport
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=bd253cba-647a-4264-97c2-b16c6196e414 | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="ovnmeta-a6e858b0-c295-41d4-8ff4-858c18695d0c"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="network:dhcp"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="1"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"=\"\"
ovn-nbctl lsp-set-options bd253cba-647a-4264-97c2-b16c6196e414 requested-chassis=

ovn-nbctl ls-add neutron-82e03259-6310-4b9d-9575-7f07d613ce09
LS_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch name=neutron-82e03259-6310-4b9d-9575-7f07d613ce09 | awk '{ print $3 }')
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:mtu\"="8942"
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:network_name\"="test-net"
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:revision_number\"="1"
ovn-nbctl set logical_switch "$LS_UUID" other_config:mcast_flood_unregistered="false"
ovn-nbctl set logical_switch "$LS_UUID" other_config:mcast_snoop="false"
ovn-nbctl set logical_switch "$LS_UUID" other_config:vlan-passthru="false"
ovn-nbctl lsp-add neutron-82e03259-6310-4b9d-9575-7f07d613ce09 a5525be6-6773-4438-8e2a-cd7b41223ff4 
ovn-nbctl lsp-set-addresses a5525be6-6773-4438-8e2a-cd7b41223ff4 "router"
ovn-nbctl lsp-set-enabled a5525be6-6773-4438-8e2a-cd7b41223ff4 enabled
ovn-nbctl lsp-set-type a5525be6-6773-4438-8e2a-cd7b41223ff4 router
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=a5525be6-6773-4438-8e2a-cd7b41223ff4 | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"="192.168.10.1/24"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="72adf022-c39d-4c21-8120-5135f066aa01"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="network:router_interface"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-82e03259-6310-4b9d-9575-7f07d613ce09"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="3"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"=\"\"
ovn-nbctl lsp-set-options a5525be6-6773-4438-8e2a-cd7b41223ff4 router-port=lrp-a5525be6-6773-4438-8e2a-cd7b41223ff4
ovn-nbctl lsp-add neutron-82e03259-6310-4b9d-9575-7f07d613ce09 1d7ec5b6-6c82-49cc-b940-e46c63c59148 
ovn-nbctl lsp-set-addresses 1d7ec5b6-6c82-49cc-b940-e46c63c59148 "fa:16:3e:88:a9:43 192.168.10.203"
ovn-nbctl lsp-set-port-security 1d7ec5b6-6c82-49cc-b940-e46c63c59148 "fa:16:3e:88:a9:43 192.168.10.203"
ovn-nbctl lsp-set-enabled 1d7ec5b6-6c82-49cc-b940-e46c63c59148 enabled
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=1d7ec5b6-6c82-49cc-b940-e46c63c59148 | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"="192.168.10.203/24"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="4082929f-7ab4-4b76-b06f-fd5dea91c609"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="compute:nova"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-82e03259-6310-4b9d-9575-7f07d613ce09"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="4"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"="6558094b-5c0d-4258-93ab-6efc734c80bc"
ovn-nbctl lsp-set-options 1d7ec5b6-6c82-49cc-b940-e46c63c59148 mcast_flood_reports=true
ovn-nbctl lsp-set-options 1d7ec5b6-6c82-49cc-b940-e46c63c59148 requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-set-dhcpv4-options 1d7ec5b6-6c82-49cc-b940-e46c63c59148 "$CIDR_1d646b6e_fcce_4b0a_8895_5e173e6648f9"
ovn-nbctl lsp-add neutron-82e03259-6310-4b9d-9575-7f07d613ce09 71d95a8d-9844-409b-b113-d40250684a69 
ovn-nbctl lsp-set-addresses 71d95a8d-9844-409b-b113-d40250684a69 "fa:16:3e:d7:da:20 192.168.10.68"
ovn-nbctl lsp-set-port-security 71d95a8d-9844-409b-b113-d40250684a69 "fa:16:3e:d7:da:20 192.168.10.68"
ovn-nbctl lsp-set-enabled 71d95a8d-9844-409b-b113-d40250684a69 enabled
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=71d95a8d-9844-409b-b113-d40250684a69 | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"="192.168.10.68/24"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="17df02cc-5577-4f07-a150-49e6499381c8"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="compute:nova"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-82e03259-6310-4b9d-9575-7f07d613ce09"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_fip\"="192.168.122.127"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="5"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"="6558094b-5c0d-4258-93ab-6efc734c80bc f7baaa68-b46e-4977-a361-66b2e21bf35b"
ovn-nbctl lsp-set-options 71d95a8d-9844-409b-b113-d40250684a69 mcast_flood_reports=true
ovn-nbctl lsp-set-options 71d95a8d-9844-409b-b113-d40250684a69 requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-set-dhcpv4-options 71d95a8d-9844-409b-b113-d40250684a69 "$CIDR_1d646b6e_fcce_4b0a_8895_5e173e6648f9"
ovn-nbctl lsp-add neutron-82e03259-6310-4b9d-9575-7f07d613ce09 d0327ab0-796b-4e17-9f4d-9aaa6b42107d 
ovn-nbctl lsp-set-addresses d0327ab0-796b-4e17-9f4d-9aaa6b42107d "fa:16:3e:1c:fa:a5 192.168.10.2"
ovn-nbctl lsp-set-enabled d0327ab0-796b-4e17-9f4d-9aaa6b42107d enabled
ovn-nbctl lsp-set-type d0327ab0-796b-4e17-9f4d-9aaa6b42107d localport
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=d0327ab0-796b-4e17-9f4d-9aaa6b42107d | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"="192.168.10.2/24"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="ovnmeta-82e03259-6310-4b9d-9575-7f07d613ce09"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="network:dhcp"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-82e03259-6310-4b9d-9575-7f07d613ce09"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"=\"\"
ovn-nbctl lsp-set-options d0327ab0-796b-4e17-9f4d-9aaa6b42107d requested-chassis=

ovn-nbctl ls-add neutron-496db99d-97cc-4d52-ab60-2d1386d3626c
LS_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch name=neutron-496db99d-97cc-4d52-ab60-2d1386d3626c | awk '{ print $3 }')
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:mtu\"="8942"
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:network_name\"="test-net2"
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:revision_number\"="1"
ovn-nbctl set logical_switch "$LS_UUID" other_config:mcast_flood_unregistered="false"
ovn-nbctl set logical_switch "$LS_UUID" other_config:mcast_snoop="false"
ovn-nbctl set logical_switch "$LS_UUID" other_config:vlan-passthru="false"
ovn-nbctl lsp-add neutron-496db99d-97cc-4d52-ab60-2d1386d3626c 23d0ed6c-a799-451c-9c69-536ae9079081 
ovn-nbctl lsp-set-addresses 23d0ed6c-a799-451c-9c69-536ae9079081 "fa:16:3e:9f:a8:7e 192.168.33.40"
ovn-nbctl lsp-set-port-security 23d0ed6c-a799-451c-9c69-536ae9079081 "fa:16:3e:9f:a8:7e 192.168.33.40"
ovn-nbctl lsp-set-enabled 23d0ed6c-a799-451c-9c69-536ae9079081 enabled
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=23d0ed6c-a799-451c-9c69-536ae9079081 | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"="192.168.33.40/24"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="1e8f9d6d-eb56-485e-8f84-6bd622a6c5dc"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="compute:nova"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-496db99d-97cc-4d52-ab60-2d1386d3626c"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="4"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"="6558094b-5c0d-4258-93ab-6efc734c80bc"
ovn-nbctl lsp-set-options 23d0ed6c-a799-451c-9c69-536ae9079081 mcast_flood_reports=true
ovn-nbctl lsp-set-options 23d0ed6c-a799-451c-9c69-536ae9079081 requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-set-dhcpv4-options 23d0ed6c-a799-451c-9c69-536ae9079081 "$CIDR_3aacef64_db3c_4f79_930d_6282a3e6b95a"
ovn-nbctl lsp-add neutron-496db99d-97cc-4d52-ab60-2d1386d3626c 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b 
ovn-nbctl lsp-set-addresses 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b "fa:16:3e:94:cf:4a 192.168.33.102"
ovn-nbctl lsp-set-port-security 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b "fa:16:3e:94:cf:4a 192.168.33.102"
ovn-nbctl lsp-set-enabled 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b enabled
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=9fbc5b58-cf4c-4e83-8801-a73b41bdf27b | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"="192.168.33.102/24"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="17df02cc-5577-4f07-a150-49e6499381c8"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="compute:nova"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-496db99d-97cc-4d52-ab60-2d1386d3626c"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="4"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"="6558094b-5c0d-4258-93ab-6efc734c80bc"
ovn-nbctl lsp-set-options 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b mcast_flood_reports=true
ovn-nbctl lsp-set-options 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-set-dhcpv4-options 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b "$CIDR_3aacef64_db3c_4f79_930d_6282a3e6b95a"
ovn-nbctl lsp-add neutron-496db99d-97cc-4d52-ab60-2d1386d3626c d54273a6-ecbd-4ab0-ba13-adbd2a2ae204 
ovn-nbctl lsp-set-addresses d54273a6-ecbd-4ab0-ba13-adbd2a2ae204 "fa:16:3e:14:05:e3 192.168.33.100"
ovn-nbctl lsp-set-enabled d54273a6-ecbd-4ab0-ba13-adbd2a2ae204 enabled
ovn-nbctl lsp-set-type d54273a6-ecbd-4ab0-ba13-adbd2a2ae204 localport
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=d54273a6-ecbd-4ab0-ba13-adbd2a2ae204 | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"="192.168.33.100/24"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="ovnmeta-496db99d-97cc-4d52-ab60-2d1386d3626c"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="network:dhcp"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-496db99d-97cc-4d52-ab60-2d1386d3626c"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"=\"\"
ovn-nbctl lsp-set-options d54273a6-ecbd-4ab0-ba13-adbd2a2ae204 requested-chassis=

ovn-nbctl ls-add neutron-561990d3-f4d5-431d-ae92-a85b83f4f570
LS_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch name=neutron-561990d3-f4d5-431d-ae92-a85b83f4f570 | awk '{ print $3 }')
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:mtu\"="9000"
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:network_name\"="public"
ovn-nbctl set logical_switch "$LS_UUID" external_ids:\"neutron:revision_number\"="1"
ovn-nbctl set logical_switch "$LS_UUID" other_config:mcast_flood_unregistered="false"
ovn-nbctl set logical_switch "$LS_UUID" other_config:mcast_snoop="false"
ovn-nbctl set logical_switch "$LS_UUID" other_config:vlan-passthru="false"
ovn-nbctl lsp-add neutron-561990d3-f4d5-431d-ae92-a85b83f4f570 provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 
ovn-nbctl lsp-set-addresses provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 "unknown"
ovn-nbctl lsp-set-type provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 localnet
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 | awk '{ print $3 }')
ovn-nbctl lsp-set-options provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 mcast_flood=false
ovn-nbctl lsp-set-options provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 mcast_flood_reports=true
ovn-nbctl lsp-set-options provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 network_name=external
ovn-nbctl lsp-add neutron-561990d3-f4d5-431d-ae92-a85b83f4f570 f6e94716-7c85-403e-98ad-f42c19e7f6e5 
ovn-nbctl lsp-set-addresses f6e94716-7c85-403e-98ad-f42c19e7f6e5 "fa:16:3e:e6:32:42"
ovn-nbctl lsp-set-enabled f6e94716-7c85-403e-98ad-f42c19e7f6e5 enabled
ovn-nbctl lsp-set-type f6e94716-7c85-403e-98ad-f42c19e7f6e5 localport
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=f6e94716-7c85-403e-98ad-f42c19e7f6e5 | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="ovnmeta-561990d3-f4d5-431d-ae92-a85b83f4f570"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="network:dhcp"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-561990d3-f4d5-431d-ae92-a85b83f4f570"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"="ff0b6fda266d4d12a0df787aa41f1bb2"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="1"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"=\"\"
ovn-nbctl lsp-set-options f6e94716-7c85-403e-98ad-f42c19e7f6e5 requested-chassis=
ovn-nbctl lsp-add neutron-561990d3-f4d5-431d-ae92-a85b83f4f570 2442a8af-8399-4cfb-86fc-5b379b8cf6b9 
ovn-nbctl lsp-set-addresses 2442a8af-8399-4cfb-86fc-5b379b8cf6b9 "router"
ovn-nbctl lsp-set-enabled 2442a8af-8399-4cfb-86fc-5b379b8cf6b9 enabled
ovn-nbctl lsp-set-type 2442a8af-8399-4cfb-86fc-5b379b8cf6b9 router
LSP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Switch_Port name=2442a8af-8399-4cfb-86fc-5b379b8cf6b9 | awk '{ print $3 }')
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:cidrs\"="192.168.122.141/24"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_id\"="72adf022-c39d-4c21-8120-5135f066aa01"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:device_owner\"="network:router_gateway"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:network_name\"="neutron-561990d3-f4d5-431d-ae92-a85b83f4f570"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:port_name\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:project_id\"=\"\"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:revision_number\"="4"
ovn-nbctl set logical_switch_port "$LSP_UUID" external_ids:\"neutron:security_group_ids\"=\"\"
ovn-nbctl lsp-set-options 2442a8af-8399-4cfb-86fc-5b379b8cf6b9 nat-addresses=router
ovn-nbctl lsp-set-options 2442a8af-8399-4cfb-86fc-5b379b8cf6b9 router-port=lrp-2442a8af-8399-4cfb-86fc-5b379b8cf6b9

ovn-nbctl lr-add neutron-72adf022-c39d-4c21-8120-5135f066aa01
LR_UUID=$(ovn-nbctl --columns=_uuid find Logical_Router name=neutron-72adf022-c39d-4c21-8120-5135f066aa01 | awk '{ print $3 }')
ovn-nbctl set Logical_Router "$LR_UUID" external_ids:\"neutron:availability_zone_hints\"=\"\"
ovn-nbctl set Logical_Router "$LR_UUID" external_ids:\"neutron:gw_port_id\"="2442a8af-8399-4cfb-86fc-5b379b8cf6b9"
ovn-nbctl set Logical_Router "$LR_UUID" external_ids:\"neutron:revision_number\"="5"
ovn-nbctl set Logical_Router "$LR_UUID" external_ids:\"neutron:router_name\"="public"
ovn-nbctl set Logical_Router neutron-72adf022-c39d-4c21-8120-5135f066aa01 options:always_learn_from_arp_request=false
ovn-nbctl set Logical_Router neutron-72adf022-c39d-4c21-8120-5135f066aa01 options:dynamic_neigh_routers=true
ovn-nbctl lrp-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 lrp-2442a8af-8399-4cfb-86fc-5b379b8cf6b9 fa:16:3e:de:b5:e3 192.168.122.141/24 
LRP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Router_Port name=lrp-2442a8af-8399-4cfb-86fc-5b379b8cf6b9 | awk '{ print $3 }')
ovn-nbctl set Logical_Router_Port "$LRP_UUID" external_ids:\"neutron:network_name\"="neutron-561990d3-f4d5-431d-ae92-a85b83f4f570"
ovn-nbctl set Logical_Router_Port "$LRP_UUID" external_ids:\"neutron:revision_number\"="4"
ovn-nbctl set Logical_Router_Port "$LRP_UUID" external_ids:\"neutron:router_name\"="72adf022-c39d-4c21-8120-5135f066aa01"
ovn-nbctl set Logical_Router_Port "$LRP_UUID" external_ids:\"neutron:subnet_ids\"="22bcd50a-d60c-43c3-bc19-1733f52b6a63"
ovn-nbctl lrp-set-gateway-chassis lrp-2442a8af-8399-4cfb-86fc-5b379b8cf6b9 103ca35b-b7ab-4349-a57d-cf9da2ec0e79 1
ovn-nbctl lrp-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 lrp-a5525be6-6773-4438-8e2a-cd7b41223ff4 fa:16:3e:ff:cd:02 192.168.10.1/24 
LRP_UUID=$(ovn-nbctl --columns=_uuid find Logical_Router_Port name=lrp-a5525be6-6773-4438-8e2a-cd7b41223ff4 | awk '{ print $3 }')
ovn-nbctl set Logical_Router_Port "$LRP_UUID" external_ids:\"neutron:network_name\"="neutron-82e03259-6310-4b9d-9575-7f07d613ce09"
ovn-nbctl set Logical_Router_Port "$LRP_UUID" external_ids:\"neutron:revision_number\"="3"
ovn-nbctl set Logical_Router_Port "$LRP_UUID" external_ids:\"neutron:router_name\"="72adf022-c39d-4c21-8120-5135f066aa01"
ovn-nbctl set Logical_Router_Port "$LRP_UUID" external_ids:\"neutron:subnet_ids\"="046cf02c-1271-4a2b-9b44-8e60028b4c37"
ovn-nbctl lr-route-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 10.20.30.0/24 192.168.122.40 
ovn-nbctl lr-route-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 0.0.0.0/0 192.168.122.1 
ovn-nbctl lr-nat-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 snat 192.168.122.141 192.168.10.0/24 
ovn-nbctl lr-nat-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 dnat_and_snat 192.168.122.127 192.168.10.68 

ovn-nbctl pg-add pg_f7baaa68_b46e_4977_a361_66b2e21bf35b
ovn-nbctl pg-set-ports pg_f7baaa68_b46e_4977_a361_66b2e21bf35b 71d95a8d-9844-409b-b113-d40250684a69 
ovn-nbctl --type=port-group acl-add pg_f7baaa68_b46e_4977_a361_66b2e21bf35b from-lport 1002 'inport == @pg_f7baaa68_b46e_4977_a361_66b2e21bf35b && ip4' allow-related
ovn-nbctl --type=port-group acl-add pg_f7baaa68_b46e_4977_a361_66b2e21bf35b from-lport 1002 'inport == @pg_f7baaa68_b46e_4977_a361_66b2e21bf35b && ip6' allow-related
ovn-nbctl --type=port-group acl-add pg_f7baaa68_b46e_4977_a361_66b2e21bf35b to-lport 1002 'outport == @pg_f7baaa68_b46e_4977_a361_66b2e21bf35b && ip4 && ip4.src == 0.0.0.0/0 && tcp && tcp.dst >= 1 && tcp.dst <= 65535' allow-related

ovn-nbctl pg-add pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4
# No ports present
ovn-nbctl --type=port-group acl-add pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 from-lport 1002 'inport == @pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 && ip6' allow-related
ovn-nbctl --type=port-group acl-add pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 from-lport 1002 'inport == @pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 && ip4' allow-related
ovn-nbctl --type=port-group acl-add pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 to-lport 1002 'outport == @pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 && ip6 && ip6.src == $pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4_ip6' allow-related
ovn-nbctl --type=port-group acl-add pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 to-lport 1002 'outport == @pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 && ip4 && ip4.src == $pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4_ip4' allow-related

ovn-nbctl pg-add neutron_pg_drop
ovn-nbctl pg-set-ports neutron_pg_drop 6189098b-df23-4ca1-9062-2b756cad6acc 23d0ed6c-a799-451c-9c69-536ae9079081 1d7ec5b6-6c82-49cc-b940-e46c63c59148 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b 71d95a8d-9844-409b-b113-d40250684a69 6cbd2589-3c49-4f6a-a139-8cd6fa93522e 8b332d14-b00a-4e75-af70-eac6e56afec4 
ovn-nbctl --type=port-group acl-add neutron_pg_drop from-lport 1001 'inport == @neutron_pg_drop && ip' drop
ovn-nbctl --type=port-group acl-add neutron_pg_drop to-lport 1001 'outport == @neutron_pg_drop && ip' drop

ovn-nbctl pg-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc
ovn-nbctl pg-set-ports pg_6558094b_5c0d_4258_93ab_6efc734c80bc 6189098b-df23-4ca1-9062-2b756cad6acc 23d0ed6c-a799-451c-9c69-536ae9079081 1d7ec5b6-6c82-49cc-b940-e46c63c59148 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b 71d95a8d-9844-409b-b113-d40250684a69 6cbd2589-3c49-4f6a-a139-8cd6fa93522e 8b332d14-b00a-4e75-af70-eac6e56afec4 
ovn-nbctl --type=port-group acl-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc from-lport 1002 'inport == @pg_6558094b_5c0d_4258_93ab_6efc734c80bc && ip6' allow-related
ovn-nbctl --type=port-group acl-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc to-lport 1002 'outport == @pg_6558094b_5c0d_4258_93ab_6efc734c80bc && ip4 && ip4.src == 0.0.0.0/0 && tcp && tcp.dst == 22' allow-related
ovn-nbctl --type=port-group acl-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc to-lport 1002 'outport == @pg_6558094b_5c0d_4258_93ab_6efc734c80bc && ip4 && ip4.src == 0.0.0.0/0 && icmp4' allow-related
ovn-nbctl --type=port-group acl-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc from-lport 1002 'inport == @pg_6558094b_5c0d_4258_93ab_6efc734c80bc && ip4' allow-related
ovn-nbctl --type=port-group acl-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc to-lport 1002 'outport == @pg_6558094b_5c0d_4258_93ab_6efc734c80bc && ip6 && ip6.src == $pg_6558094b_5c0d_4258_93ab_6efc734c80bc_ip6' allow-related
ovn-nbctl --type=port-group acl-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc to-lport 1002 'outport == @pg_6558094b_5c0d_4258_93ab_6efc734c80bc && ip4 && ip4.src == $pg_6558094b_5c0d_4258_93ab_6efc734c80bc_ip4' allow-related

ovn-nbctl ha-chassis-group-add default_ha_chassis_group
ovn-nbctl ha-chassis-group-add-chassis default_ha_chassis_group 103ca35b-b7ab-4349-a57d-cf9da2ec0e79 32766
