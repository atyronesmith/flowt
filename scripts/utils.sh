#!/bin/bash

MAX_LEN=50

OUTPUT_DIR="data"

declare -A flows_keywords=(
    [cookie]="true"
    [duration]="true"
    [n_packets]="true"
    [n_bytes]="true"
    [idle_age]="true"
    [priority]="true"
    [action]="true"
)

declare -A EDGES

export flows_keywords

function gen_gather_filename() {
    host="$1"
    bridge="$2"

    if [[ "$host" =~ ^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$ ]]; then
        fname=${host//./_}
    else
        fname=$host
    fi

    mkdir -p "$OUTPUT_DIR"

    fname+="_$bridge.txt"

    echo "$OUTPUT_DIR/$fname"
}

function html_encode() {
    local rv="$1"
    local max_len="$2"

    rv=${rv:0:max_len}
    rv=${rv//&/&amp;}
    rv=${rv//>/&gt;}
    rv=${rv//</&lt;}

    echo "$rv"
}

function gen_html_flowtable_hdr() {
    table_num="$1"

    cat <<EOF
      <TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0" CELLPADDING="1">
      <TR>
        <TD PORT="input">Priority</TD>
        <TD>Match</TD>
        <TD ALIGN="center">Action (Table ${table_num})</TD>
      </TR>
EOF
}

function gen_html_action_td() {
    local action="$1"

    if [[ $action =~ resubmit\([^,]*,([^\)]+)\) ]]; then
        # We can have multiple resubmit(,X) in a list of actions
        # Need to search for resubmit(,X) and if it is present capture X
        # add an output port
        EDGES[$edge_name]="table${BASH_REMATCH[1]}:input"
        if [ "$rule_count" -gt 1 ]; then
            gen_html_tr "${port_name}" "${act}" ""
        fi
    fi

}

function gen_html_flowtable_row() {
    local content="$1"

    if [[ $content =~ resubmit() ]]; then
        cat <<EOF
      <TR>
        <TD ALIGN="left" PORT=>$content</TD>
      </TR>
EOF
    else
        cat <<EOF
      <TR>
        <TD ALIGN="left">$content</TD>
      </TR>
EOF
    fi
}

function gen_dot_flowtable_start() {
    local num="$1"

    cat <<EOF
    "table$num" [
    label=<
EOF
    gen_html_flowtable_hdr "$num"
}

function gen_html_flowtable_first_row() {
    local priority="$1"
    local inport="$2"
    local rowspan="$3"
    local action="$4"

    action=$(html_encode "$action" "$MAX_LEN")

    cat <<EOF
      <TR>
        <TD rowspan="$rowspan">$priority</TD>
        <TD rowspan="$rowspan">$inport</TD>
        <TD ALIGN="left">$action</TD>
      </TR>
EOF
}

function gen_html_tr() {
    local port_name="$1"
    local content="$2"
    local bgcolor="$3"

    content=$(html_encode "$content" "$MAX_LEN")

    attribs="ALIGN=\"left\""

    [ -n "$port_name" ] && attribs+=" PORT=\"${port_name}\""
    [ -n "$bgcolor" ] && attribs+=" BGCOLOR=\"${bgcolor}\""

    cat <<EOF
      <TR>
        <TD $attribs>${content}</TD>
      </TR>
EOF
}

function gen_dot_flowtable_rule() {
    local priority_str="$1"
    local actions="$2"  # Commas separate string of actions
    local rule_num="$3" # The rule number in the current flow table
    local table_num="$4"

    # A SUBMIT(,N) action contains a comma.  There a more complicated
    # regex for splitting the action string is needed
    scrubed=$(sed -E 's/([^,\(\)]+(\(.*?\))*)+/\1 /gm;t;d' <<<"$actions")
    scrubed=${scrubed// ,/ }
    IFS=' ' read -r -a action_list <<<"$scrubed"
    # At this point the action_list array contains individual actions

    # The Priority and Match cells span the action list in the HTML table
    # The first action in action_list[@] needs to be part of the first
    # row definition
    rowspan=$((${#action_list[@]}))

    # Rule priority is the first item in the Priority string
    # Match to find the priority value and the rest of the string
    [[ $priority_str =~ ([^,]+),(.*) ]] || printf "Invalid priority format: %s" "$priority_str"
    priority="${BASH_REMATCH[1]}"
    # The match criteria follow the priority
    match=$(html_encode "${BASH_REMATCH[2]}" "100")
    match=${match//,/<BR/>}

    # The first row must container the first action so that rowspan works 
    # The first action could be a Submit(,N)
    gen_html_flowtable_first_row "$priority" "$match" "$rowspan" "${action_list[0]}"

    count=1
    rule_count=1
    for act in "${action_list[@]}"; do
        port_name="o${rule_num}_${count}"
        edge_name="table${table_num}:${port_name}"
        if [[ $act =~ resubmit\([^,]*,([^\)]+)\) ]]; then
            # We can have multiple resubmit(,X) in a list of actions
            # Need to search for resubmit(,X) and if it is present capture X
            # add an output port
            EDGES[$edge_name]="table${BASH_REMATCH[1]}:input"
            if [ $rule_count -gt 1 ]; then
                gen_html_tr "${port_name}" "${act}" ""
            fi
            count=$((count + 1))
        elif [[ $act =~ drop ]]; then
            #            EDGES[$edge_name]="drop"
            if [ $count -gt 1 ]; then
                gen_html_tr "${port_name}" "${act}" "red"
            fi
        else
            gen_html_tr "" "${act}" ""
        fi
        rule_count=$((rule_count + 1))
    done
}

function gen_dot_flowtable_end() {
    cat <<EOF
      </TABLE>
      >
    ];
EOF
}

function gen_dot_graph_start() {
    cat <<EOF
digraph {
    concentrate=True;
    rankdir=LR;
    node [shape="none" fontsize="6"]
    edge [fontname="Helvetica,Arial,sans-serif"]
    graph [
        rankdir = "LR"
    ];
EOF
}

function gen_dot_graph_end() {
    cat <<EOF
  "drop" [
    shape="invtriangle"
  ]
EOF
    for edge in "${!EDGES[@]}"; do
        cat <<EOF
    $edge -> ${EDGES[$edge]}
EOF
    done

    cat <<EOF
}
EOF
}
