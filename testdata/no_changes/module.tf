# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

module "foo" {
  source = "./foo"

  bar = "baz"
  one = "two"
}
