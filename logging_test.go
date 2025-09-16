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
			`{"@level":"info","@message":"list.concept_pet.pets: Result found","@module":"terraform.ui","@timestamp":"2025-08-28T18:07:11.534589+00:00","list_resource_found":{"address":"list.concept_pet.pets","display_name":"This is a easy-antelope","identity":{"id":"easy-antelope","legs":6},"identity_version":1,"resource_type":"concept_pet"},"type":"list_resource_found"}`,
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
					IdentityVersion: 1,
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
