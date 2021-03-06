package provider

import (
	"log"
	"sync"

	// "os"

	"github.com/erigones/godanube/cloudapi"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Meta ...
type Meta struct {
	Client *cloudapi.Client
	data   *schema.ResourceData

	// Used to lock some operations
	sync.Mutex
}

// New ...
func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
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
			"danube_image":   dataSourceImage(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"danube_machine": resourceMachine(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	// m := &Meta{data: d}

	log.Println("[DEBUG] Initializing danube provider")
	config := Config{
		APIKey: d.Get("api_key").(string),
		VDC:    d.Get("vdc").(string),
		URL:    d.Get("url").(string),

		// InsecureSkipTLSVerify: d.Get("insecure_skip_tls_verify").(bool),
	}

	client, err := config.Client()
	if err != nil {
		return nil, err
	}

	return client, nil
}
