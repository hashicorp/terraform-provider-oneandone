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

func TestAccOneandoneSshKey_Basic(t *testing.T) {
	var sshKey oneandone.SSHKey

	name := "test"
	name_updated := "test1"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDOneandoneSshKeyDestroyCheck,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckOneandoneSshKey_basic, name),
				Check: resource.ComposeTestCheckFunc(
					func(*terraform.State) error {
						time.Sleep(10 * time.Second)
						return nil
					},
					testAccCheckOneandoneSshKeyExists("oneandone_ssh_key.ssh_key", &sshKey),
					testAccCheckOneandoneSshKeyAttributes("oneandone_ssh_key.ssh_key", name),
					resource.TestCheckResourceAttr("oneandone_ssh_key.ssh_key", "name", name),
				),
			},
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckOneandoneSshKey_update, name_updated),
				Check: resource.ComposeTestCheckFunc(
					func(*terraform.State) error {
						time.Sleep(10 * time.Second)
						return nil
					},
					testAccCheckOneandoneSshKeyExists("oneandone_ssh_key.ssh_key", &sshKey),
					testAccCheckOneandoneSshKeyAttributes("oneandone_ssh_key.ssh_key", name_updated),
					resource.TestCheckResourceAttr("oneandone_ssh_key.ssh_key", "name", name_updated),
				),
			},
		},
	})
}

func testAccCheckDOneandoneSshKeyDestroyCheck(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneandone_server" {
			continue
		}

		api := oneandone.New(os.Getenv("ONEANDONE_TOKEN"), oneandone.BaseUrl)

		_, err := api.GetSSHKey(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("SSH Key still exists %s %s", rs.Primary.ID, err.Error())
		}
	}

	return nil
}

func testAccCheckOneandoneSshKeyAttributes(n string, reverse_dns string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.Attributes["name"] != reverse_dns {
			return fmt.Errorf("Bad name: expected %s : found %s ", reverse_dns, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func testAccCheckOneandoneSshKeyExists(n string, server *oneandone.SSHKey) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		api := oneandone.New(os.Getenv("ONEANDONE_TOKEN"), oneandone.BaseUrl)

		found_server, err := api.GetSSHKey(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Error occured while fetching SSH Key: %s", rs.Primary.ID)
		}
		if found_server.Id != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}
		server = found_server

		return nil
	}
}

const testAccCheckOneandoneSshKey_basic = `
resource "oneandone_ssh_key" "ssh_key" {
  name = "%s"
  description = "test ssh descr",
}`

const testAccCheckOneandoneSshKey_update = `
resource "oneandone_ssh_key" "ssh_key" {
  name = "%s"
  description = "test ssh descr"
}`
