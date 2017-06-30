---
layout: "oneandone"
page_title: "1&1: oneandone_vpn"
sidebar_current: "docs-oneandone-resource-vpn"
description: |-
  Creates and manages 1&1 VPN.
---

# oneandone\_vpn

Manages a VPN on 1&1

## Example Usage

```hcl
resource "oneandone_vpn" "vpn" {
  datacenter  = "GB"
  name        = "%s"
  description = "ttest descr"
}
```

## Argument Reference

The following arguments are supported:

* `datacenter` - (Optional) Location of desired 1and1 datacenter. Can be `DE`, `GB`, `US` or `ES`.
* `name` - (Required) The name of the VPN
* `description` - (Optional)
* `download_path` - (Optional)
* `file_name` - (Optional)

