package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMachine() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMachineRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceMachineRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
