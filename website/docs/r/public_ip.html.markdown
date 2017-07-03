---
layout: "oneandone"
page_title: "1&1: oneandone_public_ip"
sidebar_current: "docs-oneandone-resource-public-ip"
description: |-
  Creates and manages 1&1 Public IP.
---

# oneandone\_ip

Manages a Public IP on 1&1

## Example Usage

```hcl
resource "oneandone_public_ip" "ip" {
	"ip_type"     = "IPV4"
	"reverse_dns" = "%s"
	"datacenter"  = "GB"
}
```

## Argument Reference

The following arguments are supported:

* `ip_type` - (Required) IP type. Can be `IPV4` or `IPV6`
* `reverse_dns` - (Optional) 
* `datacenter` - (Optional) Location of desired 1and1 datacenter. Can be `DE`, `GB`, `US` or `ES`.
* `ip_address` - (Computed) The IP address.
