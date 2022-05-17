# flowt

Debugging tools for openflow/OVN/OVS

## Commands

```text
kni@rhos-nfv-01 flowt]$ go run cmd/schema/main.go 
Generate an ./ovn_nortbound.go or ./ovn_southbound.go file depending on
the type of database contained in the file_to_parse.  The ***.go file
contains go structures that map the database data contents.

Usage: main [options] file_to_parse
       file_to_parse  -- Path to flow rules.
  -h Print usage information.
  -o string
     Go Schema file output file (Defaults to name in schema in current directory)
  -pkg string
     Target package for schema. (default "dbtypes")
  -v Print extra runtime information.
```

```text
[kni@rhos-nfv-01 flowt]$ go run cmd/db/main.go -h
Read an OVN NB or SB database, generate a json file of just the data,
and generate a .dot file that represents the schema of the database.

Usage: main [options] db_to_parse
       db_to_parse  -- Path to NB or SB database.
  -c string
     Name of chart to generate.
  -chart string
     Name of chart to generate.
  -help
     Print usage information.
  -o string
     Directory to place the results (Defaults to local directory) (default ".")
  -v Print extra runtime information.
```
