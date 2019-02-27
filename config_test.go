package tfjson

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestConfigValidate(t *testing.T) {
	f, err := os.Open("test-fixtures/basic/plan.json")
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

func TestConfigResourceUnmarshalExpressions(t *testing.T) {
	in := []byte(`
{
  "address": "aws_instance.foo",
  "mode": "managed",
  "type": "aws_instance",
  "name": "foo",
  "provider_config_key": "provider.aws",
  "expressions": {
    "ami": {
      "constant_value": "ami-foobar"
    },
    "ebs_block_device": [
      {
        "device_name": {
          "references": [
            "var.foo"
          ]
        }
      }
    ],
    "instance_type": {
      "constant_value": "t2.micro"
    }
  },
  "schema_version": 1
}
`)

	expectedInner := map[string]*Expression{
		"device_name": &Expression{
			ExpressionData: &ExpressionData{
				References: []string{"var.foo"},
			},
		},
	}

	expected := &ConfigResource{
		Address:           "aws_instance.foo",
		Mode:              ManagedResourceMode,
		Type:              "aws_instance",
		Name:              "foo",
		ProviderConfigKey: "provider.aws",
		Expressions: map[string]*Expression{
			"ami": &Expression{
				ExpressionData: &ExpressionData{
					ConstantValue: "ami-foobar",
				},
			},
			"ebs_block_device": &Expression{
				ExpressionData: &ExpressionData{
					NestedBlocks: []map[string]*Expression{expectedInner},
				},
			},
			"instance_type": &Expression{
				ExpressionData: &ExpressionData{
					ConstantValue: "t2.micro",
				},
			},
		},
		SchemaVersion: 1,
	}

	var actual *ConfigResource
	if err := json.Unmarshal(in, &actual); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v", expected, actual)
	}
}

func TestProviderConfigKey(t *testing.T) {
	cases := []struct {
		Name     string
		In       ProviderConfigKey
		Expected *ProviderConfigKeyData
	}{
		{
			Name: "root module, no alias",
			In:   "null",
			Expected: &ProviderConfigKeyData{
				Provider: "null",
			},
		},
		{
			Name: "root module, alias",
			In:   "null.foo",
			Expected: &ProviderConfigKeyData{
				Provider: "null",
				Alias:    "foo",
			},
		},
		{
			Name: "module, no alias",
			In:   "foo:null",
			Expected: &ProviderConfigKeyData{
				ModuleAddress: "foo",
				Provider:      "null",
			},
		},
		{
			Name: "module, alias",
			In:   "foo:null.bar",
			Expected: &ProviderConfigKeyData{
				ModuleAddress: "foo",
				Provider:      "null",
				Alias:         "bar",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := tc.In.Data()
			if !reflect.DeepEqual(tc.Expected, actual) {
				t.Fatalf("expected %#v, got %#v", tc.Expected, actual)
			}

			var out ProviderConfigKey
			out.SetData(actual)
			if tc.In != out {
				t.Fatalf("output data mismatch: expected %q, got %q", tc.In, out)
			}
		})
	}
}
