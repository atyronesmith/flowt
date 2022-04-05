#!/bin/bash

usage() {
    prog=$(basename "$0")
    cat <<-EOM
    Parse flows from a node/bridge
    Usage:
        $prog [-h] [-d] [-v] [-b bridge_name] host
            ip_address           -- hostname or IP Address of the node to query
            
    Options
            -b  -- The name of the bridge to query. Defaults to br-int.
            -t  -- Number of tables to process
            -h  -- Print this usage and exit.
            -d  -- set -x
            -v  -- Provide more info
    ENVIRONEMENT VARIABLES
            None

EOM
    exit 0
}

VERBOSE="false"
export VERBOSE

bridge="br-int"
max_tables=100

while getopts ":hvdb:t:" opt; do
    case ${opt} in
    t)
        max_tables=$OPTARG
        ;;
    b)
        bridge=$OPTARG
        ;;
    v)
        VERBOSE="true"
        ;;
    d)
        set -x
        ;;
    h)
        usage
        exit 0
        ;;
    \?)
        echo "Invalid Option: -$OPTARG" 1>&2
        exit 1
        ;;
    esac
done
shift $((OPTIND - 1))

if [ "$#" -gt 0 ]; then
    host=$1
    shift
else
    usage
fi

# shellcheck source=utils.sh
source ./utils.sh

fname=$(gen_gather_filename "$host" "$bridge")

declare -A TABLES

gen_dot_graph_start

flowtable_count=$((0))
table_count=$((0))

while IFS= read -r line; do
    if [[ $line =~ NXST_FLOW ]]; then
        continue
    fi
    line=${line//, / }

    IFS=" " read -r -a field_array <<<"$line"

    declare -A PARSED

    for field in "${field_array[@]}"; do
        if [[ $field =~ ([^=]+)=(.*) ]]; then
            name="${BASH_REMATCH[1]}"
            value="${BASH_REMATCH[2]}"
            if [ -n "${flows_keywords[name]+abc}" ]; then
                printf "Unknown field <%s> in line:\n\t%s" "$name" "$line"
                exit 1
            fi
            PARSED["$name"]="$value"
        else
            printf "Cannot parse line:\n\t%s\n" "$line"
            exit 1
        fi
    done

    if [ -z "${TABLES[${PARSED[table]}]+abc}" ]; then
        if [ $flowtable_count -gt 0 ]; then
            gen_dot_flowtable_end
        fi
        TABLES[${PARSED[table]}]=$((0))
        gen_dot_flowtable_start "${PARSED[table]}"
        table_count=$((table_count + 1))
    else
        TABLES[${PARSED[table]}]=$((TABLES[${PARSED[table]}] + 1))
    fi

    gen_dot_flowtable_rule "${PARSED[priority]}" "${PARSED[actions]}" "${TABLES[${PARSED[table]}]}" "${PARSED[table]}"

    if [ $table_count -ge "$max_tables" ]; then
        break;
    fi

    flowtable_count=$((flowtable_count + 1))

    # if [ $table_count -gt 20 ]; then
    #   exit 1
    # fi
done <"$fname"

if [ $table_count -gt 0 ]; then
    gen_dot_flowtable_end
fi
gen_dot_graph_end
