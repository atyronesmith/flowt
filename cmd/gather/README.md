# gather

This command automates the acquisition of the nort and southbound databases either from an OpenStack deployment, or from an ovn-scale-test deployment.

## Example

`go run cmd/gather/main.go -nb -sb -o test osp 192.168.25.30:22`

## Usage

After the database(s) have been retrieve other commands in the repo can be use to generate go datastructures, information about the database, or scripts to recreate the databases using **ovn-scale-test**. .
