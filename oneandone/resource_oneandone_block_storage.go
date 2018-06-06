package oneandone

import (
	"fmt"
	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

func resourceOneandOneBlockStorage() *schema.Resource {
	return &schema.Resource{
		Create: resourceOneandOneBlockStorageCreate,
		Read:   resourceOneandOneBlockStorageRead,
		Update: resourceOneandOneBlockStorageUpdate,
		Delete: resourceOneandOneBlockStorageDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"datacenter": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceOneandOneBlockStorageCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	req := oneandone.BlockStorageRequest{
		Name: d.Get("name").(string),
		Size: oneandone.Int2Pointer(d.Get("size").(int)),
	}

	if raw, ok := d.GetOk("description"); ok {
		req.Description = raw.(string)
	}

	if raw, ok := d.GetOk("server_id"); ok {
		req.ServerId = raw.(string)
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

	blockStorageId, blockStorage, err := config.API.CreateBlockStorage(&req)
	if err != nil {
		return err
	}

	err = config.API.WaitForState(blockStorage, "POWERED_ON", 10, config.Retries)

	if err != nil {
		return err
	}

	d.SetId(blockStorageId)

	return resourceOneandOneBlockStorageRead(d, meta)
}

func resourceOneandOneBlockStorageRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	blockStorage, err := config.API.GetBlockStorage(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", blockStorage.Name)
	d.Set("description", blockStorage.Description)
	d.Set("size", blockStorage.Size)
	d.Set("datacenter", blockStorage.Datacenter.CountryCode)
	if blockStorage.Server != nil && len(blockStorage.Server.Id) > 0 {
		d.Set("server_id", blockStorage.Server.Id)
	}

	_, serverIdSetInConfig := d.GetOk("server_id")
	if blockStorage.Server == nil && serverIdSetInConfig {
		d.Set("server_id", "")
	}

	return nil
}

func resourceOneandOneBlockStorageUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	blockStorage, err := config.API.GetBlockStorage(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return err
	}

	bsChanged := false
	if d.HasChange("name") || d.HasChange("description") {
		updateBlksReq := oneandone.UpdateBlockStorageRequest{}

		_, new_name := d.GetChange("name")
		if blockStorage.Name != new_name {
			updateBlksReq.Name = new_name.(string)
			bsChanged = true
		}

		_, new_description := d.GetChange("description")
		if blockStorage.Description != new_description {
			updateBlksReq.Description = new_description.(string)
			bsChanged = true
		}

		if bsChanged {
			_, err := config.API.UpdateBlockStorage(blockStorage.Id, &updateBlksReq)
			if err != nil {
				return err
			}
		}
	}

	if d.HasChange("server_id") {
		o, new_server_id := d.GetChange("server_id")

		if (blockStorage.Server != nil && blockStorage.Server.Id != new_server_id) ||
			(blockStorage.Server == nil && len(new_server_id.(string)) > 0) {
			_, err := config.API.AddBlockStorageServer(blockStorage.Id, new_server_id.(string))
			if err != nil {
				return err
			}
		}

		if blockStorage.Server != nil && len(new_server_id.(string)) == 0 {
			_, err := config.API.RemoveBlockStorageServer(blockStorage.Id, o.(string))
			if err != nil {
				return err
			}
		}
	}

	err = config.API.WaitForState(blockStorage, "POWERED_ON", 10, config.Retries)

	if err != nil {
		return err
	}

	return resourceOneandOneBlockStorageRead(d, meta)
}

func resourceOneandOneBlockStorageDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	bs, err := config.API.GetBlockStorage(d.Id())

	if bs.Server != nil {
		config.API.RemoveBlockStorageServer(bs.Id, bs.Server.Id)
		err = config.API.WaitForState(bs, "POWERED_ON", 10, config.Retries)
		if err != nil {
			return err
		}
	}

	_, err = config.API.DeleteBlockStorage(d.Id())
	if err != nil {
		return err
	}

	return nil
}
