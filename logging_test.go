// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package tfjson

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-version"
)

var cmpOpts = cmp.AllowUnexported(allLogMessageTypes...)

func TestLogging_generic(t *testing.T) {
	testCases := []struct {
		rawMessage      string
		expectedMessage LogMsg
	}{
		{
			`{"@level":"info","@message":"Installing provider version: hashicorp/aws v6.8.0...","@module":"terraform.ui","@timestamp":"2025-08-11T15:09:18.827459+00:00","type":"log"}`,
			LogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "Installing provider version: hashicorp/aws v6.8.0...",
					Time: time.Date(2025, 8, 11, 15, 9, 18, 827459000, time.UTC),
				},
			},
		},
		{
			`{"@level":"info","@message":"Terraform 1.9.0","@module":"terraform.ui","@timestamp":"2025-08-11T15:09:15.919212+00:00","terraform":"1.9.0","type":"version","ui":"1.2"}`,
			VersionLogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "Terraform 1.9.0",
					Time: time.Date(2025, 8, 11, 15, 9, 15, 919212000, time.UTC),
				},
				Terraform: version.Must(version.NewVersion("1.9.0")),
				UI:        version.Must(version.NewVersion("1.2")),
			},
		},
		{
			`{"@level":"error","@message":"Error: Unclosed configuration block","@module":"terraform.ui","@timestamp":"2025-08-13T10:40:46.749685+00:00","diagnostic":{"severity":"error","summary":"Unclosed configuration block","detail":"There is no closing brace for this block before the end of the file. This may be caused by incorrect brace nesting elsewhere in this file.","range":{"filename":"main.tf","start":{"line":11,"column":30,"byte":153},"end":{"line":11,"column":31,"byte":154}},"snippet":{"context":"resource \"random_pet\" \"name\"","code":"resource \"random_pet\" \"name\" {","start_line":11,"highlight_start_offset":29,"highlight_end_offset":30,"values":[]}},"type":"diagnostic"}`,
			DiagnosticLogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Error,
					Msg:  "Error: Unclosed configuration block",
					Time: time.Date(2025, 8, 13, 10, 40, 46, 749685000, time.UTC),
				},
				Diagnostic: Diagnostic{
					Severity: DiagnosticSeverityError,
					Summary:  "Unclosed configuration block",
					Detail:   "There is no closing brace for this block before the end of the file. This may be caused by incorrect brace nesting elsewhere in this file.",
					Range: &Range{
						Filename: "main.tf",
						Start: Pos{
							Line:   11,
							Column: 30,
							Byte:   153,
						},
						End: Pos{
							Line:   11,
							Column: 31,
							Byte:   154,
						},
					},
					Snippet: &DiagnosticSnippet{
						Context:              ptrToString(`resource "random_pet" "name"`),
						Code:                 `resource "random_pet" "name" {`,
						StartLine:            11,
						HighlightStartOffset: 29,
						HighlightEndOffset:   30,
						Values:               []DiagnosticExpressionValue{},
					},
				},
			},
		},
		{
			`{"@level":"debug","@message":"Foobar","@module":"terraform.ui","@timestamp":"2025-08-11T15:09:18.827459+00:00","type":"FOO"}`,
			UnknownLogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Debug,
					Msg:  "Foobar",
					Time: time.Date(2025, 8, 11, 15, 9, 18, 827459000, time.UTC),
				},
			},
		},
	}

	for _, tc := range testCases {
		msg, err := UnmarshalLogMessage([]byte(tc.rawMessage))
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(tc.expectedMessage, msg, cmpOpts); diff != "" {
			t.Fatalf("unexpected message: %s", diff)
		}
	}
}

func TestLogging_query(t *testing.T) {
	testCases := []struct {
		rawMessage      string
		expectedMessage LogMsg
	}{
		{
			`{"@level":"info","@message":"list.concept_pet.pets: Starting query...","@module":"terraform.ui","@timestamp":"2025-08-28T18:07:11.534006+00:00","list_start":{"address":"list.concept_pet.pets","resource_type":"concept_pet"},"type":"list_start"}`,
			ListStartMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "list.concept_pet.pets: Starting query...",
					Time: time.Date(2025, 8, 28, 18, 7, 11, 534006000, time.UTC),
				},
				ListStart: ListStartData{
					Address:      "list.concept_pet.pets",
					ResourceType: "concept_pet",
					InputConfig:  nil,
				},
			},
		},
		{
			`{"@level":"info","@message":"list.concept_pet.pets: Result found","@module":"terraform.ui","@timestamp":"2025-08-28T18:07:11.534589+00:00","list_resource_found":{"address":"list.concept_pet.pets","display_name":"This is a easy-antelope","identity":{"id":"easy-antelope","legs":6},"resource_type":"concept_pet"},"type":"list_resource_found"}`,
			ListResourceFoundMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "list.concept_pet.pets: Result found",
					Time: time.Date(2025, 8, 28, 18, 7, 11, 534589000, time.UTC),
				},
				ListResourceFound: ListResourceFoundData{
					Address:      "list.concept_pet.pets",
					ResourceType: "concept_pet",
					DisplayName:  "This is a easy-antelope",
					Identity: map[string]json.RawMessage{
						"id":   json.RawMessage(`"easy-antelope"`),
						"legs": json.RawMessage("6"),
					},
				},
			},
		},
		{
			`{"@level":"info","@message":"list.concept_pet.pets: List complete","@module":"terraform.ui","@timestamp":"2025-08-28T18:07:11.534661+00:00","list_complete":{"address":"list.concept_pet.pets","resource_type":"concept_pet","total":5},"type":"list_complete"}`,
			ListCompleteMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "list.concept_pet.pets: List complete",
					Time: time.Date(2025, 8, 28, 18, 7, 11, 534661000, time.UTC),
				},
				ListComplete: ListCompleteData{
					Address:      "list.concept_pet.pets",
					ResourceType: "concept_pet",
					Total:        5,
				},
			},
		},
	}

	for _, tc := range testCases {
		msg, err := UnmarshalLogMessage([]byte(tc.rawMessage))
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(tc.expectedMessage, msg, cmpOpts); diff != "" {
			t.Fatalf("unexpected message: %s", diff)
		}
	}
}

// Includes a typical sequence of logs that happen when initializing a working directory
//
// Currently `init` creates some logs with "type":"log" and others with "type":"init_output"
// Type "init_output" logs include a specific field called "message_code" that takes a string value.
func TestLogging_init(t *testing.T) {
	testCases := []struct {
		rawMessage      string
		expectedMessage LogMsg
	}{
		{
			`{"@level":"info","@message":"Terraform 1.15.0-dev","@module":"terraform.ui","@timestamp":"2025-11-17T17:17:58.540604Z","terraform":"1.15.0-dev","type":"version","ui":"1.2"}`,
			VersionLogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "Terraform 1.15.0-dev",
					Time: time.Date(2025, 11, 17, 17, 17, 58, 540604000, time.UTC),
				},
				Terraform: version.Must(version.NewSemver("1.15.0-dev")),
				UI:        version.Must(version.NewSemver("1.2.0")),
			},
		},
		{
			`{"@level":"info","@message":"Initializing provider plugins found in the configuration...","@module":"terraform.ui","@timestamp":"2025-11-17T17:18:04.314Z","message_code":"initializing_provider_plugin_from_config_message","type":"init_output"}`,
			UnknownLogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "Initializing provider plugins found in the configuration...",
					Time: time.Date(2025, 11, 17, 17, 18, 04, 314000000, time.UTC),
				},
			},
		},
		{
			`{"@level":"info","@message":"hashicorp/aws: Finding latest version...","@module":"terraform.ui","@timestamp":"2025-11-17T17:18:04.314594Z","type":"log"}`,
			LogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "hashicorp/aws: Finding latest version...",
					Time: time.Date(2025, 11, 17, 17, 18, 04, 314594000, time.UTC),
				},
			},
		},
		{
			`{"@level":"info","@message":"Installing provider version: hashicorp/aws v6.21.0...","@module":"terraform.ui","@timestamp":"2025-11-17T17:18:04.784659Z","type":"log"}`,
			LogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "Installing provider version: hashicorp/aws v6.21.0...",
					Time: time.Date(2025, 11, 17, 17, 18, 04, 784659000, time.UTC),
				},
			},
		},
		{
			`{"@level":"info","@message":"Installed provider version: hashicorp/aws v6.21.0 (signed by HashiCorp)","@module":"terraform.ui","@timestamp":"2025-11-17T17:18:26.345919Z","type":"log"}`,
			LogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "Installed provider version: hashicorp/aws v6.21.0 (signed by HashiCorp)",
					Time: time.Date(2025, 11, 17, 17, 18, 26, 345919000, time.UTC),
				},
			},
		},
		{
			`{"@level":"info","@message":"Initializing the backend...","@module":"terraform.ui","@timestamp":"2025-11-17T17:18:52.256Z","message_code":"initializing_backend_message","type":"init_output"}`,
			UnknownLogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "Initializing the backend...",
					Time: time.Date(2025, 11, 17, 17, 18, 52, 256000000, time.UTC),
				},
			},
		},
		// At this point in an init command's output there is a log message that isn't presented in JSON format:
		// /*
		//  Successfully configured the backend "local"! Terraform will automatically
		//  use this backend unless the backend configuration changes.
		// */
		//
		// See this GitHub issue: https://github.com/hashicorp/terraform/issues/37911
		{
			`{"@level":"info","@message":"Terraform has created a lock file .terraform.lock.hcl to record the provider\nselections it made above. Include this file in your version control repository\nso that Terraform can guarantee to make the same selections by default when\nyou run \"terraform init\" in the future.","@module":"terraform.ui","@timestamp":"2025-11-17T17:19:06.698Z","message_code":"lock_info","type":"init_output"}`,
			UnknownLogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "Terraform has created a lock file .terraform.lock.hcl to record the provider\nselections it made above. Include this file in your version control repository\nso that Terraform can guarantee to make the same selections by default when\nyou run \"terraform init\" in the future.",
					Time: time.Date(2025, 11, 17, 17, 19, 06, 698000000, time.UTC),
				},
			},
		},
		{
			`{"@level":"info","@message":"Terraform has been successfully initialized!","@module":"terraform.ui","@timestamp":"2025-11-17T17:19:09.915Z","message_code":"output_init_success_message","type":"init_output"}`,
			UnknownLogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "Terraform has been successfully initialized!",
					Time: time.Date(2025, 11, 17, 17, 19, 9, 915000000, time.UTC),
				},
			},
		},
		{
			`{"@level":"info","@message":"You may now begin working with Terraform. Try running \"terraform plan\" to see\nany changes that are required for your infrastructure. All Terraform commands\nshould now work.\n\nIf you ever set or change modules or backend configuration for Terraform,\nrerun this command to reinitialize your working directory. If you forget, other\ncommands will detect it and remind you to do so if necessary.","@module":"terraform.ui","@timestamp":"2025-11-17T17:19:10.553Z","message_code":"output_init_success_cli_message","type":"init_output"}`,
			UnknownLogMessage{
				baseLogMessage: baseLogMessage{
					Lvl:  Info,
					Msg:  "You may now begin working with Terraform. Try running \"terraform plan\" to see\nany changes that are required for your infrastructure. All Terraform commands\nshould now work.\n\nIf you ever set or change modules or backend configuration for Terraform,\nrerun this command to reinitialize your working directory. If you forget, other\ncommands will detect it and remind you to do so if necessary.",
					Time: time.Date(2025, 11, 17, 17, 19, 10, 553000000, time.UTC),
				},
			},
		},
	}

	for _, tc := range testCases {
		msg, err := UnmarshalLogMessage([]byte(tc.rawMessage))
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(tc.expectedMessage, msg, cmpOpts); diff != "" {
			t.Fatalf("unexpected message: %s", diff)
		}
	}
}
