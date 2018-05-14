---
layout: "oneandone"
page_title: "1&1: oneandone_ssh_key"
sidebar_current: "docs-oneandone-resource-ssh_key"
description: |-
  Creates and manages 1&1 SSH Key.
---

# oneandone\_ssh\_key

Manages SSH Keys on 1&1

## Example Usage

```hcl
resource "oneandone_ssh_key" "sshkey" {
  name = "test_ssh_key"
  description = "testing_ssh_keys"
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) Description for the ssh key
* `name` - (Required) The name of the storage
* `public_key` - (Optional) Public key to import. If not given, new SSH key pair will be created and the private key is returned in the response
