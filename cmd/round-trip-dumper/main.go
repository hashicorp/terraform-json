package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"

	tfjson "github.com/hashicorp/terraform-json"
)

var (
	diff   = flag.Bool("diff", false, "diff output instead of writing")
	schema = flag.Bool("schema", false, "input is a schema, not a plan")
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "usage: %s FILE\n\n", os.Args[0])
		os.Exit(1)
	}

	path := flag.Arg(0)

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer f.Close()

	var parsed interface{}
	if *schema {
		parsed = &tfjson.ProviderSchemas{}
	} else {
		parsed = &tfjson.Plan{}
	}

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()
	if err = dec.Decode(parsed); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	out, err := json.MarshalIndent(parsed, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	out = append(out, byte('\n'))

	if *diff {
		var diffCmd string
		if _, err := exec.LookPath("colordiff"); err == nil {
			diffCmd = "colordiff"
		} else {
			diffCmd = "diff"
		}

		cmd := exec.Command(diffCmd, "-urN", path, "-")
		cmd.Stdin = bytes.NewBuffer(out)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			if err.(*exec.ExitError).ProcessState.ExitCode() > 1 {
				os.Exit(1)
			}
		} else {
			fmt.Fprintln(os.Stderr, "[no diff]")
		}
	} else {
		os.Stdout.Write(out)
	}
}
