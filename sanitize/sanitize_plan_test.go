package sanitize

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/sebdah/goldie"
)

const testDataDir = "testdata"

func TestSanitizePlanGolden(t *testing.T) {
	cases, err := goldenCases()
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range cases {
		t.Run(tc.Name(), testSanitizePlanGoldenEntry(tc))
	}
}

func testSanitizePlanGoldenEntry(c testGoldenCase) func(t *testing.T) {
	return func(t *testing.T) {
		p := new(tfjson.Plan)
		err := json.Unmarshal(c.InputData, p)
		if err != nil {
			t.Fatal(err)
		}

		p, err = SanitizePlan(p)
		if err != nil {
			t.Fatal(err)
		}

		goldie.AssertJson(t, c.Name(), p)
	}
}

type testGoldenCase struct {
	FileName  string
	InputData []byte
}

func (c *testGoldenCase) Name() string {
	return strings.TrimSuffix(c.FileName, filepath.Ext(c.FileName))
}

func goldenCases() ([]testGoldenCase, error) {
	d, err := os.Open(testDataDir)
	if err != nil {
		return nil, err
	}

	entries, err := d.ReadDir(0)
	if err != nil {
		return nil, err
	}

	result := make([]testGoldenCase, 0)
	for _, e := range entries {
		if !e.Type().IsRegular() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}

		data, err := os.ReadFile(filepath.Join(testDataDir, e.Name()))
		if err != nil {
			return nil, err
		}

		result = append(result, testGoldenCase{
			FileName:  e.Name(),
			InputData: data,
		})
	}

	return result, err
}

func init() {
	goldie.FixtureDir = testDataDir
}
