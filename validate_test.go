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

func TestValidateOutput_versioned(t *testing.T) {
	errOutput := `{
  "format_version": "0.1",
  "valid": false,
  "error_count": 1,
  "warning_count": 1,
  "diagnostics": [
    {
      "severity": "warning",
      "summary": "Deprecated Attribute",
      "detail": "Deprecated in favor of project_id",
      "range": {
        "filename": "main.tf",
        "start": {
          "line": 21,
          "column": 25,
          "byte": 408
        },
        "end": {
          "line": 21,
          "column": 42,
          "byte": 425
        }
      },
      "snippet": {
        "context": "resource \"google_project_access_approval_settings\" \"project_access_approval\"",
        "code": "  project             = \"my-project-name\"",
        "start_line": 21,
        "highlight_start_offset": 24,
        "highlight_end_offset": 41,
        "values": []
      }
    },
    {
      "severity": "error",
      "summary": "Missing required argument",
      "detail": "The argument \"enrolled_services\" is required, but no definition was found.",
      "range": {
        "filename": "main.tf",
        "start": {
          "line": 19,
          "column": 78,
          "byte": 340
        },
        "end": {
          "line": 19,
          "column": 79,
          "byte": 341
        }
      },
      "snippet": {
        "context": "resource \"google_project_access_approval_settings\" \"project_access_approval\"",
        "code": "resource \"google_project_access_approval_settings\" \"project_access_approval\" {",
        "start_line": 19,
        "highlight_start_offset": 77,
        "highlight_end_offset": 78,
        "values": []
      }
    }
  ]
}`
	var parsed ValidateOutput
	if err := json.Unmarshal([]byte(errOutput), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &ValidateOutput{
		FormatVersion: "0.1",
		ErrorCount:    1,
		WarningCount:  1,
		Diagnostics: []Diagnostic{
			{
				Severity: "warning",
				Summary:  "Deprecated Attribute",
				Detail:   "Deprecated in favor of project_id",
				Range: &Range{
					Filename: "main.tf",
					Start:    Pos{Line: 21, Column: 25, Byte: 408},
					End:      Pos{Line: 21, Column: 42, Byte: 425},
				},
				Snippet: &DiagnosticSnippet{
					Context:              ptrToString(`resource "google_project_access_approval_settings" "project_access_approval"`),
					Code:                 `  project             = "my-project-name"`,
					StartLine:            21,
					HighlightStartOffset: 24,
					HighlightEndOffset:   41,
					Values:               []DiagnosticExpressionValue{},
				},
			},
			{
				Severity: "error",
				Summary:  "Missing required argument",
				Detail:   `The argument "enrolled_services" is required, but no definition was found.`,
				Range: &Range{
					Filename: "main.tf",
					Start:    Pos{Line: 19, Column: 78, Byte: 340},
					End:      Pos{Line: 19, Column: 79, Byte: 341},
				},
				Snippet: &DiagnosticSnippet{
					Context:              ptrToString(`resource "google_project_access_approval_settings" "project_access_approval"`),
					Code:                 `resource "google_project_access_approval_settings" "project_access_approval" {`,
					StartLine:            19,
					HighlightStartOffset: 77,
					HighlightEndOffset:   78,
					Values:               []DiagnosticExpressionValue{},
				},
			},
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

func TestValidateOutput_versioned10(t *testing.T) {
	errOutput := `{
  "format_version": "1.0",
  "valid": false,
  "error_count": 1,
  "warning_count": 1,
  "diagnostics": [
    {
      "severity": "warning",
      "summary": "Deprecated Attribute",
      "detail": "Deprecated in favor of project_id",
      "range": {
        "filename": "main.tf",
        "start": {
          "line": 21,
          "column": 25,
          "byte": 408
        },
        "end": {
          "line": 21,
          "column": 42,
          "byte": 425
        }
      },
      "snippet": {
        "context": "resource \"google_project_access_approval_settings\" \"project_access_approval\"",
        "code": "  project             = \"my-project-name\"",
        "start_line": 21,
        "highlight_start_offset": 24,
        "highlight_end_offset": 41,
        "values": []
      }
    },
    {
      "severity": "error",
      "summary": "Missing required argument",
      "detail": "The argument \"enrolled_services\" is required, but no definition was found.",
      "range": {
        "filename": "main.tf",
        "start": {
          "line": 19,
          "column": 78,
          "byte": 340
        },
        "end": {
          "line": 19,
          "column": 79,
          "byte": 341
        }
      },
      "snippet": {
        "context": "resource \"google_project_access_approval_settings\" \"project_access_approval\"",
        "code": "resource \"google_project_access_approval_settings\" \"project_access_approval\" {",
        "start_line": 19,
        "highlight_start_offset": 77,
        "highlight_end_offset": 78,
        "values": []
      }
    }
  ]
}`
	var parsed ValidateOutput
	if err := json.Unmarshal([]byte(errOutput), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &ValidateOutput{
		FormatVersion: "1.0",
		ErrorCount:    1,
		WarningCount:  1,
		Diagnostics: []Diagnostic{
			{
				Severity: "warning",
				Summary:  "Deprecated Attribute",
				Detail:   "Deprecated in favor of project_id",
				Range: &Range{
					Filename: "main.tf",
					Start:    Pos{Line: 21, Column: 25, Byte: 408},
					End:      Pos{Line: 21, Column: 42, Byte: 425},
				},
				Snippet: &DiagnosticSnippet{
					Context:              ptrToString(`resource "google_project_access_approval_settings" "project_access_approval"`),
					Code:                 `  project             = "my-project-name"`,
					StartLine:            21,
					HighlightStartOffset: 24,
					HighlightEndOffset:   41,
					Values:               []DiagnosticExpressionValue{},
				},
			},
			{
				Severity: "error",
				Summary:  "Missing required argument",
				Detail:   `The argument "enrolled_services" is required, but no definition was found.`,
				Range: &Range{
					Filename: "main.tf",
					Start:    Pos{Line: 19, Column: 78, Byte: 340},
					End:      Pos{Line: 19, Column: 79, Byte: 341},
				},
				Snippet: &DiagnosticSnippet{
					Context:              ptrToString(`resource "google_project_access_approval_settings" "project_access_approval"`),
					Code:                 `resource "google_project_access_approval_settings" "project_access_approval" {`,
					StartLine:            19,
					HighlightStartOffset: 77,
					HighlightEndOffset:   78,
					Values:               []DiagnosticExpressionValue{},
				},
			},
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

func ptrToString(val string) *string {
	return &val
}
