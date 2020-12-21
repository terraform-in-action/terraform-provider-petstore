

provider "aws" {
    profile = "swinkler"
    region = "us-west-2"
}

module "petstore" {
  source  = "scottwinkler/petstore/aws"
}

output "address" {
  value = module.petstore.address
}
