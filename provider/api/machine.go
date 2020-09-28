package api

import (
	"log"

	"github.com/erigones/godanube/cloudapi"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func MachineRead(hostname_or_uuid string, d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudapi.Client)

	vmInfo, err := client.GetMachine(hostname_or_uuid)

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
