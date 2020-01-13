resource "null_resource" "foo" {}

resource "null_resource" "bar" {
  depends_on = ["null_resource.foo"]
}
