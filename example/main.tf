terraform {
  required_providers {
    petstore = {
      source  = "terraform-in-action/petstore"
      version = "~> 1.0"
    }
  }
}

provider "petstore" {
  address = "https://xfc42dh782.execute-api.us-west-2.amazonaws.com/v1"
}

resource "petstore_pet" "my_pet" {
  name    = "princess"
  species = "cat"
  age     = 3
}
