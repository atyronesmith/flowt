#!/bin/bash

set -e

CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348=$(ovn-nbctl create dhcp_options "cidr"="10.0.0.0/16" \
  options='"dns_server"="{192.168.122.1}" "lease_time"="43200" "mtu"="8942" "router"="10.0.0.1" "server_id"="10.0.0.1" "server_mac"="fa:16:3e:4c:61:9c" "classless_static_route"="{169.254.169.254/32,10.0.0.10, 0.0.0.0/0,10.0.0.1}"')
CIDR_1d646b6e_fcce_4b0a_8895_5e173e6648f9=$(ovn-nbctl create dhcp_options "cidr"="192.168.10.0/24" \
  options='"mtu"="8942" "router"="192.168.10.1" "server_id"="192.168.10.1" "server_mac"="fa:16:3e:53:b0:2f" "classless_static_route"="{169.254.169.254/32,192.168.10.2, 0.0.0.0/0,192.168.10.1}" "dns_server"="{10.11.5.19, 10.10.160.2, 10.5.30.160}" "lease_time"="43200"')
CIDR_3aacef64_db3c_4f79_930d_6282a3e6b95a=$(ovn-nbctl create dhcp_options "cidr"="192.168.33.0/24" \
  options='"mtu"="8942" "router"="192.168.33.1" "server_id"="192.168.33.1" "server_mac"="fa:16:3e:1f:5d:84" "classless_static_route"="{169.254.169.254/32,192.168.33.100, 0.0.0.0/0,192.168.33.1}" "dns_server"="{10.11.5.19, 10.10.160.2, 10.5.30.160}" "lease_time"="43200"')

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

ovn-nbctl ls-add neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c
ovn-nbctl lsp-add neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 
ovn-nbctl lsp-set-addresses provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 "unknown"
ovn-nbctl lsp-set-type provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 localnet
ovn-nbctl lsp-set-options provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 mcast_flood=false
ovn-nbctl lsp-set-options provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 mcast_flood_reports=true
ovn-nbctl lsp-set-options provnet-4dfd6d92-5415-4c83-890a-5aa4f28be252 network_name=tenant
ovn-nbctl lsp-add neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c bd253cba-647a-4264-97c2-b16c6196e414 
ovn-nbctl lsp-set-addresses bd253cba-647a-4264-97c2-b16c6196e414 "fa:16:3e:57:37:78"
ovn-nbctl lsp-set-enabled bd253cba-647a-4264-97c2-b16c6196e414 enabled
ovn-nbctl lsp-set-type bd253cba-647a-4264-97c2-b16c6196e414 localport
ovn-nbctl lsp-set-options bd253cba-647a-4264-97c2-b16c6196e414 requested-chassis=
ovn-nbctl lsp-add neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c 6cbd2589-3c49-4f6a-a139-8cd6fa93522e 
ovn-nbctl lsp-set-addresses 6cbd2589-3c49-4f6a-a139-8cd6fa93522e "fa:16:3e:a9:84:91 192.0.11.219"
ovn-nbctl lsp-set-port-security 6cbd2589-3c49-4f6a-a139-8cd6fa93522e "fa:16:3e:a9:84:91 192.0.11.219"
ovn-nbctl lsp-set-enabled 6cbd2589-3c49-4f6a-a139-8cd6fa93522e enabled
ovn-nbctl lsp-set-options 6cbd2589-3c49-4f6a-a139-8cd6fa93522e mcast_flood_reports=true
ovn-nbctl lsp-set-options 6cbd2589-3c49-4f6a-a139-8cd6fa93522e requested-chassis=sos-novacompute-0.localdomain

ovn-nbctl ls-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad b58a497d-08ea-4dfd-9018-0bcd4a16140f 
ovn-nbctl lsp-set-addresses b58a497d-08ea-4dfd-9018-0bcd4a16140f "fa:16:3e:db:df:0d 10.0.2.87"
ovn-nbctl lsp-set-port-security b58a497d-08ea-4dfd-9018-0bcd4a16140f "fa:16:3e:db:df:0d 10.0.2.87"
ovn-nbctl lsp-set-enabled b58a497d-08ea-4dfd-9018-0bcd4a16140f enabled
ovn-nbctl lsp-set-options b58a497d-08ea-4dfd-9018-0bcd4a16140f mcast_flood_reports=true
ovn-nbctl lsp-set-options b58a497d-08ea-4dfd-9018-0bcd4a16140f requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-set-dhcpv4-options b58a497d-08ea-4dfd-9018-0bcd4a16140f "$CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348"
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad b4e4d22b-7576-49e9-9759-9b7672b21b41 
ovn-nbctl lsp-set-addresses b4e4d22b-7576-49e9-9759-9b7672b21b41 "fa:16:3e:23:47:0b 10.0.0.10" "fa:16:3e:23:47:0b" "fa:16:3e:23:47:0b 10.0.0.10"
ovn-nbctl lsp-set-enabled b4e4d22b-7576-49e9-9759-9b7672b21b41 enabled
ovn-nbctl lsp-set-type b4e4d22b-7576-49e9-9759-9b7672b21b41 localport
ovn-nbctl lsp-set-options b4e4d22b-7576-49e9-9759-9b7672b21b41 requested-chassis=
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad 1b346a67-483e-440b-92cd-d1fc205ac4fd 
ovn-nbctl lsp-set-addresses 1b346a67-483e-440b-92cd-d1fc205ac4fd "fa:16:3e:e9:98:5a 10.0.0.67"
ovn-nbctl lsp-set-port-security 1b346a67-483e-440b-92cd-d1fc205ac4fd "fa:16:3e:e9:98:5a 10.0.0.67"
ovn-nbctl lsp-set-enabled 1b346a67-483e-440b-92cd-d1fc205ac4fd enabled
ovn-nbctl lsp-set-options 1b346a67-483e-440b-92cd-d1fc205ac4fd mcast_flood_reports=true
ovn-nbctl lsp-set-options 1b346a67-483e-440b-92cd-d1fc205ac4fd requested-chassis=
ovn-nbctl lsp-set-dhcpv4-options 1b346a67-483e-440b-92cd-d1fc205ac4fd "$CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348"
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad e70452bc-fef4-4f87-bd66-6c88b1f8ba8a 
ovn-nbctl lsp-set-addresses e70452bc-fef4-4f87-bd66-6c88b1f8ba8a "fa:16:3e:ee:9a:d8 10.0.0.5"
ovn-nbctl lsp-set-port-security e70452bc-fef4-4f87-bd66-6c88b1f8ba8a "fa:16:3e:ee:9a:d8 10.0.0.5"
ovn-nbctl lsp-set-enabled e70452bc-fef4-4f87-bd66-6c88b1f8ba8a enabled
ovn-nbctl lsp-set-type e70452bc-fef4-4f87-bd66-6c88b1f8ba8a virtual
ovn-nbctl lsp-set-options e70452bc-fef4-4f87-bd66-6c88b1f8ba8a mcast_flood_reports=true
ovn-nbctl lsp-set-options e70452bc-fef4-4f87-bd66-6c88b1f8ba8a requested-chassis=
ovn-nbctl lsp-set-options e70452bc-fef4-4f87-bd66-6c88b1f8ba8a virtual-ip=10.0.0.5
ovn-nbctl lsp-set-options e70452bc-fef4-4f87-bd66-6c88b1f8ba8a virtual-parents=03a0961b-7ad6-487a-bb2a-5753e522826c,1b346a67-483e-440b-92cd-d1fc205ac4fd
ovn-nbctl lsp-set-dhcpv4-options e70452bc-fef4-4f87-bd66-6c88b1f8ba8a "$CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348"
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 
ovn-nbctl lsp-set-addresses 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 "fa:16:3e:3b:8d:92 10.0.0.7"
ovn-nbctl lsp-set-port-security 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 "fa:16:3e:3b:8d:92 10.0.0.7"
ovn-nbctl lsp-set-enabled 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 enabled
ovn-nbctl lsp-set-type 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 virtual
ovn-nbctl lsp-set-options 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 mcast_flood_reports=true
ovn-nbctl lsp-set-options 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 requested-chassis=
ovn-nbctl lsp-set-options 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 virtual-ip=10.0.0.7
ovn-nbctl lsp-set-options 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 virtual-parents=b58a497d-08ea-4dfd-9018-0bcd4a16140f,03a0961b-7ad6-487a-bb2a-5753e522826c,1b346a67-483e-440b-92cd-d1fc205ac4fd
ovn-nbctl lsp-set-dhcpv4-options 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 "$CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348"
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad 63085c3a-0a4d-43c4-b527-fa5ff9112aba 
ovn-nbctl lsp-set-addresses 63085c3a-0a4d-43c4-b527-fa5ff9112aba "fa:16:3e:16:ec:e9 10.0.2.249"
ovn-nbctl lsp-set-port-security 63085c3a-0a4d-43c4-b527-fa5ff9112aba "fa:16:3e:16:ec:e9 10.0.2.249"
ovn-nbctl lsp-set-enabled 63085c3a-0a4d-43c4-b527-fa5ff9112aba enabled
ovn-nbctl lsp-set-options 63085c3a-0a4d-43c4-b527-fa5ff9112aba mcast_flood_reports=true
ovn-nbctl lsp-set-options 63085c3a-0a4d-43c4-b527-fa5ff9112aba requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-set-dhcpv4-options 63085c3a-0a4d-43c4-b527-fa5ff9112aba "$CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348"
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad 14dd446e-4e07-43bf-bde0-12655ebb437a 
ovn-nbctl lsp-set-addresses 14dd446e-4e07-43bf-bde0-12655ebb437a "router"
ovn-nbctl lsp-set-enabled 14dd446e-4e07-43bf-bde0-12655ebb437a enabled
ovn-nbctl lsp-set-type 14dd446e-4e07-43bf-bde0-12655ebb437a router
ovn-nbctl lsp-set-options 14dd446e-4e07-43bf-bde0-12655ebb437a router-port=lrp-14dd446e-4e07-43bf-bde0-12655ebb437a
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad 03a0961b-7ad6-487a-bb2a-5753e522826c 
ovn-nbctl lsp-set-addresses 03a0961b-7ad6-487a-bb2a-5753e522826c "fa:16:3e:0b:a2:5d 10.0.3.21"
ovn-nbctl lsp-set-port-security 03a0961b-7ad6-487a-bb2a-5753e522826c "fa:16:3e:0b:a2:5d 10.0.3.21"
ovn-nbctl lsp-set-enabled 03a0961b-7ad6-487a-bb2a-5753e522826c enabled
ovn-nbctl lsp-set-options 03a0961b-7ad6-487a-bb2a-5753e522826c mcast_flood_reports=true
ovn-nbctl lsp-set-options 03a0961b-7ad6-487a-bb2a-5753e522826c requested-chassis=
ovn-nbctl lsp-set-dhcpv4-options 03a0961b-7ad6-487a-bb2a-5753e522826c "$CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348"
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad 63085c3a-0a4d-43c4-b527-fa5ff9112aba 
ovn-nbctl lsp-set-addresses 63085c3a-0a4d-43c4-b527-fa5ff9112aba "fa:16:3e:16:ec:e9 10.0.2.249"
ovn-nbctl lsp-set-port-security 63085c3a-0a4d-43c4-b527-fa5ff9112aba "fa:16:3e:16:ec:e9 10.0.2.249"
ovn-nbctl lsp-set-enabled 63085c3a-0a4d-43c4-b527-fa5ff9112aba enabled
ovn-nbctl lsp-set-options 63085c3a-0a4d-43c4-b527-fa5ff9112aba mcast_flood_reports=true
ovn-nbctl lsp-set-options 63085c3a-0a4d-43c4-b527-fa5ff9112aba requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-set-dhcpv4-options 63085c3a-0a4d-43c4-b527-fa5ff9112aba "$CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348"
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 
ovn-nbctl lsp-set-addresses 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 "fa:16:3e:3b:8d:92 10.0.0.7"
ovn-nbctl lsp-set-port-security 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 "fa:16:3e:3b:8d:92 10.0.0.7"
ovn-nbctl lsp-set-enabled 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 enabled
ovn-nbctl lsp-set-type 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 virtual
ovn-nbctl lsp-set-options 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 mcast_flood_reports=true
ovn-nbctl lsp-set-options 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 requested-chassis=
ovn-nbctl lsp-set-options 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 virtual-ip=10.0.0.7
ovn-nbctl lsp-set-options 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 virtual-parents=b58a497d-08ea-4dfd-9018-0bcd4a16140f,03a0961b-7ad6-487a-bb2a-5753e522826c,1b346a67-483e-440b-92cd-d1fc205ac4fd
ovn-nbctl lsp-set-dhcpv4-options 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4 "$CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348"
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad b58a497d-08ea-4dfd-9018-0bcd4a16140f 
ovn-nbctl lsp-set-addresses b58a497d-08ea-4dfd-9018-0bcd4a16140f "fa:16:3e:db:df:0d 10.0.2.87"
ovn-nbctl lsp-set-port-security b58a497d-08ea-4dfd-9018-0bcd4a16140f "fa:16:3e:db:df:0d 10.0.2.87"
ovn-nbctl lsp-set-enabled b58a497d-08ea-4dfd-9018-0bcd4a16140f enabled
ovn-nbctl lsp-set-options b58a497d-08ea-4dfd-9018-0bcd4a16140f mcast_flood_reports=true
ovn-nbctl lsp-set-options b58a497d-08ea-4dfd-9018-0bcd4a16140f requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-set-dhcpv4-options b58a497d-08ea-4dfd-9018-0bcd4a16140f "$CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348"
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad e70452bc-fef4-4f87-bd66-6c88b1f8ba8a 
ovn-nbctl lsp-set-addresses e70452bc-fef4-4f87-bd66-6c88b1f8ba8a "fa:16:3e:ee:9a:d8 10.0.0.5"
ovn-nbctl lsp-set-port-security e70452bc-fef4-4f87-bd66-6c88b1f8ba8a "fa:16:3e:ee:9a:d8 10.0.0.5"
ovn-nbctl lsp-set-enabled e70452bc-fef4-4f87-bd66-6c88b1f8ba8a enabled
ovn-nbctl lsp-set-type e70452bc-fef4-4f87-bd66-6c88b1f8ba8a virtual
ovn-nbctl lsp-set-options e70452bc-fef4-4f87-bd66-6c88b1f8ba8a mcast_flood_reports=true
ovn-nbctl lsp-set-options e70452bc-fef4-4f87-bd66-6c88b1f8ba8a requested-chassis=
ovn-nbctl lsp-set-options e70452bc-fef4-4f87-bd66-6c88b1f8ba8a virtual-ip=10.0.0.5
ovn-nbctl lsp-set-options e70452bc-fef4-4f87-bd66-6c88b1f8ba8a virtual-parents=03a0961b-7ad6-487a-bb2a-5753e522826c,1b346a67-483e-440b-92cd-d1fc205ac4fd
ovn-nbctl lsp-set-dhcpv4-options e70452bc-fef4-4f87-bd66-6c88b1f8ba8a "$CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348"
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad 14dd446e-4e07-43bf-bde0-12655ebb437a 
ovn-nbctl lsp-set-addresses 14dd446e-4e07-43bf-bde0-12655ebb437a "router"
ovn-nbctl lsp-set-enabled 14dd446e-4e07-43bf-bde0-12655ebb437a enabled
ovn-nbctl lsp-set-type 14dd446e-4e07-43bf-bde0-12655ebb437a router
ovn-nbctl lsp-set-options 14dd446e-4e07-43bf-bde0-12655ebb437a router-port=lrp-14dd446e-4e07-43bf-bde0-12655ebb437a
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad 03a0961b-7ad6-487a-bb2a-5753e522826c 
ovn-nbctl lsp-set-addresses 03a0961b-7ad6-487a-bb2a-5753e522826c "fa:16:3e:0b:a2:5d 10.0.3.21"
ovn-nbctl lsp-set-port-security 03a0961b-7ad6-487a-bb2a-5753e522826c "fa:16:3e:0b:a2:5d 10.0.3.21"
ovn-nbctl lsp-set-enabled 03a0961b-7ad6-487a-bb2a-5753e522826c enabled
ovn-nbctl lsp-set-options 03a0961b-7ad6-487a-bb2a-5753e522826c mcast_flood_reports=true
ovn-nbctl lsp-set-options 03a0961b-7ad6-487a-bb2a-5753e522826c requested-chassis=
ovn-nbctl lsp-set-dhcpv4-options 03a0961b-7ad6-487a-bb2a-5753e522826c "$CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348"
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad 1b346a67-483e-440b-92cd-d1fc205ac4fd 
ovn-nbctl lsp-set-addresses 1b346a67-483e-440b-92cd-d1fc205ac4fd "fa:16:3e:e9:98:5a 10.0.0.67"
ovn-nbctl lsp-set-port-security 1b346a67-483e-440b-92cd-d1fc205ac4fd "fa:16:3e:e9:98:5a 10.0.0.67"
ovn-nbctl lsp-set-enabled 1b346a67-483e-440b-92cd-d1fc205ac4fd enabled
ovn-nbctl lsp-set-options 1b346a67-483e-440b-92cd-d1fc205ac4fd mcast_flood_reports=true
ovn-nbctl lsp-set-options 1b346a67-483e-440b-92cd-d1fc205ac4fd requested-chassis=
ovn-nbctl lsp-set-dhcpv4-options 1b346a67-483e-440b-92cd-d1fc205ac4fd "$CIDR_0dfbe2bf_ebf7_4582_8643_a986315a1348"
ovn-nbctl lsp-add neutron-71f14623-55d2-4ad3-930a-eb41e74499ad b4e4d22b-7576-49e9-9759-9b7672b21b41 
ovn-nbctl lsp-set-addresses b4e4d22b-7576-49e9-9759-9b7672b21b41 "fa:16:3e:23:47:0b 10.0.0.10" "fa:16:3e:23:47:0b" "fa:16:3e:23:47:0b 10.0.0.10"
ovn-nbctl lsp-set-enabled b4e4d22b-7576-49e9-9759-9b7672b21b41 enabled
ovn-nbctl lsp-set-type b4e4d22b-7576-49e9-9759-9b7672b21b41 localport
ovn-nbctl lsp-set-options b4e4d22b-7576-49e9-9759-9b7672b21b41 requested-chassis=

ovn-nbctl ls-add neutron-82e03259-6310-4b9d-9575-7f07d613ce09
ovn-nbctl lsp-add neutron-82e03259-6310-4b9d-9575-7f07d613ce09 a5525be6-6773-4438-8e2a-cd7b41223ff4 
ovn-nbctl lsp-set-addresses a5525be6-6773-4438-8e2a-cd7b41223ff4 "router"
ovn-nbctl lsp-set-enabled a5525be6-6773-4438-8e2a-cd7b41223ff4 enabled
ovn-nbctl lsp-set-type a5525be6-6773-4438-8e2a-cd7b41223ff4 router
ovn-nbctl lsp-set-options a5525be6-6773-4438-8e2a-cd7b41223ff4 router-port=lrp-a5525be6-6773-4438-8e2a-cd7b41223ff4
ovn-nbctl lsp-add neutron-82e03259-6310-4b9d-9575-7f07d613ce09 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb 
ovn-nbctl lsp-set-addresses 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb "fa:16:3e:90:63:65 192.168.10.244"
ovn-nbctl lsp-set-port-security 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb "fa:16:3e:90:63:65 192.168.10.244"
ovn-nbctl lsp-set-enabled 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb enabled
ovn-nbctl lsp-set-options 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb mcast_flood_reports=true
ovn-nbctl lsp-set-options 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-set-dhcpv4-options 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb "$CIDR_1d646b6e_fcce_4b0a_8895_5e173e6648f9"
ovn-nbctl lsp-add neutron-82e03259-6310-4b9d-9575-7f07d613ce09 d0327ab0-796b-4e17-9f4d-9aaa6b42107d 
ovn-nbctl lsp-set-addresses d0327ab0-796b-4e17-9f4d-9aaa6b42107d "fa:16:3e:1c:fa:a5 192.168.10.2"
ovn-nbctl lsp-set-enabled d0327ab0-796b-4e17-9f4d-9aaa6b42107d enabled
ovn-nbctl lsp-set-type d0327ab0-796b-4e17-9f4d-9aaa6b42107d localport
ovn-nbctl lsp-set-options d0327ab0-796b-4e17-9f4d-9aaa6b42107d requested-chassis=
ovn-nbctl lsp-add neutron-82e03259-6310-4b9d-9575-7f07d613ce09 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb 
ovn-nbctl lsp-set-addresses 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb "fa:16:3e:90:63:65 192.168.10.244"
ovn-nbctl lsp-set-port-security 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb "fa:16:3e:90:63:65 192.168.10.244"
ovn-nbctl lsp-set-enabled 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb enabled
ovn-nbctl lsp-set-options 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb mcast_flood_reports=true
ovn-nbctl lsp-set-options 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-set-dhcpv4-options 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb "$CIDR_1d646b6e_fcce_4b0a_8895_5e173e6648f9"
ovn-nbctl lsp-add neutron-82e03259-6310-4b9d-9575-7f07d613ce09 71d95a8d-9844-409b-b113-d40250684a69 
ovn-nbctl lsp-set-addresses 71d95a8d-9844-409b-b113-d40250684a69 "fa:16:3e:d7:da:20 192.168.10.68"
ovn-nbctl lsp-set-port-security 71d95a8d-9844-409b-b113-d40250684a69 "fa:16:3e:d7:da:20 192.168.10.68"
ovn-nbctl lsp-set-enabled 71d95a8d-9844-409b-b113-d40250684a69 enabled
ovn-nbctl lsp-set-options 71d95a8d-9844-409b-b113-d40250684a69 mcast_flood_reports=true
ovn-nbctl lsp-set-options 71d95a8d-9844-409b-b113-d40250684a69 requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-set-dhcpv4-options 71d95a8d-9844-409b-b113-d40250684a69 "$CIDR_1d646b6e_fcce_4b0a_8895_5e173e6648f9"

ovn-nbctl ls-add neutron-496db99d-97cc-4d52-ab60-2d1386d3626c
ovn-nbctl lsp-add neutron-496db99d-97cc-4d52-ab60-2d1386d3626c d54273a6-ecbd-4ab0-ba13-adbd2a2ae204 
ovn-nbctl lsp-set-addresses d54273a6-ecbd-4ab0-ba13-adbd2a2ae204 "fa:16:3e:14:05:e3" "fa:16:3e:14:05:e3" "fa:16:3e:14:05:e3 192.168.33.100"
ovn-nbctl lsp-set-enabled d54273a6-ecbd-4ab0-ba13-adbd2a2ae204 enabled
ovn-nbctl lsp-set-type d54273a6-ecbd-4ab0-ba13-adbd2a2ae204 localport
ovn-nbctl lsp-set-options d54273a6-ecbd-4ab0-ba13-adbd2a2ae204 requested-chassis=
ovn-nbctl lsp-add neutron-496db99d-97cc-4d52-ab60-2d1386d3626c 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b 
ovn-nbctl lsp-set-addresses 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b "fa:16:3e:94:cf:4a 192.168.33.102"
ovn-nbctl lsp-set-port-security 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b "fa:16:3e:94:cf:4a 192.168.33.102"
ovn-nbctl lsp-set-enabled 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b enabled
ovn-nbctl lsp-set-options 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b mcast_flood_reports=true
ovn-nbctl lsp-set-options 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b requested-chassis=sos-novacompute-0.localdomain
ovn-nbctl lsp-set-dhcpv4-options 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b "$CIDR_3aacef64_db3c_4f79_930d_6282a3e6b95a"

ovn-nbctl ls-add neutron-561990d3-f4d5-431d-ae92-a85b83f4f570
ovn-nbctl lsp-add neutron-561990d3-f4d5-431d-ae92-a85b83f4f570 provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 
ovn-nbctl lsp-set-addresses provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 "unknown"
ovn-nbctl lsp-set-type provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 localnet
ovn-nbctl lsp-set-options provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 mcast_flood=false
ovn-nbctl lsp-set-options provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 mcast_flood_reports=true
ovn-nbctl lsp-set-options provnet-58b5c9c3-7645-4eb9-9c2f-ae5e21cf66f0 network_name=external
ovn-nbctl lsp-add neutron-561990d3-f4d5-431d-ae92-a85b83f4f570 f6e94716-7c85-403e-98ad-f42c19e7f6e5 
ovn-nbctl lsp-set-addresses f6e94716-7c85-403e-98ad-f42c19e7f6e5 "fa:16:3e:e6:32:42"
ovn-nbctl lsp-set-enabled f6e94716-7c85-403e-98ad-f42c19e7f6e5 enabled
ovn-nbctl lsp-set-type f6e94716-7c85-403e-98ad-f42c19e7f6e5 localport
ovn-nbctl lsp-set-options f6e94716-7c85-403e-98ad-f42c19e7f6e5 requested-chassis=
ovn-nbctl lsp-add neutron-561990d3-f4d5-431d-ae92-a85b83f4f570 ab5dc20a-fb92-4f4f-97ef-04282384b598 
ovn-nbctl lsp-set-addresses ab5dc20a-fb92-4f4f-97ef-04282384b598 "router"
ovn-nbctl lsp-set-enabled ab5dc20a-fb92-4f4f-97ef-04282384b598 enabled
ovn-nbctl lsp-set-type ab5dc20a-fb92-4f4f-97ef-04282384b598 router
ovn-nbctl lsp-set-options ab5dc20a-fb92-4f4f-97ef-04282384b598 nat-addresses=router
ovn-nbctl lsp-set-options ab5dc20a-fb92-4f4f-97ef-04282384b598 router-port=lrp-ab5dc20a-fb92-4f4f-97ef-04282384b598
ovn-nbctl lsp-add neutron-561990d3-f4d5-431d-ae92-a85b83f4f570 2442a8af-8399-4cfb-86fc-5b379b8cf6b9 
ovn-nbctl lsp-set-addresses 2442a8af-8399-4cfb-86fc-5b379b8cf6b9 "router"
ovn-nbctl lsp-set-enabled 2442a8af-8399-4cfb-86fc-5b379b8cf6b9 enabled
ovn-nbctl lsp-set-type 2442a8af-8399-4cfb-86fc-5b379b8cf6b9 router
ovn-nbctl lsp-set-options 2442a8af-8399-4cfb-86fc-5b379b8cf6b9 nat-addresses=router
ovn-nbctl lsp-set-options 2442a8af-8399-4cfb-86fc-5b379b8cf6b9 router-port=lrp-2442a8af-8399-4cfb-86fc-5b379b8cf6b9
ovn-nbctl lsp-add neutron-561990d3-f4d5-431d-ae92-a85b83f4f570 ab5dc20a-fb92-4f4f-97ef-04282384b598 
ovn-nbctl lsp-set-addresses ab5dc20a-fb92-4f4f-97ef-04282384b598 "router"
ovn-nbctl lsp-set-enabled ab5dc20a-fb92-4f4f-97ef-04282384b598 enabled
ovn-nbctl lsp-set-type ab5dc20a-fb92-4f4f-97ef-04282384b598 router
ovn-nbctl lsp-set-options ab5dc20a-fb92-4f4f-97ef-04282384b598 nat-addresses=router
ovn-nbctl lsp-set-options ab5dc20a-fb92-4f4f-97ef-04282384b598 router-port=lrp-ab5dc20a-fb92-4f4f-97ef-04282384b598

ovn-nbctl lr-add neutron-72adf022-c39d-4c21-8120-5135f066aa01
ovn-nbctl set logical_router neutron-72adf022-c39d-4c21-8120-5135f066aa01 options:always_learn_from_arp_request=false
ovn-nbctl set logical_router neutron-72adf022-c39d-4c21-8120-5135f066aa01 options:dynamic_neigh_routers=true
ovn-nbctl lrp-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 lrp-2442a8af-8399-4cfb-86fc-5b379b8cf6b9 fa:16:3e:de:b5:e3 192.168.122.141/24 
ovn-nbctl lrp-set-gateway-chassis lrp-2442a8af-8399-4cfb-86fc-5b379b8cf6b9 7417d7c9-d686-49fe-8557-ab4afc38ae40
ovn-nbctl lrp-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 lrp-a5525be6-6773-4438-8e2a-cd7b41223ff4 fa:16:3e:ff:cd:02 192.168.10.1/24 
ovn-nbctl lr-route-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 0.0.0.0/0 192.168.122.1 
ovn-nbctl lr-route-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 10.20.30.0/24 192.168.122.40 
ovn-nbctl  lr-nat-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 snat 192.168.122.141 192.168.10.0/24 
ovn-nbctl  lr-nat-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 dnat_and_snat 192.168.122.138 192.168.10.244 
ovn-nbctl  lr-nat-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 dnat_and_snat 192.168.122.138 192.168.10.244 
ovn-nbctl  lr-nat-add neutron-72adf022-c39d-4c21-8120-5135f066aa01 dnat_and_snat 192.168.122.127 192.168.10.68 

ovn-nbctl lr-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2
ovn-nbctl set logical_router neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 options:always_learn_from_arp_request=false
ovn-nbctl set logical_router neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 options:dynamic_neigh_routers=true
ovn-nbctl lrp-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 lrp-ab5dc20a-fb92-4f4f-97ef-04282384b598 fa:16:3e:90:ad:39 192.168.122.135/24 
ovn-nbctl lrp-set-gateway-chassis lrp-ab5dc20a-fb92-4f4f-97ef-04282384b598 f2e3f73d-ac37-4dcc-a194-2e8a16178c08
ovn-nbctl lrp-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 lrp-14dd446e-4e07-43bf-bde0-12655ebb437a fa:16:3e:96:51:d2 10.0.0.1/16 
ovn-nbctl lrp-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 lrp-14dd446e-4e07-43bf-bde0-12655ebb437a fa:16:3e:96:51:d2 10.0.0.1/16 
ovn-nbctl lrp-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 lrp-ab5dc20a-fb92-4f4f-97ef-04282384b598 fa:16:3e:90:ad:39 192.168.122.135/24 
ovn-nbctl lrp-set-gateway-chassis lrp-ab5dc20a-fb92-4f4f-97ef-04282384b598 f2e3f73d-ac37-4dcc-a194-2e8a16178c08
ovn-nbctl lr-route-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 0.0.0.0/0 192.168.122.1 
ovn-nbctl lr-route-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 0.0.0.0/0 192.168.122.1 
ovn-nbctl  lr-nat-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 dnat_and_snat 192.168.122.147 10.0.2.249 
ovn-nbctl  lr-nat-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 dnat_and_snat 192.168.122.151 10.0.0.7 
ovn-nbctl  lr-nat-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 dnat_and_snat 192.168.122.150 10.0.0.5 
ovn-nbctl  lr-nat-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 snat 192.168.122.135 10.0.0.0/16 
ovn-nbctl  lr-nat-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 dnat_and_snat 192.168.122.147 10.0.2.249 
ovn-nbctl  lr-nat-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 dnat_and_snat 192.168.122.151 10.0.0.7 
ovn-nbctl  lr-nat-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 dnat_and_snat 192.168.122.150 10.0.0.5 
ovn-nbctl  lr-nat-add neutron-73b70493-8cd4-48f8-8389-a40b85b67de2 snat 192.168.122.135 10.0.0.0/16 

ovn-nbctl pg-add pg_f7baaa68_b46e_4977_a361_66b2e21bf35b
ovn-nbctl --type=port-group acl-add pg_f7baaa68_b46e_4977_a361_66b2e21bf35b from-lport 1002 'inport == @pg_f7baaa68_b46e_4977_a361_66b2e21bf35b && ip4' allow-related
ovn-nbctl --type=port-group acl-add pg_f7baaa68_b46e_4977_a361_66b2e21bf35b from-lport 1002 'inport == @pg_f7baaa68_b46e_4977_a361_66b2e21bf35b && ip6' allow-related
ovn-nbctl --type=port-group acl-add pg_f7baaa68_b46e_4977_a361_66b2e21bf35b to-lport 1002 'outport == @pg_f7baaa68_b46e_4977_a361_66b2e21bf35b && ip4 && ip4.src == 0.0.0.0/0 && tcp && tcp.dst == 3306' allow-related
ovn-nbctl --type=port-group acl-add pg_f7baaa68_b46e_4977_a361_66b2e21bf35b to-lport 1002 'outport == @pg_f7baaa68_b46e_4977_a361_66b2e21bf35b && ip4 && ip4.src == 192.168.3.5/32 && icmp4' allow-related

ovn-nbctl pg-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d
ovn-nbctl pg-set-ports pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4
ovn-nbctl pg-set-ports pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 10.0.0.0/16 && udp && udp.dst == 6081' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 10.0.0.0/16 && udp && udp.dst == 4500' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 10.0.0.0/16 && udp && udp.dst == 4789' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst == 10250' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 10.0.0.0/16 && ip.proto == 112' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 0.0.0.0/0 && icmp4' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst == 22' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 0.0.0.0/0 && udp && udp.dst >= 30000 && udp.dst <= 32767' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst == 1936' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 10.0.0.0/16 && udp && udp.dst >= 9000 && udp.dst <= 9999' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 10.0.0.0/16 && ip.proto == 50' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 0.0.0.0/0 && tcp && tcp.dst == 443' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 0.0.0.0/0 && tcp && tcp.dst == 80' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst >= 9000 && tcp.dst <= 9999' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 10.0.0.0/16 && udp && udp.dst == 500' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d from-lport 1002 'inport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip6' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d from-lport 1002 'inport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4' allow-related
ovn-nbctl --type=port-group acl-add pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d to-lport 1002 'outport == @pg_b002956d_bba1_4f5e_bdb3_907aeef62e2d && ip4 && ip4.src == 0.0.0.0/0 && tcp && tcp.dst >= 30000 && tcp.dst <= 32767' allow-related

ovn-nbctl pg-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d
ovn-nbctl pg-set-ports pg_d685ccbb_4b47_44bf_9c8f_8914999b851d b58a497d-08ea-4dfd-9018-0bcd4a16140f
ovn-nbctl pg-set-ports pg_d685ccbb_4b47_44bf_9c8f_8914999b851d 1b346a67-483e-440b-92cd-d1fc205ac4fd
ovn-nbctl pg-set-ports pg_d685ccbb_4b47_44bf_9c8f_8914999b851d e70452bc-fef4-4f87-bd66-6c88b1f8ba8a
ovn-nbctl pg-set-ports pg_d685ccbb_4b47_44bf_9c8f_8914999b851d 63085c3a-0a4d-43c4-b527-fa5ff9112aba
ovn-nbctl pg-set-ports pg_d685ccbb_4b47_44bf_9c8f_8914999b851d 03a0961b-7ad6-487a-bb2a-5753e522826c
ovn-nbctl pg-set-ports pg_d685ccbb_4b47_44bf_9c8f_8914999b851d 63085c3a-0a4d-43c4-b527-fa5ff9112aba
ovn-nbctl pg-set-ports pg_d685ccbb_4b47_44bf_9c8f_8914999b851d b58a497d-08ea-4dfd-9018-0bcd4a16140f
ovn-nbctl pg-set-ports pg_d685ccbb_4b47_44bf_9c8f_8914999b851d e70452bc-fef4-4f87-bd66-6c88b1f8ba8a
ovn-nbctl pg-set-ports pg_d685ccbb_4b47_44bf_9c8f_8914999b851d 03a0961b-7ad6-487a-bb2a-5753e522826c
ovn-nbctl pg-set-ports pg_d685ccbb_4b47_44bf_9c8f_8914999b851d 1b346a67-483e-440b-92cd-d1fc205ac4fd
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && udp && udp.dst == 53' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && ip.proto == 112' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 0.0.0.0/0 && udp && udp.dst >= 30000 && udp.dst <= 32767' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 0.0.0.0/0 && tcp && tcp.dst == 80' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst >= 2379 && tcp.dst <= 2380' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 0.0.0.0/0 && tcp && tcp.dst == 6443' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst >= 6641 && tcp.dst <= 6642' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst == 10257' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && udp && udp.dst >= 9000 && udp.dst <= 9999' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst == 22' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && udp && udp.dst == 4789' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst == 1936' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst == 53' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 0.0.0.0/0 && icmp4' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && udp && udp.dst == 4500' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d from-lport 1002 'inport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 0.0.0.0/0 && tcp && tcp.dst >= 30000 && tcp.dst <= 32767' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst == 22623' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && udp && udp.dst == 500' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 0.0.0.0/0 && tcp && tcp.dst == 443' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && udp && udp.dst == 6081' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst == 10250' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d from-lport 1002 'inport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip6' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst == 10259' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && ip.proto == 50' allow-related
ovn-nbctl --type=port-group acl-add pg_d685ccbb_4b47_44bf_9c8f_8914999b851d to-lport 1002 'outport == @pg_d685ccbb_4b47_44bf_9c8f_8914999b851d && ip4 && ip4.src == 10.0.0.0/16 && tcp && tcp.dst >= 9000 && tcp.dst <= 9999' allow-related

ovn-nbctl pg-add pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4
ovn-nbctl --type=port-group acl-add pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 from-lport 1002 'inport == @pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 && ip6' allow-related
ovn-nbctl --type=port-group acl-add pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 from-lport 1002 'inport == @pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 && ip4' allow-related
ovn-nbctl --type=port-group acl-add pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 to-lport 1002 'outport == @pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 && ip6 && ip6.src == $pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4_ip6' allow-related
ovn-nbctl --type=port-group acl-add pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 to-lport 1002 'outport == @pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4 && ip4 && ip4.src == $pg_706cec15_e52f_4a0a_8139_2ff14b5fd1a4_ip4' allow-related

ovn-nbctl pg-add neutron_pg_drop
ovn-nbctl pg-set-ports neutron_pg_drop b58a497d-08ea-4dfd-9018-0bcd4a16140f
ovn-nbctl pg-set-ports neutron_pg_drop 1b346a67-483e-440b-92cd-d1fc205ac4fd
ovn-nbctl pg-set-ports neutron_pg_drop e70452bc-fef4-4f87-bd66-6c88b1f8ba8a
ovn-nbctl pg-set-ports neutron_pg_drop 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4
ovn-nbctl pg-set-ports neutron_pg_drop 63085c3a-0a4d-43c4-b527-fa5ff9112aba
ovn-nbctl pg-set-ports neutron_pg_drop 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb
ovn-nbctl pg-set-ports neutron_pg_drop 03a0961b-7ad6-487a-bb2a-5753e522826c
ovn-nbctl pg-set-ports neutron_pg_drop 63085c3a-0a4d-43c4-b527-fa5ff9112aba
ovn-nbctl pg-set-ports neutron_pg_drop 9bb9cbcc-a1ea-496b-a0ea-3396864b1ac4
ovn-nbctl pg-set-ports neutron_pg_drop b58a497d-08ea-4dfd-9018-0bcd4a16140f
ovn-nbctl pg-set-ports neutron_pg_drop e70452bc-fef4-4f87-bd66-6c88b1f8ba8a
ovn-nbctl pg-set-ports neutron_pg_drop 03a0961b-7ad6-487a-bb2a-5753e522826c
ovn-nbctl pg-set-ports neutron_pg_drop 1b346a67-483e-440b-92cd-d1fc205ac4fd
ovn-nbctl pg-set-ports neutron_pg_drop 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb
ovn-nbctl pg-set-ports neutron_pg_drop 71d95a8d-9844-409b-b113-d40250684a69
ovn-nbctl pg-set-ports neutron_pg_drop 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b
ovn-nbctl pg-set-ports neutron_pg_drop 8b332d14-b00a-4e75-af70-eac6e56afec4
ovn-nbctl pg-set-ports neutron_pg_drop 6cbd2589-3c49-4f6a-a139-8cd6fa93522e
ovn-nbctl --type=port-group acl-add neutron_pg_drop from-lport 1001 'inport == @neutron_pg_drop && ip' drop
ovn-nbctl --type=port-group acl-add neutron_pg_drop to-lport 1001 'outport == @neutron_pg_drop && ip' drop

ovn-nbctl pg-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc
ovn-nbctl pg-set-ports pg_6558094b_5c0d_4258_93ab_6efc734c80bc 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb
ovn-nbctl pg-set-ports pg_6558094b_5c0d_4258_93ab_6efc734c80bc 748a84ff-60ff-47e9-b2cd-dcc91cc31cdb
ovn-nbctl pg-set-ports pg_6558094b_5c0d_4258_93ab_6efc734c80bc 71d95a8d-9844-409b-b113-d40250684a69
ovn-nbctl pg-set-ports pg_6558094b_5c0d_4258_93ab_6efc734c80bc 9fbc5b58-cf4c-4e83-8801-a73b41bdf27b
ovn-nbctl pg-set-ports pg_6558094b_5c0d_4258_93ab_6efc734c80bc 8b332d14-b00a-4e75-af70-eac6e56afec4
ovn-nbctl pg-set-ports pg_6558094b_5c0d_4258_93ab_6efc734c80bc 6cbd2589-3c49-4f6a-a139-8cd6fa93522e
ovn-nbctl --type=port-group acl-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc from-lport 1002 'inport == @pg_6558094b_5c0d_4258_93ab_6efc734c80bc && ip6' allow-related
ovn-nbctl --type=port-group acl-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc to-lport 1002 'outport == @pg_6558094b_5c0d_4258_93ab_6efc734c80bc && ip4 && ip4.src == 0.0.0.0/0 && icmp4' allow-related
ovn-nbctl --type=port-group acl-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc from-lport 1002 'inport == @pg_6558094b_5c0d_4258_93ab_6efc734c80bc && ip4' allow-related
ovn-nbctl --type=port-group acl-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc to-lport 1002 'outport == @pg_6558094b_5c0d_4258_93ab_6efc734c80bc && ip6 && ip6.src == $pg_6558094b_5c0d_4258_93ab_6efc734c80bc_ip6' allow-related
ovn-nbctl --type=port-group acl-add pg_6558094b_5c0d_4258_93ab_6efc734c80bc to-lport 1002 'outport == @pg_6558094b_5c0d_4258_93ab_6efc734c80bc && ip4 && ip4.src == $pg_6558094b_5c0d_4258_93ab_6efc734c80bc_ip4' allow-related
