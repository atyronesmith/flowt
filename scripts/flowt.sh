#!/bin/bash

SCRIPT_DIR=./cmds

declare -A cmds=(
    [gather]="${SCRIPT_DIR}/gather.sh"
)

usage() {
    local out_dir="$1"

    prog=$(basename "$0")
    cat <<-EOM
    Analyize OVS/OpenFlow/OVN
    Usage:
        $prog [-h] [-d] [-v] gather
            gather           -- Gather OVS/OVN/OpenFlow information from nodes in inventory
            
    Options
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

while getopts ":hvd" opt; do
    case ${opt} in
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
    COMMAND=$1
    shift
else
    usage
fi

case "$COMMAND" in
# Parse options to the install sub command
gather)
    if [ "$#" -lt 1 ]; then
        usage
    fi
    deploy "$@"
    ;;
destroy)
    if [ "$#" -lt 1 ]; then
        usage
    fi
    destroy "$@"
    ;;
create-deploy)
    create_deploy
    ;;
prep-osp)
    prepare_openstack
    ;;
test-sriov)
    test_sriov
    ;;
update-clouds)
    update_clouds
    ;;
test-vlan)
    test_vlan
    ;;
patch-ocp)
    patch_ocp
    ;;
prep-ocp)
    prepare_for_ocp_worker
    ;;
csr)
    sign_csr
    ;;
clean)
    rm -rf "$BUILD_DIR"
    ;;
label-nodes)
    label_nodes feature.node.kubernetes.io/network-sriov.capable="true"
    ;;
deploy-operator)
    if [ "$#" -lt 1 ]; then
        usage
    fi
    deploy_operator "$1"
    ;;
deploy-policy)
    deploy_policy
    ;;
pull-secret)
    if [ "$#" -lt 1 ]; then
        usage
    fi
    create_pull_secret "$1"
    ;;
install-manifests)
    install_manifests
    ;;
ingress-fip)
    create_ingress_fip
    ;;
print-config)
    print_install_config
    ;;
*)
    echo "Unknown command: $COMMAND"
    usage "$out_dir"
    ;;
esac
