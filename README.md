# stdcsv

stdcsv is a tool and a Golang package to transform "non standard" CSV data to a "standard" one AKA: Comma separated, correctly quoted, UTF-8 charset.

## Install the cli tool

```sh
go install github.com/colbee1/stdcsv
```

### Usage

```
stdcsv -h
Usage of stdcsv:
  -charset string
        charset of input (default "utf-8")
  -charset-list
        list recognized charset
  -colpad string
        column content when padding
  -columns int
        number of columns per rows (default -1)
  -comma-in string
        comma used in input (default ",")
  -comma-out string
        comma to use for output (default ",")
  -comment string
        char used as comment (default "#")
  -headers string
        headers separated by comma
  -in string
        input file (default "-")
  -lazy
        lazy quoting is used in input
  -limit int
        maximum number of rows in output
  -offset int
        number of rows to skip before start
  -out string
        output file (default "-")
  -quiet
        don't be to verbose
  -skip-bom
        skip the BOM
  -skip-trailing-comma
        skip the trailing comma if present
  -trim-sapces
        trim cells spaces
```
