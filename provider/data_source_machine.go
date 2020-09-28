package provider

import (
	"github.com/fishman/terraform-provider-danube/provider/api"
	"github.com/fishman/terraform-provider-danube/provider/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMachine() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMachineRead,

		Schema: schemas.MachineDataSchema(),
	}
}

func dataSourceMachineRead(d *schema.ResourceData, meta interface{}) error {
	hostname := d.Get("hostname").(string)

	return api.MachineRead(hostname, d, meta)
}
