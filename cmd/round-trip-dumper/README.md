# Round-Trip Dumper/Diagnostic Tool

This directory contains a simple tool that will load a plan JSON file, and then
immediately dump it back out to stdout. It's helpful when troubleshooting large
parsing errors, which should (hopefully) be rare.

`go build ./` in this directory to build the binary. `go run` also works if you
don't need the binary permanently.

## Diffing

The `-diff` flag will automatically diff the result for you. `colordiff` is used
if it's present, otherwise regular `diff` is used.
