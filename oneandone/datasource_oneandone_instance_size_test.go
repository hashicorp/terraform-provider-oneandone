package oneandone

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceOneandOneServerSize(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			testAccDataSourceOneandOneServerStep(
				`name = "XL"`,
				"XL", "F94E3B12D06231D9CDA1859E09133D8A"),
			testAccDataSourceOneandOneServerStep(
				`vcores = 4
				ram = 8`,
				"XXL", "2E4092AAB9D31CBA368B05AC70ABEC5A"),
		},
	})
}

func testAccDataSourceOneandOneServerStep(filter string, name string, id string) resource.TestStep {
	return resource.TestStep{
		Config: testAccDataSourceOneandOneServerSizeConfig(filter),
		Check: resource.ComposeTestCheckFunc(
			testAccDataSourceOneandOneServerSize("data.oneandone_instance_size.serversize", name, id),
		),
	}
}

func testAccDataSourceOneandOneServerSize(data_source_name string, sizeName string, sizeId string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[data_source_name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", data_source_name)
		}

		ds_attr := ds.Primary.Attributes

		if ds_attr["name"] != sizeName {
			return fmt.Errorf("name is %s instead of %s", ds_attr["name"], sizeName)
		}
		if ds_attr["id"] != sizeId {
			return fmt.Errorf("id of %s is %s instead of %s", sizeName, ds_attr["id"], sizeId)
		}
		return nil
	}
}

func testAccDataSourceOneandOneServerSizeConfig(name string) string {
	return fmt.Sprintf(`
data "oneandone_instance_size" "serversize" {
	%s
}`, name)
}
