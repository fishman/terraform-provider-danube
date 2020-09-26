package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DANUBE_USER", "admin"),
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DANUBE_APIKEY", nil),
			},
			"vdc": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DANUBE_VDC", "main"),
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DANUBE_URL", "https://localhost:46483/api/"),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"danube_machine": dataSourceMachine(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"danube_machine": resourceMachine(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		User:   d.Get("user").(string),
		APIKey: d.Get("api_key").(string),
		VDC:    d.Get("vdc").(string),
		URL:    d.Get("url").(string),

		InsecureSkipTLSVerify: d.Get("insecure_skip_tls_verify").(bool),
	}

	if user, ok := d.GetOk("user"); ok {
		config.User = user.(string)
	}

	client, err := config.Client()
	if err != nil {
		return nil, err
	}

	return client, nil
}
