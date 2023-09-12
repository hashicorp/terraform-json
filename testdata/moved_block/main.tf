moved {
  from = random_id.test
  to = random_id.test2
}

resource "random_id" "test2" {
  byte_length = 10
}
