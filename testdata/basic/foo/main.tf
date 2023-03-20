# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "bar" {
  type = "string"
}

variable "one" {
  type = "string"
}

provider "null" {
  alias = "aliased"
}

resource "null_resource" "foo" {
  triggers = {
    foo = "bar"
  }
}

resource "null_resource" "aliased" {
  provider = "null.aliased"
}

output "foo" {
  value = "bar"
}
