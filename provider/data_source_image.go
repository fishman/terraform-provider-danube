package provider

import (
	"log"

	"github.com/erigones/godanube/cloudapi"
	"github.com/fishman/terraform-provider-danube/provider/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceImage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMachineRead,

		Schema: schemas.ImageSchema(),
	}
}

func dataSourceImageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudapi.Client)

	name := d.Get("name").(string)

	info, err := client.GetImage(name)

	if err != nil {
		log.Println("[ERR] ", err.Error)
		return err
	}

	d.SetId(info.Uuid)
	d.Set("name", info.Name)

	return err
}
