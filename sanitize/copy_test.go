package sanitize

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
)

func TestCopyStructureCopy(t *testing.T) {
	in := tfjson.UnknownConstantValue
	out, err := copyStructureCopy(in)
	if err != nil {
		t.Fatal(err)
	}

	if in != out {
		t.Fatal("did not shallow copy")
	}
}
