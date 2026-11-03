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

// StateIdentitiesFormatVersionConstraints defines the versions of the JSON
// state identities format that are supported by this package.
var StateIdentitiesFormatVersionConstraints = ">= 1.0, < 2.0"

// StateIdentitiesOutput represents JSON output from terraform state identities -json
type StateIdentitiesOutput struct {
	FormatVersion string                    `json:"format_version"`
	Identities    map[string]map[string]any `json:"identities"`
	Diagnostics   []Diagnostic              `json:"diagnostics,omitempty"`
}

// Validate checks to ensure that data is present, and the
// version matches the version supported by this library.
func (sio *StateIdentitiesOutput) Validate() error {
	if sio == nil {
		return errors.New("output is nil")
	}

	if sio.FormatVersion == "" {
		return errors.New("unexpected input, format version is missing")
	}

	constraint, err := version.NewConstraint(StateIdentitiesFormatVersionConstraints)
	if err != nil {
		return fmt.Errorf("invalid version constraint: %w", err)
	}

	version, err := version.NewVersion(sio.FormatVersion)
	if err != nil {
		return fmt.Errorf("invalid format version %q: %w", sio.FormatVersion, err)
	}

	if !constraint.Check(version) {
		return fmt.Errorf("unsupported state identities format version: %q does not satisfy %q",
			version, constraint)
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler for StateIdentitiesOutput.
//
// As per established convention this method should only ever
// be invoked *indirectly* via [encoding/json] library.
func (sio *StateIdentitiesOutput) UnmarshalJSON(b []byte) error {
	type rawOutput StateIdentitiesOutput
	var output rawOutput

	dec := json.NewDecoder(bytes.NewReader(b))

	// decode numbers as json.Number to avoid losing precision
	dec.UseNumber()

	err := dec.Decode(&output)
	if err != nil {
		return err
	}

	*sio = *(*StateIdentitiesOutput)(&output)

	return sio.Validate()
}
