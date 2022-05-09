resource "null_resource" "foo" {
  triggers = {
    foo = "bar"
  }

  provisioner "local-exec" {
    command = "echo hello"
  }
}

resource "null_resource" "bar" {
  triggers = {
    foo_id = "${null_resource.foo.id}"
  }
}

resource "null_resource" "baz" {
  count = 3

  triggers = {
    foo_id = "${null_resource.foo.id}"
  }
}
