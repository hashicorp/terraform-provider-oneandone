package oneandone

import (
	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

func resourceOneandOneSshKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceOneandOneSshKeyCreate,
		Read:   resourceOneandOneSshKeyRead,
		Update: resourceOneandOneSshKeyUpdate,
		Delete: resourceOneandOneSshKeyDelete,
		Schema: map[string]*schema.Schema{

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"servers": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Computed: true,
			},
			"md5": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceOneandOneSshKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	req := oneandone.SSHKeyRequest{
		Name: d.Get("name").(string),
	}

	if raw, ok := d.GetOk("description"); ok {
		req.Description = raw.(string)
	}

	if raw, ok := d.GetOk("public_key"); ok {
		req.PublicKey = raw.(string)
	}

	sshKey_id, _, err := config.API.CreateSSHKey(&req)
	if err != nil {
		return err
	}

	d.SetId(sshKey_id)

	return resourceOneandOneSshKeyRead(d, meta)
}

func resourceOneandOneSshKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	_, err := config.API.RenameSSHKey(d.Id(), d.Get("name").(string), d.Get("description").(string))
	if err != nil {
		return err
	}

	return resourceOneandOneSshKeyRead(d, meta)
}

func resourceOneandOneSshKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	sshKey, err := config.API.GetSSHKey(d.Id())

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", sshKey.Name)
	d.Set("description", sshKey.Description)

	pkey := sshKey.PublicKey
	if !strings.HasPrefix(pkey, "ssh-rsa ") {
		pkey = "ssh-rsa " + pkey
	}

	d.Set("public_key", pkey)
	d.Set("md5", sshKey.Md5)
	d.Set("servers", getSshServers(*sshKey.Servers))

	return nil
}

func getSshServers(servers []oneandone.SSHServer) []map[string]interface{} {
	raw := make([]map[string]interface{}, 0, len(servers))

	for _, server := range servers {

		toadd := map[string]interface{}{
			"id":   server.Id,
			"name": server.Name,
		}

		raw = append(raw, toadd)
	}

	return raw
}

func resourceOneandOneSshKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	_, err := config.API.DeleteSSHKey(d.Id())
	if err != nil {
		return err
	}

	return nil
}
