// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConfigValidate(t *testing.T) {
	f, err := os.Open("testdata/basic/plan.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var plan *Plan
	if err := json.NewDecoder(f).Decode(&plan); err != nil {
		t.Fatal(err)
	}

	if err := plan.Config.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestConfig_actions(t *testing.T) {
	f, err := os.Open("testdata/config_actions/plan.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var plan *Plan
	if err := json.NewDecoder(f).Decode(&plan); err != nil {
		t.Fatal(err)
	}

	if err := plan.Validate(); err != nil {
		t.Fatal(err)
	}

	if plan.Config == nil || plan.Config.RootModule == nil {
		t.Fatal("expected configuration root module to be present")
	}

	expectedRootActions := []*ConfigAction{
		{
			Address:           "action.bufo_print.each",
			Type:              "bufo_print",
			Name:              "each",
			ProviderConfigKey: "bufo",
			ForEachExpression: &Expression{
				ExpressionData: &ExpressionData{
					ConstantValue: UnknownConstantValue,
					References:    []string{"var.names"},
				},
			},
		},
		{
			Address:           "action.bufo_print.many",
			Type:              "bufo_print",
			Name:              "many",
			ProviderConfigKey: "bufo",
			CountExpression: &Expression{
				ExpressionData: &ExpressionData{
					ConstantValue: float64(3),
				},
			},
		},
		{
			Address:           "action.bufo_print.success",
			Type:              "bufo_print",
			Name:              "success",
			ProviderConfigKey: "bufo",
		},
	}

	if diff := cmp.Diff(expectedRootActions, plan.Config.RootModule.Actions); diff != "" {
		t.Fatalf("unexpected root module actions: %s", diff)
	}

	child, ok := plan.Config.RootModule.ModuleCalls["child"]
	if !ok || child.Module == nil {
		t.Fatal("expected child module call to be present")
	}

	expectedChildActions := []*ConfigAction{
		{
			Address:           "action.bufo_print.nested",
			Type:              "bufo_print",
			Name:              "nested",
			ProviderConfigKey: "bufo",
		},
	}

	if diff := cmp.Diff(expectedChildActions, child.Module.Actions); diff != "" {
		t.Fatalf("unexpected child module actions: %s", diff)
	}
}
