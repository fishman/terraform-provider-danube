package provider

import (
	// "fmt"
	// "github.com/erigones/godanube/cloudapi"

	"log"

	"github.com/erigones/godanube/cloudapi"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMachine() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMachineRead,

		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alias": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ram": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
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

	return err
}
