# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "one" {
  type = "string"
}

resource "null_resource" "foo" {
  triggers = {
    foo = null
  }
}

resource "null_resource" "bar" {
  triggers = {
    foo = var.one
  }
}

resource "null_resource" "baz" {
  triggers = {
    foo = var.one
    bar = null
  }
}
