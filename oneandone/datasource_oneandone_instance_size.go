package oneandone

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceOneandOneServerSize() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOneandOneFixedInstanceSizesRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"vcores", "ram"},
			},
			"vcores": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ram": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func dataSourceOneandOneFixedInstanceSizesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	fixed_instance_sizes, err := config.API.ListFixedInstanceSizes()
	if err != nil {
		return err
	}

	name := d.Get("name").(string)
	vcores := d.Get("vcores").(int)
	ram := d.Get("ram").(float64)

	if name == "" && vcores == 0 && ram == 0 {
		return errors.New("no FixedInstanceSizes selectors set")
	}

	for _, size := range fixed_instance_sizes {
		if name != "" && size.Name != name {
			continue
		}
		if vcores != 0 && size.Hardware.Vcores != vcores {
			continue
		}
		if ram != 0 && size.Hardware.Ram != float32(ram) {
			continue
		}

		d.SetId(size.Id)
		d.Set("name", size.Name)
		d.Set("vcores", size.Hardware.Vcores)
		d.Set("ram", size.Hardware.Ram)
		d.Set("coresPerProcessor", size.Hardware.CoresPerProcessor)
		return nil
	}
	return errors.New("size not found")
}
