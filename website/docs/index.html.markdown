---
layout: "petstore"
page_title: "Provider: Petstore"
sidebar_current: "docs-petstore-index"
description: |-
  Manage Pets as Code - with Terraform! An instructional provider for use with chapter 11 of "Terraform in Action"
---

# Petstore Provider

This provider is used to deploy pets as code using the petstore API.

Use the navigation to the left to read about the available resources.

## Versions

For production use, you should constrain the acceptable provider versions via
configuration, to ensure that new versions with breaking changes will not be
automatically installed by `terraform init` in the future:

```hcl
terraform {
  required_providers {
    petstore = {
      source  = "terraform-in-action/petstore"
      version = "~> 1.0"
    }
  }
}
```

For more information on provider installation and constraining provider versions, see the [Provider Requirements documentation](https://www.terraform.io/docs/configuration/provider-requirements.html).

## Example Usage

```hcl
provider "petstore" {
  address  = var.address
}

# Create a pet
resource "petstore_pet" "pet" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `petstore` - (Optional) The Petstore API address to connect to.
  Can be overridden by setting the `PETSTORE_ADDRESS` environment variable.