package oneandone

import (
	"fmt"
	"testing"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"time"
)

func TestAccOneandoneImage_Basic(t *testing.T) {
	var image oneandone.Image

	name := "test"
	name_updated := "test1"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDOneandoneImageDestroyCheck,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckOneandoneImage_basic, name),

				Check: resource.ComposeTestCheckFunc(
					func(*terraform.State) error {
						time.Sleep(10 * time.Second)
						return nil
					},
					testAccCheckOneandoneImageExists("oneandone_image.img", &image),
					resource.TestCheckResourceAttr("oneandone_image.img", "name", name),
				),
			},
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckOneandoneImage_update, name_updated),

				Check: resource.ComposeTestCheckFunc(
					func(*terraform.State) error {
						time.Sleep(10 * time.Second)
						return nil
					},
					testAccCheckOneandoneImageExists("oneandone_image.img", &image),
					resource.TestCheckResourceAttr("oneandone_image.img", "name", name_updated),
				),
			},
		},
	})
}

func testAccCheckDOneandoneImageDestroyCheck(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneandone_image.img" {
			continue
		}

		api := oneandone.New(os.Getenv("ONEANDONE_TOKEN"), oneandone.BaseUrl)

		_, err := api.GetImage(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Image still exists %s %s", rs.Primary.ID, err.Error())
		}
	}

	return nil
}

func testAccCheckOneandoneImageExists(n string, img_p *oneandone.Image) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		api := oneandone.New(os.Getenv("ONEANDONE_TOKEN"), oneandone.BaseUrl)

		found_img, err := api.GetImage(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Error occured while fetching Image: %s", rs.Primary.ID)
		}
		if found_img.Id != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}
		img_p = found_img

		return nil
	}
}

const testAccCheckOneandoneImage_basic = `
resource "oneandone_image" "img" {
  name = "%s"
  server_id = "C72CF0A681B0CCE7EC624DD194D585C6",
  description = "Testing terraform 1and1 image create",
  frequency = "WEEKLY",
  num_images = 5
}`

const testAccCheckOneandoneImage_update = `
resource "oneandone_image" "img" {
  name = "%s"
  description = "Testing terraform 1and1 image update",
  frequency = "ONCE",
}`
