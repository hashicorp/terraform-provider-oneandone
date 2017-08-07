---
layout: "oneandone"
page_title: "1&1: oneandone_server"
sidebar_current: "docs-oneandone-resource-server"
description: |-
  Creates and manages 1&1 Server.
---

# oneandone\_server

Manages a Server on 1&1

## Example Usage

```hcl
resource "oneandone_server" "server" {
  name = "Example"
  description = "Terraform 1and1 tutorial"
  image = "ubuntu"
  datacenter = "GB"
  vcores = 1
  cores_per_processor = 1
  ram = 2
  ssh_key_path = "/path/to/private/ssh_key"
  ssh_key_public = "${file("/path/to/public/key.pub")}"
  hdds = [
    {
      disk_size = 60
      is_main = true
    }
  ]

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

* `cores_per_processor` -(Optional) Number of cores per processor
* `datacenter` - (Optional) Location of desired 1and1 datacenter. Can be `DE`, `GB`, `US` or `ES`
* `description` - (Optional) Description of the server
* `firewall_policy_id` - (Optional) ID of firewall policy
* `fixed_instance_size` - (Optional) ID of a fixed instance size
* `hdds` - (Optional) List of HDDs. One HDD must be main.
* `*disk_size` -(Required) The size of HDD
* `*is_main` - (Optional) Indicates if HDD is to be used as main hard disk of the server
* `image` -(Required) The name of a desired image to be provisioned with the server
* `ip` - (Optional) IP address for the server
* `loadbalancer_id` - (Optional) ID of the load balancer
* `monitoring_policy_id` - (Optional) ID of monitoring policy
* `name` -(Required) The name of the server.
* `password` - (Optional) Desired password.
* `ram` -(Optional) Size of ram.
* `ssh_key_path` - (Optional) Path to private ssh key
* `ssh_key_public` - (Optional) The public key data in OpenSSH authorized_keys format.
* `vcores` -(Optional) Number of virtual cores.

Either `fixed_instance_size` or all of `vcores`, `cores_per_processor`, `ram` and `hdds` are required.

IPs (`ips`) expose the following attributes

* `id` - (Computed) The ID of the attached IP
* `ip` - (Computed) The IP 
* `firewall_policy_id` - (Computed) The attached firewall policy
