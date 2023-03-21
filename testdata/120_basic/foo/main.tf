terraform {
  required_providers {
    null = {
      source = "hashicorp/null"
      configuration_aliases = [null.aliased]
    }
  }
}

variable "bar" {
  type = string
}

variable "one" {
  type = string
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
