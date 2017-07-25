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
				"XL", 2, 4),
			testAccDataSourceOneandOneServerStep(
				`vcores = 4
				ram = 8`,
				"XXL", 4, 8),
		},
	})
}

func testAccDataSourceOneandOneServerStep(filter string, name string, vcores int, ram int) resource.TestStep {
	return resource.TestStep{
		Config: testAccDataSourceOneandOneServerSizeConfig(filter),
		Check: resource.ComposeTestCheckFunc(
			testAccDataSourceOneandOneServerSize("data.oneandone_instance_size.serversize", name, vcores, ram),
		),
	}
}

func testAccDataSourceOneandOneServerSize(data_source_name string, sizeName string, vcores int, ram int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[data_source_name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", data_source_name)
		}

		ds_attr := ds.Primary.Attributes

		if len(ds_attr["id"]) < 32 {
			return fmt.Errorf("id is %s(%d) which is too short", ds_attr["id"], len(ds_attr["id"]))
		}
		if ds_attr["name"] != sizeName {
			return fmt.Errorf("name is %s instead of %s", ds_attr["name"], sizeName)
		}
		if ds_attr["vcores"] != fmt.Sprintf("%d", vcores) {
			return fmt.Errorf("vcores of %s is %s instead of %d", sizeName, ds_attr["vcores"], vcores)
		}
		if ds_attr["ram"] != fmt.Sprintf("%d", ram) {
			return fmt.Errorf("ram of %s is %s instead of %d", sizeName, ds_attr["ram"], ram)
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
