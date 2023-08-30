variable "file_names" {
  type = set(string)
  default = [
    "file1.txt",
  ]
}

resource "local_file" "foo" {
  for_each = var.file_names
  content  = "Hello, World!"
  filename = each.value

  lifecycle {
    postcondition {
      condition     = self.content == "Hello, World!"
      error_message = "File content is not correct"
    }
  }
}

