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

func TestAccOneandoneBlockStorage_Basic(t *testing.T) {
	var server oneandone.Server
	image := "centos6-64min"

	var storage oneandone.BlockStorage
	name := "test_storage1"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDOneandoneBlockStorageDestroyCheck,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckOneandoneBlockStorage_basic, name, name, image, name),
				Check: resource.ComposeTestCheckFunc(
					func(*terraform.State) error {
						time.Sleep(10 * time.Second)
						return nil
					},
					testAccCheckOneandoneServerExists("oneandone_server.server", &server),
					testAccCheckOneandoneBlockStorageExists("oneandone_block_storage.storage", &storage),
					resource.TestCheckResourceAttr("oneandone_block_storage.storage", "name", name),
				),
			},
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckOneandoneBlockStorage_update, name, name, image, name),
				Check: resource.ComposeTestCheckFunc(
					func(*terraform.State) error {
						time.Sleep(10 * time.Second)
						return nil
					},
					testAccCheckOneandoneServerExists("oneandone_server.server", &server),
					testAccCheckOneandoneBlockStorageExists("oneandone_block_storage.storage", &storage),
					resource.TestCheckResourceAttr("oneandone_block_storage.storage", "name", name),
					resource.TestCheckResourceAttr("oneandone_block_storage.storage", "server_id", ""),
				),
			},
		},
	})
}

func testAccCheckDOneandoneBlockStorageDestroyCheck(s *terraform.State) error {
	for _, blockStorage := range s.RootModule().Resources {
		if blockStorage.Type != "oneandone_block_storage" {
			continue
		}

		api := oneandone.New(os.Getenv("ONEANDONE_TOKEN"), oneandone.BaseUrl)

		_, err := api.GetBlockStorage(blockStorage.Primary.ID)

		if err == nil {
			return fmt.Errorf("Block storage still exists %s %s", blockStorage.Primary.ID, err.Error())
		}
	}

	return nil
}

func testAccCheckOneandoneBlockStorageExists(n string, storage *oneandone.BlockStorage) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		api := oneandone.New(os.Getenv("ONEANDONE_TOKEN"), oneandone.BaseUrl)

		found_storage, err := api.GetBlockStorage(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Error occured while fetching BlockStorage: %s", rs.Primary.ID)
		}
		if found_storage.Id != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}
		storage = found_storage

		return nil
	}
}

const testAccCheckOneandoneBlockStorage_basic = `
resource "oneandone_server" "server" {
  name = "%s"
  description = "%s"
  image = "%s"
  datacenter = "GB"
  vcores = 1
  cores_per_processor = 1
  ram = 2
  password = "Kv40kd8PQb"
  hdds = [
    {
      disk_size = 30
      is_main = true
    }
  ]
}

resource "oneandone_block_storage" "storage" {
	name = "%s"
	description = "ttt"
	size = 20
	datacenter = "GB"
	server_id = "${oneandone_server.server.id}"
}`

const testAccCheckOneandoneBlockStorage_update = `
resource "oneandone_server" "server" {
  name = "%s"
  description = "%s"
  image = "%s"
  datacenter = "GB"
  vcores = 1
  cores_per_processor = 1
  ram = 2
  password = "Kv40kd8PQb"
  hdds = [
    {
      disk_size = 30
      is_main = true
    }
  ]
}

resource "oneandone_block_storage" "storage" {
	name = "%s"
	description = "ttt"
	size = 20
	datacenter = "GB"
}`
