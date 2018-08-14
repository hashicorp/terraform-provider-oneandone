package oneandone

import (
	"fmt"
	"log"
	"strings"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOneandOneBaremetalServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceOneandOneBaremetalServerCreate,
		Read:   resourceOneandOneServerRead,
		Update: resourceOneandOneBaremetalServerUpdate,
		Delete: resourceOneandOneServerDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"image": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"baremetal_model_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ssh_key_path": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ssh_key_public": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"ssh_key_path"},
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"datacenter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"firewall_policy_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				Computed: true,
			},
			"firewall_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"monitoring_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"loadbalancer_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceOneandOneBaremetalServerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	saps, _ := config.API.ListServerAppliances(0, 0, "", "baremetal", "")
	var sa oneandone.ServerAppliance
	imageName := d.Get("image").(string)

	for _, a := range saps {

		if a.Type == "BAREMETAL" && strings.ToLower(a.Name) == strings.ToLower(imageName) {
			sa = a
			break
		}
	}
	//if the exact name is not found we will try to match
	if sa.Name != imageName {
		for _, a := range saps {

			if a.Type == "BAREMETAL" && strings.Contains(strings.ToLower(a.Name), strings.ToLower(d.Get("image").(string))) {
				sa = a
				break
			}
		}
	}

	req := oneandone.ServerRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		ApplianceId: sa.Id,
		PowerOn:     true,
		ServerType:  "baremetal",
	}

	if baremetal_model_id := d.Get("baremetal_model_id").(string); baremetal_model_id != "" {
		req.Hardware = oneandone.Hardware{
			BaremetalModelId: baremetal_model_id,
		}
	} else {
		return fmt.Errorf(fmt.Sprintf("must provide a valid baremetal model id %d", baremetal_model_id))
	}

	if raw, ok := d.GetOk("ip"); ok {

		new_ip := raw.(string)

		ips, err := config.API.ListPublicIps()
		if err != nil {
			return err
		}

		for _, ip := range ips {
			if ip.IpAddress == new_ip {
				req.IpId = ip.Id
				break
			}
		}

		log.Println("[DEBUG] req.IP", req.IpId)
	}

	if raw, ok := d.GetOk("datacenter"); ok {

		dcs, err := config.API.ListDatacenters()

		if err != nil {
			return fmt.Errorf("An error occured while fetching list of datacenters %s", err)

		}

		decenter := raw.(string)
		for _, dc := range dcs {
			if strings.ToLower(dc.CountryCode) == strings.ToLower(decenter) {
				req.DatacenterId = dc.Id
				break
			}
		}
	}

	if fwp_id, ok := d.GetOk("firewall_policy_id"); ok {
		req.FirewallPolicyId = fwp_id.(string)
	}

	if mp_id, ok := d.GetOk("monitoring_policy_id"); ok {
		req.MonitoringPolicyId = mp_id.(string)
	}

	if lb_id, ok := d.GetOk("loadbalancer_id"); ok {
		req.LoadBalancerId = lb_id.(string)
	}

	var privateKey string
	if raw, ok := d.GetOk("ssh_key_path"); ok {
		rawpath := raw.(string)

		priv, publicKey, err := getSshKey(rawpath)
		privateKey = priv
		if err != nil {
			return err
		}

		req.SSHKey = publicKey
	}

	if raw, ok := d.GetOk("ssh_key_public"); ok {
		req.SSHKey = raw.(string)
	}

	var password string
	if raw, ok := d.GetOk("password"); ok {
		req.Password = raw.(string)
		password = req.Password
	}

	server_id, server, err := config.API.CreateServer(&req)
	if err != nil {
		return err
	}

	err = config.API.WaitForState(server, "POWERED_ON", 10, config.Retries)

	d.SetId(server_id)
	server, err = config.API.GetServer(d.Id())
	if err != nil {
		return err
	}

	if password == "" {
		password = server.FirstPassword
	}
	d.SetConnInfo(map[string]string{
		"type":        "ssh",
		"host":        server.Ips[0].Ip,
		"password":    password,
		"private_key": privateKey,
	})

	return resourceOneandOneServerRead(d, meta)
}

func resourceOneandOneBaremetalServerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	var baremetal_model_id string
	if tmp := d.Get("baremetal_model_id").(string); tmp != "" {
		baremetal_model_id = tmp
	}

	if d.HasChange("name") || d.HasChange("description") {
		_, name := d.GetChange("name")
		_, description := d.GetChange("description")
		server, err := config.API.RenameServer(d.Id(), name.(string), description.(string))
		if err != nil {
			return err
		}

		err = config.API.WaitForState(server, "POWERED_ON", 10, config.Retries)

	}

	if d.HasChange("monitoring_policy_id") {
		o, n := d.GetChange("monitoring_policy_id")

		if n == nil {
			mp, err := config.API.RemoveMonitoringPolicyServer(o.(string), d.Id())

			if err != nil {
				return err
			}

			err = config.API.WaitForState(mp, "ACTIVE", 30, config.Retries)
			if err != nil {
				return err
			}
		} else {
			mp, err := config.API.AttachMonitoringPolicyServers(n.(string), []string{d.Id()})
			if err != nil {
				return err
			}

			err = config.API.WaitForState(mp, "ACTIVE", 30, config.Retries)
			if err != nil {
				return err
			}
		}
	}

	if d.HasChange("loadbalancer_id") {
		o, n := d.GetChange("loadbalancer_id")
		server, err := config.API.GetServer(d.Id())
		if err != nil {
			return err
		}

		if n == nil || n.(string) == "" {
			log.Println("[DEBUG] Removing")
			log.Println("[DEBUG] IPS:", server.Ips)

			for _, ip := range server.Ips {
				mp, err := config.API.DeleteLoadBalancerServerIp(o.(string), ip.Id)

				if err != nil {
					return err
				}

				err = config.API.WaitForState(mp, "ACTIVE", 30, config.Retries)
				if err != nil {
					return err
				}
			}
		} else {
			log.Println("[DEBUG] Adding")
			ip_ids := []string{}
			for _, ip := range server.Ips {
				ip_ids = append(ip_ids, ip.Id)
			}
			mp, err := config.API.AddLoadBalancerServerIps(n.(string), ip_ids)
			if err != nil {
				return err
			}

			err = config.API.WaitForState(mp, "ACTIVE", 30, config.Retries)
			if err != nil {
				return err
			}

		}
	}

	if d.HasChange("firewall_policy_id") {
		server, err := config.API.GetServer(d.Id())
		if err != nil {
			return err
		}

		_, n := d.GetChange("firewall_policy_id")
		if n != nil {
			ip_ids := []string{}
			for _, ip := range server.Ips {
				ip_ids = append(ip_ids, ip.Id)
			}

			mp, err := config.API.AddFirewallPolicyServerIps(n.(string), ip_ids)
			if err != nil {
				return err
			}

			err = config.API.WaitForState(mp, "ACTIVE", 30, config.Retries)
			if err != nil {
				return err
			}
		}
	}

	var baremetalModelId string
	if d.HasChange("baremetal_model_id") {
		baremetalModelId = baremetal_model_id
	}

	hw := &oneandone.Hardware{}

	if baremetalModelId != "" {
		hw.BaremetalModelId = baremetalModelId
	}
	if hw != nil && hw.FixedInsSizeId != "" {
		srv, err := config.API.UpdateServerHardware(d.Id(), hw)
		if err != nil {
			return err
		}
		err = config.API.WaitForState(srv, "POWERED_ON", 30, config.Retries)
	}

	return resourceOneandOneServerRead(d, meta)
}
