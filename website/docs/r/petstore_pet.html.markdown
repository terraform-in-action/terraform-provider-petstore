---
layout: "petstore"
page_title: "Petstore: petstore_pet"
sidebar_current: "docs-resource-petstore-pet"
description: |-
  Manages pets.
---

# petstore_pet

Manages pets.

## Example Usage

Basic usage:

```hcl
resource "petstore_pet" "pet" {
    name    = "princess"
    species = "cat"
    age     = 3
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the pet.
* `species` - (Required) species of pet.
* `age` - (Required) Age of pet.

## Attributes Reference

* `id` - The id of the pet.

## Import

Pets can be imported; use `<PET ID>` as the import ID. For
example:

```shell
terraform import petstore_pet.pet pet-id
```