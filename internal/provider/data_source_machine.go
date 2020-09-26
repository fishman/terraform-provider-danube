package provider

import (
	"fmt"
	"github.com/erigones/godanube/cloudapi"
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
	client := meta.(*cloudapi.Client)

	vmInfo, err := client.GetMachine("test")
	if err != nil {
		fmt.Println("error:" + err.Error())
	} else {
		fmt.Println(vmInfo)
	}
	return nil
}
