## 1.1.1 (Unreleased)

FEATURES:

* **New Data Source:** `oneandone_instance_size` (#4)[https://github.com/terraform-providers/terraform-provider-oneandone/issues/4]

IMPROVEMENTS: 

* `oneandone_server` - added `fixed_instance_size` parameter [GH-5]
* `oneandone_server` - added `ssh_key_public` parameter [GH-6]

BUG FIXES:

* resource/resource_oneandone_server.go: Added missing update hardware function [GH-2]
* resource/resource_oneandone_server.go: Added `ForceNew` on image parameter of server resource [GH-8]



## 0.1.0 (June 21, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
