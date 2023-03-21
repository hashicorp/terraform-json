module "foo" {
    source = "vancluever/module/null"

    depends_on = [
        null_resource.bar
    ]
}

resource "null_resource" "bar" {}