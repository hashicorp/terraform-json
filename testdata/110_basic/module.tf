module "foo" {
  source = "./foo"

  bar = "baz"
  one = "two"

  providers = {
    null.aliased = null
  }
}
