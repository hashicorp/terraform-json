// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/go-version"
)

// ModulesFormatVersionConstraints defines the versions of
// the JSON format that are supported by this package.
var ModulesFormatVersionConstraints = ">= 1.0, < 2.0"

// ModulesOutput represents JSON output from terraform modules -json
// (available from 1.10 onwards)
type ModulesOutput struct {
	FormatVersion string         `json:"format_version"`
	Modules       []ModuleOutput `json:"modules"`
	Diagnostics   []Diagnostic   `json:"diagnostics,omitempty"`
}

// ModuleOutput represents a single module entry in the list of modules
type ModuleOutput struct {
	Key     string `json:"key"`
	Source  string `json:"source"`
	Version string `json:"version"`
}

// Validate checks to ensure that data is present, and the
// version matches the version supported by this library.
func (mo *ModulesOutput) Validate() error {
	if mo == nil {
		return errors.New("output is nil")
	}

	if mo.FormatVersion == "" {
		return errors.New("unexpected input, format version is missing")
	}

	constraint, err := version.NewConstraint(ModulesFormatVersionConstraints)
	if err != nil {
		return fmt.Errorf("invalid version constraint: %w", err)
	}

	version, err := version.NewVersion(mo.FormatVersion)
	if err != nil {
		return fmt.Errorf("invalid format version %q: %w", mo.FormatVersion, err)
	}

	if !constraint.Check(version) {
		return fmt.Errorf("unsupported modules format version: %q does not satisfy %q",
			version, constraint)
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler for ModulesOutput.
//
// As per established convention this method should only ever
// be invoked *indirectly* via [encoding/json] library.
func (mo *ModulesOutput) UnmarshalJSON(b []byte) error {
	type rawOutput ModulesOutput
	var output rawOutput

	dec := json.NewDecoder(bytes.NewReader(b))
	err := dec.Decode(&output)
	if err != nil {
		return err
	}

	*mo = *(*ModulesOutput)(&output)

	return mo.Validate()
}
