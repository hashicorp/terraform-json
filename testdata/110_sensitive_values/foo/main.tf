# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "bar" {
  type = string
}

variable "one" {
  type = string
}

terraform {
  required_providers {
    null = {
      source                = "hashicorp/null"
      configuration_aliases = [null.aliased]
    }
  }
}

resource "null_resource" "foo" {
  triggers = {
    foo = "bar"
  }
}

resource "null_resource" "aliased" {
  provider = null.aliased
}

output "foo" {
  value = "bar"
}
