package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	tfjson "github.com/hashicorp/terraform-json"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s FILE\n\n", os.Args[0])
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var parsed tfjson.Plan
	if err = json.Unmarshal(data, &parsed); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	out, err := json.MarshalIndent(parsed, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	out = append(out, byte('\n'))
	os.Stdout.Write(out)
}
