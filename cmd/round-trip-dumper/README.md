# Round-Trip Dumper/Diagnostic Tool

This directory contains a simple tool that will load a plan JSON file, and then
immediately dump it back out to stdout. It's helpful when troubleshooting large
parsing errors, which should (hopefully) be rare.

`go build ./` in this directory to build the binary.

## Shortcuts

Quickly diff the basic test fixture against what tfjson is actually outputting
for it:

```
./round-trip-dumper ../../test-fixtures/basic/plan.json | diff -urN ../../test-fixtures/basic/plan.json - | less -R
```

You can also use [`colordiff`](https://www.colordiff.org/) to get colorized
output which may help visualize things.
