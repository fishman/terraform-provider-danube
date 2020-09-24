package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceMachineCreate,
		Read:   resourceMachineRead,
		Update: resourceMachineUpdate,
		Delete: resourceMachineDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceMachineCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceMachineRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceMachineUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceMachineDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
