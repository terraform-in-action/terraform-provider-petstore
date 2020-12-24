---
layout: "petstore"
page_title: "Petstore: petstore_pet_ids"
sidebar_current: "docs-datasource-petstore-pet-ids"
description: |-
  Get IDs for pet resources
---

# Data Source: petstore_pet_ids

Use this data source to get a list of pet IDs

## Example Usage

```hcl
data "petstore_pet_ids" "pets" {
    names = ["Snowball", "Princess"]
}

data "petstore_pet_ids" "all" {
    names = ["*"]
}
```

## Argument Reference

The following arguments are supported:

* `names` - (Required) A list of pet names to search for. Names that don't
  match a real pet will be omitted from the results, but are not an error.

    To select _all_ pets, provide a list with a single
    asterisk, like `["*"]`. No other use of wildcards are supported.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `ids` - A list of pet IDs