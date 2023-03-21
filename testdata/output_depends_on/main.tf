resource "null_resource" "foo" {}

resource "null_resource" "bar" {}

output "id" {
  depends_on = ["null_resource.bar"]
  value      = "${null_resource.foo.id}"
}
