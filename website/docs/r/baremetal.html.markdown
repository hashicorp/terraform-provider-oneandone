---
layout: "oneandone"
page_title: "1&1: oneandone_baremetal"
sidebar_current: "docs-oneandone-resource-server-baremetal"
description: |-
  Creates and manages 1&1 Baremetal Server.
---

# oneandone\_baremetal

Manages a Baremetal Server on 1&1

## Example Usage

```hcl
resource "oneandone_baremetal" "server" {
  name = "%s"
  description = "%s"
  image = "%s"
  password = "Kv40kd8PQb"
  datacenter = "US"
  baremetal_model_id = "%s"
  ssh_key_path = "/path/to/private/ssh_key"
  ssh_key_public = "${file("/path/to/public/key.pub")}"


  provisioner "remote-exec" {
    inline = [
      "apt-get update",
      "apt-get -y install nginx",
    ]
  }
}
```

## Argument Reference

The following arguments are supported:

* `datacenter` - (Optional) Location of desired 1and1 datacenter. Can be `DE`, `GB`, `US` or `ES`
* `description` - (Optional) Description of the server
* `firewall_policy_id` - (Optional) ID of firewall policy
* `baremetal_model_id` - (Required) ID of a baremetal model
* `image` -(Required) The name of a desired image to be provisioned with the server
* `ip` - (Optional) IP address for the server
* `loadbalancer_id` - (Optional) ID of the load balancer
* `monitoring_policy_id` - (Optional) ID of monitoring policy
* `name` -(Required) The name of the server.
* `password` - (Optional) Desired password.
* `ssh_key_path` - (Optional) Path to private ssh key
* `ssh_key_public` - (Optional) The public key data in OpenSSH authorized_keys format.




IPs (`ips`) expose the following attributes

* `id` - (Computed) The ID of the attached IP
* `ip` - (Computed) The IP
* `firewall_policy_id` - (Computed) The attached firewall policy
