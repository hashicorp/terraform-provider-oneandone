---
layout: "oneandone"
page_title: "1&1: oneandone_block_storage"
sidebar_current: "docs-oneandone-resource-block-storage"
description: |-
  Creates and manages 1&1 Block Storage.
---

# oneandone\_block\_storage

Manages a Block Storage on 1&1

## Example Usage

```hcl
resource "oneandone_block_storage" "storage" {
  name = "test_blk_storage1"
  description = "testing_blk_storage"
  size = 20
  datacenter = "US"
}
```

## Argument Reference

The following arguments are supported:

* `datacenter` - (Optional) Location of desired 1and1 datacenter, where the block storage will be created. Can be `DE`, `GB`, `US` or `ES`
* `description` - (Optional) Description for the block storage
* `name` - (Required) The name of the storage
* `server_id` - (Optional) ID of the server that the block storage will be attached to
* `size` - (Required) Size of the block storage (`min: 20, max: 500, multipleOf: 10`)
