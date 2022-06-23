# docs

Generate a json-based tooltip file from xml documenation for Northbound or Southbound OVN database.  The xml documentation is taken directly from the upstream repository [ovn-nb.xml](https://github.com/ovn-org/ovn/blob/main/ovn-nb.xml).

The generated JSON file is used to add tooltips to the visual schema for a Northbound or Southbound OVN database.  The visual schema is in the form a '.dot' file that can be used to generate a picture file (svg,jpg,etc...).

## Example

`go run cmd/docs/main.go examples/ovn-nb.xml > cmd/db/data/nb_table_tooltip.json`

The *nb_table_tooltip.json* file is used by the cmd/db/data/main.go tool to generate the visual schema.
