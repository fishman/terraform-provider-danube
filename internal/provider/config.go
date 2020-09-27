package provider

import (
	"os"

	"log"

	"github.com/erigones/godanube/auth"
	"github.com/erigones/godanube/client"
	"github.com/erigones/godanube/cloudapi"
	// dcerrors "github.com/erigones/godanube/errors"
	// "github.com/hashicorp/errwrap"
)

// Config  ...
type Config struct {
	Debug                 bool
	User                  string
	APIKey                string
	VDC                   string
	URL                   string
	InsecureSkipTLSVerify bool
}

// Client  ...
func (c *Config) Client() (*cloudapi.Client, error) {
	// ctx := context.Background()

	userAuth, err := auth.NewAuth("", "", c.APIKey)
	if err != nil {
		return nil, err
	}

	creds := &auth.Credentials{
		UserAuthentication: userAuth,
		ApiEndpoint:        auth.Endpoint{URL: c.URL},
		VirtDatacenter:     c.VDC,
	}

	return cloudapi.New(client.NewClient(
		creds,
		cloudapi.DefaultAPIVersion,
		log.New(os.Stderr, "", log.LstdFlags),
		// newGoLogger(),
	)), nil
}
