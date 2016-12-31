# csvjoin
Join columns from two CSV files without having to use bloated Python modules.

## Compile and Run
Compile with Go:

    go build

And run:

    ./csvjoin  -left ~/left.csv -right ~/right.csv -leftcol 1 -rightcol 2 -trimandignorecase

## Copying

This program uses 2-claus BSD license, text of which can be found in `LICENSE` file.
