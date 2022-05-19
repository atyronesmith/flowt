#!/bin/bash

for uuid in $(ovn-nbctl --bare --column=_uuid list Logical_Router | sed '/^[[:space:]]*$/d'); do ovn-nbctl lr-del "$uuid"; done
for uuid in $(ovn-nbctl --bare --column=_uuid list Logical_Switch | sed '/^[[:space:]]*$/d'); do ovn-nbctl ls-del "$uuid"; done
for name in $(ovn-nbctl --bare --column=name list Port_Group | sed '/^[[:space:]]*$/d'); do ovn-nbctl pg-del "$name"; done
for uuid in $(ovn-nbctl --bare --column=_uuid list Dhcp_Options | sed '/^[[:space:]]*$/d'); do ovn-nbctl dhcp-options-del "$uuid"; done

#TODO HA_CHASSIS