# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "foo" {
  description = "foobar"
  default     = "bar"
}

variable "number" {
  default = 42
}

variable "map" {
  default = {
    foo    = "bar"
    number = 42
  }
}
