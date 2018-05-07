package oneandone

import (
	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"strings"
)

func resourceOneandOneImage() *schema.Resource {
	return &schema.Resource{

		Create: resourceOneandOneImageCreate,
		Read:   resourceOneandOneImageRead,
		Update: resourceOneandOneImageUpdate,
		Delete: resourceOneandOneImageDelete,
		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"frequency": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"datacenter_id"},
			},
			"num_images": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1, 50),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"datacenter_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"frequency"},
			},
			"source": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceOneandOneImageCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	req := oneandone.ImageRequest{
		Name:     d.Get("name").(string),
	}

	if serverId, ok := d.GetOk("server_id"); ok {
		req.ServerId = serverId.(string)
	}

	if frequency, ok := d.GetOk("frequency"); ok {
		req.Frequency = frequency.(string)
	}

	if numImages, ok := d.GetOk("num_images"); ok {
		req.NumImages = oneandone.Int2Pointer(numImages.(int))
	}

	if desc, ok := d.GetOk("description"); ok {
		req.Description = desc.(string)
	}

	if datacenterId, ok := d.GetOk("datacenter_id"); ok {
		req.DatacenterId = datacenterId.(string)
	}

	if source, ok := d.GetOk("source"); ok {
		req.Source = source.(string)
	}

	if url, ok := d.GetOk("url"); ok {
		req.Url = url.(string)
	}

	if osId, ok := d.GetOk("os_id"); ok {
		req.OsId = osId.(string)
	}

	if typ, ok := d.GetOk("type"); ok {
		req.Type = typ.(string)
	}

	imgId, img, err := config.API.CreateImage(&req)
	if err != nil {
		return err
	}

	err = config.API.WaitForState(img, "ENABLED", 10, config.Retries)
	if err != nil {
		return err
	}

	d.SetId(imgId)

	if err != nil {
		return err
	}

	return resourceOneandOneImageRead(d, meta)
}

func resourceOneandOneImageUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	if d.HasChange("name") || d.HasChange("description") || d.HasChange("frequency") {
		img, err := config.API.UpdateImage(d.Id(), d.Get("name").(string),
			d.Get("description").(string), d.Get("frequency").(string))
		if err != nil {
			return err
		}
		err = config.API.WaitForState(img, "ENABLED", 10, config.Retries)
		if err != nil {
			return err
		}
	}

	return resourceOneandOneImageRead(d, meta)
}

func resourceOneandOneImageRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	img, err := config.API.GetImage(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", img.Name)
	d.Set("datacenter", img.Datacenter.CountryCode)
	d.Set("os_family", img.OsFamily)
	d.Set("os", img.Os)
	d.Set("os_version", img.OsVersion)
	d.Set("architecture", img.Architecture)
	d.Set("os_image_type", img.OsImageType)
	d.Set("type", img.Type)
	d.Set("min_hdd_size", img.MinHddSize)
	d.Set("licenses", img.Licenses)
	d.Set("state", img.State)
	d.Set("description", img.Description)
	d.Set("hdds", readImageHdds(img))
	d.Set("server_id", img.ServerId)
	d.Set("frequency", img.Frequency)
	d.Set("num_images", img.NumImages)
	d.Set("creation_date", img.CreationDate)

	return nil

}

func resourceOneandOneImageDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	img, err := config.API.DeleteImage(d.Id())
	if err != nil {
		return err
	}

	err = config.API.WaitUntilDeleted(img)
	if err != nil {
		return err
	}

	return nil
}

func readImageHdds(image *oneandone.Image) []map[string]interface{} {
	hdds := make([]map[string]interface{}, 0, len(image.Hdds))

	for _, hd := range image.Hdds {
		hdds = append(hdds, map[string]interface{}{
			"id":      hd.Id,
			"size":    hd.Size,
			"is_main": hd.IsMain,
		})
	}

	return hdds
}
