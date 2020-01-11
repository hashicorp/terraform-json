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
