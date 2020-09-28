package provider

import (
	"log"

	"github.com/erigones/godanube/cloudapi"
	"github.com/fishman/terraform-provider-danube/provider/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMachine() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMachineRead,

		Schema: schemas.MachineSchema(),
	}
}

func dataSourceMachineRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudapi.Client)

	hostname := d.Get("hostname").(string)

	vmInfo, err := client.GetMachine(hostname)

	if err != nil {
		log.Println("[ERR] ", err.Error)
		return err
	}

	d.SetId(vmInfo.Uuid)
	d.Set("hostname", vmInfo.Hostname)
	d.Set("alias", vmInfo.Alias)
	d.Set("owner", vmInfo.Owner)
	d.Set("vcpus", vmInfo.Vcpus)
	d.Set("ram", vmInfo.Ram)
	d.Set("disk", vmInfo.Disk)
	d.Set("locked", vmInfo.Locked)
	d.Set("ips", vmInfo.Ips)
	d.Set("tags", vmInfo.Tags)

	return err
}
