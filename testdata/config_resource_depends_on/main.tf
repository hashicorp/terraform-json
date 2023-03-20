# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "null_resource" "foo" {}

resource "null_resource" "bar" {
  depends_on = ["null_resource.foo"]
}
