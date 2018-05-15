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
resource "oneandone_image" "example1" {
  name = "example1"
  server_id = "932F8ABA5060571E5D3C2119E0E31360"
  description = "Weekly server snapshot"
  frequency = "WEEKLY"
  num_images = 5
}
```
```hcl
resource "oneandone_image" "example2" {
  name = "example2"
  description = "Custom imported image"
  datacenter = "DE"
  os_id = "B77E19E062D5818532EFF11C747BD104"
  source = "image"
  url = "https://example.net/image.vdi"
}
```

## Argument Reference

The following arguments are supported:

* `datacenter` - (Optional) Country code of the datacenter where the image will be created (`US`, `DE`, `GB`, and `ES`).
* `description` - (Optional) Image description.
* `frequency` - (Optional) Creation policy frequency. Frecuency policy is only allowed in default datacenter. (`ONCE`, `DAILY`, `WEEKLY`)
* `name` - (Required) The name of the image.
* `num_images` - (Optional) Maximum number of images. Required when image is created with frequency policy.
* `os_id` - (Optional) ID of the Operating System to import.
* `server_id` - (Optional) Server ID - Required when image `source` is `server`.
* `source` - (Optional) Source of the new image: `server` (from an existing server), `image` (from an imported image) or `iso` (from an imported iso).
* `type` - (Optional) Type of the ISO to import: `os` (Operating System) or `app` (Application). It is required when the source is iso.
* `url` - (Optional) URL where the image can be downloaded. It is required when the source is image or iso.