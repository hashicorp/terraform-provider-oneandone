---
layout: "oneandone"
page_title: "1&1: oneandone_image"
sidebar_current: "docs-oneandone-resource-image"
description: |-
  Creates and manages 1&1 Image.
---

# oneandone\_image

Manages Images on 1&1

## Example Usage

```hcl
resource "oneandone_image" "img" {
  name = "test"
  server_id = "C72CF0A681B0CCE7EC624DD194D585C6",
  description = "Testing terraform 1and1 image create",
  frequency = "WEEKLY",
  num_images = 5
}
```

## Argument Reference

The following arguments are supported:

* `datacenter_id` - (Optional) ID of the datacenter where the image will be created.
* `description` - (Optional) Image description.
* `frequency` - (Optional) Creation policy frequency. Frecuency policy is only allowed in default datacenter. (`ONCE`, `DAILY`, `WEEKLY`)
* `name` - (Required) The name of the image.
* `num_images` - (Optional) Maximum number of images. Required when image is created with frequency policy.
* `os_id` - (Optional) ID of the Operating System to import.
* `server_id` - (Required) Server ID.
* `source` - (Optional) Source of the new image: `server` (from an existing server), `image` (from an imported image) or `iso` (from an imported iso).
* `type` - (Optional) Type of the ISO to import: `os` (Operating System) or `app` (Application). It is required when the source is iso.
* `url` - (Optional) URL where the image can be downloaded. It is required when the source is image or iso.