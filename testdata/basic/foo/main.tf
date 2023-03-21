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
