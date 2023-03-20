# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

provider "null" {
  version = "~> 2.1"
}

resource "null_resource" "foo" {}
