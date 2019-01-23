provider "null" {}

provider "aws" {}

provider "aws" {
  alias  = "east"
  region = "us-east-1"
}
