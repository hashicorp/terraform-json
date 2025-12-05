// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const testFixtureDir = "testdata"
const testGoldenPlanFileName = "plan.json"
const testGoldenStateFileName = "state.json"
const testGoldenSchemasFileName = "schemas.json"
const testInvalidDir = "invalid"

func testParse(t *testing.T, filename string, typ reflect.Type) {
	entries, err := os.ReadDir(testFixtureDir)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		t.Run(e.Name(), func(t *testing.T) {
			if e.Name() == testInvalidDir {
				t.Skip("Skipping known invalid test fixture")
			}

			expected, err := os.ReadFile(filepath.Join(testFixtureDir, e.Name(), filename))
			if err != nil {
				if os.IsNotExist(err) {
					t.Skip(err.Error())
				}
				t.Fatal(err)
			}

			parsed := reflect.New(typ).Interface()
			dec := json.NewDecoder(bytes.NewBuffer(expected))
			dec.DisallowUnknownFields()
			if err = dec.Decode(parsed); err != nil {
				t.Fatal(err)
			}

			actual, err := json.Marshal(parsed)
			if err != nil {
				t.Fatal(err)
			}

			// Add a newline at the end
			actual = append(actual, byte('\n'))

			// TODO: Compare the actual struct instead of byte slice
			// because JSON does not guarantee consistent key ordering

			if diff := cmp.Diff(expected, actual); diff != "" {
				t.Fatalf("unexpected: %s", diff)
			}
		})
	}
}

func TestParsePlan(t *testing.T) {
	testParse(t, testGoldenPlanFileName, reflect.TypeOf(Plan{}))
}

func TestParseSchemas(t *testing.T) {
	testParse(t, testGoldenSchemasFileName, reflect.TypeOf(ProviderSchemas{}))
}

func TestParseState(t *testing.T) {
	testParse(t, testGoldenStateFileName, reflect.TypeOf(State{}))
}

func lineAt(text []byte, offs int) []byte {
	i := offs
	for i < len(text) && text[i] != '\n' {
		i++
	}
	return text[offs:i]
}
