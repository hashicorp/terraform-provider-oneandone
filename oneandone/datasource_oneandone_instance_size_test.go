package oneandone

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceOneandOneServerSize(t *testing.T) {
	sizeName := "XL"
	sizeId := "F94E3B12D06231D9CDA1859E09133D8A"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceOneandOneServerSizeConfig(sizeName),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceOneandOneServerSize("data.oneandone_instance_size.serversize", sizeName, sizeId),
				),
			},
		},
	})
}

func testAccDataSourceOneandOneServerSize(data_source_name string, sizeName string, sizeId string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[data_source_name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", data_source_name)
		}

		ds_attr := ds.Primary.Attributes

		if ds_attr["name"] != sizeName {
			return fmt.Errorf("name is not %s", sizeName)
		}
		if ds_attr["id"] != sizeId {
			return fmt.Errorf("id of %s is not %s", sizeName, sizeId)
		}
		return nil
	}
}

func testAccDataSourceOneandOneServerSizeConfig(name string) string {
	return fmt.Sprintf(`
data "oneandone_instance_size" "serversize" {
	name = "%s"
}`, name)
}
