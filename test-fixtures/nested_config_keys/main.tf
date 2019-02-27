provider "aws" {
  region = "us-east-1"
}

variable "foo" {
  default = "/dev/sda1"
}

resource "aws_instance" "foo" {
  ami           = "ami-foobar"
  instance_type = "t2.micro"

  ebs_block_device {
    device_name = "${var.foo}"
  }
}
