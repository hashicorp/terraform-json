terraform {
  required_providers {
    null = {
      source = "hashicorp/null"
    }
  }
}
provider "aws" {
  region = "us-west-2"
}

provider "aws" {
  alias  = "east"
  region = "us-east-1"
}
