package tfjson

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestUnmarshalExpressions(t *testing.T) {
	cases := []struct {
		name     string
		in       string
		expected *ConfigResource
	}{
		{
			name: "basic",
			in: `
{
  "address": "aws_instance.foo",
  "mode": "managed",
  "type": "aws_instance",
  "name": "foo",
  "provider_config_key": "aws",
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
`,
			expected: &ConfigResource{
				Address:           "aws_instance.foo",
				Mode:              ManagedResourceMode,
				Type:              "aws_instance",
				Name:              "foo",
				ProviderConfigKey: "aws",
				Expressions: map[string]*Expression{
					"ami": {
						ExpressionData: &ExpressionData{
							ConstantValue: "ami-foobar",
						},
					},
					"ebs_block_device": {
						ExpressionData: &ExpressionData{
							NestedBlocks: []map[string]*Expression{
								{
									"device_name": {
										ExpressionData: &ExpressionData{
											ConstantValue: UnknownConstantValue,
											References:    []string{"var.foo"},
										},
									},
								},
							},
						},
					},
					"instance_type": {
						ExpressionData: &ExpressionData{
							ConstantValue: "t2.micro",
						},
					},
				},
				SchemaVersion: 1,
			},
		},
		{
			name: "explicit null in contstant value",
			in: `
{
  "address": "null_resource.foo",
  "mode": "managed",
  "type": "null_resource",
  "name": "foo",
  "provider_config_key": "null",
  "expressions": {
    "triggers": {
      "constant_value": {
        "foo": null
      }
    }
  },
  "schema_version": 0
}
`,
			expected: &ConfigResource{
				Address:           "null_resource.foo",
				Mode:              ManagedResourceMode,
				Type:              "null_resource",
				Name:              "foo",
				ProviderConfigKey: "null",
				Expressions: map[string]*Expression{
					"triggers": {
						ExpressionData: &ExpressionData{
							ConstantValue: map[string]interface{}{
								"foo": nil,
							},
						},
					},
				},
				SchemaVersion: 0,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var actual *ConfigResource
			if err := json.Unmarshal([]byte(tc.in), &actual); err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.expected, actual) {
				t.Fatalf("expected:\n\n%s\n\ngot:\n\n%s\n\n", spew.Sdump(tc.expected), spew.Sdump(actual))
			}
		})
	}
}
