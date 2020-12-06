package tfjson

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValidateOutput_error(t *testing.T) {
	errOutput := `{
  "valid": false,
  "error_count": 1,
  "warning_count": 0,
  "diagnostics": [
    {
      "severity": "error",
      "summary": "Could not load plugin",
      "detail": "\nPlugin reinitialization required..."
    }
  ]
}`
	var parsed ValidateOutput
	if err := json.Unmarshal([]byte(errOutput), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &ValidateOutput{
		ErrorCount: 1,
		Diagnostics: []Diagnostic{
			{
				Severity: "error",
				Summary:  "Could not load plugin",
				Detail:   "\nPlugin reinitialization required...",
			},
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

func TestValidateOutput_basic(t *testing.T) {
	errOutput := `{
  "valid": false,
  "error_count": 1,
  "warning_count": 1,
  "diagnostics": [
    {
      "severity": "warning",
      "summary": "\"anonymous\": [DEPRECATED] For versions later than 3.0.0, absence of a token enables this mode"
    },
    {
      "severity": "error",
      "summary": "Missing required argument",
      "detail": "The argument \"name\" is required, but no definition was found.",
      "range": {
        "filename": "main.tf",
        "start": {
          "line": 14,
          "column": 37,
          "byte": 200
        },
        "end": {
          "line": 14,
          "column": 37,
          "byte": 200
        }
      }
    }
  ]
}`
	var parsed ValidateOutput
	if err := json.Unmarshal([]byte(errOutput), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &ValidateOutput{
		ErrorCount:   1,
		WarningCount: 1,
		Diagnostics: []Diagnostic{
			{
				Severity: "warning",
				Summary:  "\"anonymous\": [DEPRECATED] For versions later than 3.0.0, absence of a token enables this mode",
			},
			{
				Severity: "error",
				Summary:  "Missing required argument",
				Detail:   "The argument \"name\" is required, but no definition was found.",
				Range: &Range{
					Filename: "main.tf",
					Start: Pos{
						Line:   14,
						Column: 37,
						Byte:   200,
					},
					End: Pos{
						Line:   14,
						Column: 37,
						Byte:   200,
					},
				},
			},
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}
