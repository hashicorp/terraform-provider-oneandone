## 1.2.1 (Unreleased)

IMPROVEMENTS:

* Firewall additions and improvements  [#20](https://github.com/terraform-providers/terraform-provider-oneandone/pull/20)

## 1.2.0 (July 24, 2018)

IMPROVEMENTS: 


* Updated 1&1 Go SDK dependency [#19](https://github.com/terraform-providers/terraform-provider-oneandone/pull/19)
* Updated block storage update method and test [#19](https://github.com/terraform-providers/terraform-provider-oneandone/pull/19)
* Added email and description properties to monitoring policy create method [#19](https://github.com/terraform-providers/terraform-provider-oneandone/pull/19)
* Updated monitoring policy test [#19](https://github.com/terraform-providers/terraform-provider-oneandone/pull/19)
* WaitUntilDeleted improvements [#19](https://github.com/terraform-providers/terraform-provider-oneandone/pull/19)

BUG FIXES:

* Fixed issue with server attach [#18](https://github.com/terraform-providers/terraform-provider-oneandone/issues/18) 

## 1.1.0 (May 17, 2018)

FEATURES:
* **New Resource:** `oneandone_image` [#16](https://github.com/terraform-providers/terraform-provider-oneandone/pull/16)
* **New Resource:** `oneandone_block_storage` [#14](https://github.com/terraform-providers/terraform-provider-oneandone/pull/14)
* **New Resource:** `oneandone_ssh_key` [#14](https://github.com/terraform-providers/terraform-provider-oneandone/pull/14)

IMPROVEMENTS: 
* README update (#13)[https://github.com/terraform-providers/terraform-provider-oneandone/pull/13]

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
