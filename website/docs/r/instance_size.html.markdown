---
layout: "oneandone"
page_title: "1&1: oneandone_instance_size"
sidebar_current: "docs-oneandone-datasource-instance-size"
description: |-
  Fetches a predefined instance type for 1&1 servers.
---

# oneandone\_instance\_size

Fetches a predefined instance type for 1&1 servers

## Example Usage

```hcl
data "oneandone_instance_size" "sizeByName" {
  name = "L"
}

data "oneandone_instance_size" "sizeByHardware" {
  vcores = 2
  ram = 4
}

resource "oneandone_server" "server" {
  name                = "Example"
  image               = "debian8-64min"
  datacenter          = "DE"
  fixed_instance_size = "${data.oneandone_instance_size.sizeByName.id}"
  ...
}
```

## Argument Reference

The following arguments are supported, at least one is required:

* `name` -(Optional) Number of cores per processor
* `ram` - (Optional) Size of ram in GB
* `vcores` - (Optional)  Number of vcores

It exposes the following attributes

* `coresPerProcessor` - (Computed) The number of vcores per processor
* `id` - (Computed) The ID of the instance type
* `name` - (Computed) The Name of the instance type
* `ram` - (Computed) The size of the ram in GB
* `vcores` - (Computed) The number of vcores
