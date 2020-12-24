package petstore

import (
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdk "github.com/terraform-in-action/go-petstore"
)

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PETSTORE_ADDRESS", nil),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"petstore_pet_ids": dataSourcePSPetIDs(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"petstore_pet": resourcePSPet(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	hostname, _ := d.Get("address").(string)
	address, _ := url.Parse(hostname)
	cfg := &sdk.Config{
		Address: address.String(),
	}
	return sdk.NewClient(cfg)
}
