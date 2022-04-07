#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

SSH_USER="heat-admin"

usage() {
    prog=$(basename "$0")
    cat <<-EOM
    Gather flows from a node/bridge
    Usage:
        $prog [-h] [-d] [-v] [-b bridge_name] host
            ip_address           -- hostname or IP Address of the node to query
            
    Options
            -b  -- The name of the bridge to query. Defaults to br-int.
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

while getopts ":hvdb:" opt; do
    case ${opt} in
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
source "$SCRIPT_DIR/utils.sh"

LOGIN="$SSH_USER@$host"

flows=$(ssh "$LOGIN" "sudo ovs-ofctl dump-flows $bridge") ||
    (
        printf "Error gathering flows from %s!\n" "$LOGIN"
        exit 1
    )

fname=$(gen_gather_filename "$host" "$bridge")

cat <<< "$flows" > "$fname"