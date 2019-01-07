package oneandone

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccOneandoneBaremetalServer_Hardware(t *testing.T) {
	_, exists := os.LookupEnv("TF_ACC_BAREMETAL_LOCK")
	if !exists {
		t.Skip("`TF_ACC_BAREMETAL_LOCK` isn't specified - skipping since test will increase test time significantly")
	}

	var server oneandone.Server

	name := "test_server_hardware"
	image := "CENTOS 6 MINIMAL SYSTEM (64BIT)"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDOneandoneBaremetalServerDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckOneandoneBaremetalServer_basic, name, name, image),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOneandoneBaremetalServerExists("oneandone_baremetal.server", &server),
					resource.TestCheckResourceAttr("oneandone_baremetal.server", "name", name),
				),
			},
			{
				Config: fmt.Sprintf(testAccCheckOneandoneBaremetalServer_hardware, name, name, image),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOneandoneBaremetalServerExists("oneandone_baremetal.server", &server),
					resource.TestCheckResourceAttr("oneandone_baremetal.server", "name", name),
				),
			},
		},
	})
}

func testAccCheckDOneandoneBaremetalServerDestroyCheck(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneandone_server" {
			continue
		}

		api := oneandone.New(os.Getenv("ONEANDONE_TOKEN"), oneandone.BaseUrl)

		_, err := api.GetServer(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Server still exists %s %s", rs.Primary.ID, err.Error())
		}
	}

	return nil
}

func testAccCheckOneandoneBaremetalServerExists(n string, server *oneandone.Server) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		api := oneandone.New(os.Getenv("ONEANDONE_TOKEN"), oneandone.BaseUrl)

		updating := true
		for updating {
			found_server, err := api.GetServer(rs.Primary.ID)

			if err != nil {
				return fmt.Errorf("Error occured while fetching Server: %s", rs.Primary.ID)
			}
			if found_server.Id != rs.Primary.ID {
				return fmt.Errorf("Record not found")
			}
			if found_server.Status.Percent == 0 {
				updating = false
				server = found_server
			} else {
				log.Print("waiting:", found_server.Status.Percent)
				time.Sleep(time.Second)
			}
		}
		return nil
	}
}

const testAccCheckOneandoneBaremetalServer_basic = `
resource "oneandone_baremetal" "server" {
  name = "%s"
  description = "%s"
  image = "%s"
  password = "Kv40kd8PQb"
  datacenter = "US"
  baremetal_model_id = "B77E19E062D5818532EFF11C747BD104"
}`

const testAccCheckOneandoneBaremetalServer_hardware = `
resource "oneandone_baremetal" "server" {
  name = "%s"
  description = "%s"
  datacenter = "US"
  image = "%s"
  baremetal_model_id = "B77E19E062D5818532EFF11C747BD104"
  password = "Kv40kd8PQb"  
}`
