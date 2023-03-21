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
