# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "bar" {
  type = "string"
}

variable "one" {
  type = "string"
}

resource "null_resource" "foo" {
  triggers = {
    foo = "bar"
  }
}
