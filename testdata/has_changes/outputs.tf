output "foo" {
  sensitive = true
  value     = "bar"
}

output "string" {
  value = "foo"
}

output "list" {
  value = [
    "foo",
    "bar",
  ]
}

output "map" {
  value = {
    foo    = "bar"
    number = 42
  }
}

output "referenced" {
  value = null_resource.foo.id
}

output "interpolated" {
  value = "${null_resource.foo.id}"
}

output "referenced_deep" {
  value = {
    foo    = "bar"
    number = 42
    map = {
      bar = "baz"
      id = null_resource.foo.id
    }
  }
}

output "interpolated_deep" {
  value = {
    foo    = "bar"
    number = 42
    map = {
      bar = "baz"
      id = "${null_resource.foo.id}"
    }
  }
}
