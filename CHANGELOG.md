## 1.0.0 (December 18, 2017)

FEATURES:

* **New Data Source:** `oneandone_instance_size` (#4)[https://github.com/terraform-providers/terraform-provider-oneandone/issues/4]

IMPROVEMENTS: 

* `oneandone_server` - added `fixed_instance_size` parameter ([#5](https://github.com/terraform-providers/terraform-provider-oneandone/issues/5))
* `oneandone_server` - added `ssh_key_public` parameter ([#6](https://github.com/terraform-providers/terraform-provider-oneandone/issues/6))

BUG FIXES:

* resource/resource_oneandone_server.go: Added missing update hardware function ([#2](https://github.com/terraform-providers/terraform-provider-oneandone/issues/2))
* resource/resource_oneandone_server.go: Added `ForceNew` on image parameter of server resource ([#8](https://github.com/terraform-providers/terraform-provider-oneandone/issues/8))



## 0.1.0 (June 21, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
