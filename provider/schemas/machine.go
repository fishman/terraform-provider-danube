package schemas

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func MachineSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"disk": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"locked": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"node_status": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ips": {
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		"tags": {
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
	}
}
