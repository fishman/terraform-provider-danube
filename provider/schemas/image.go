package schemas

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// "os"

func ImageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}
