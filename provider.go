package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GD_KEY", nil),
				Description: "GoDaddy API Key.",
			},

			"secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GD_SECRET", nil),
				Description: "GoDaddy API Secret.",
			},

			"baseurl": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "https://api.ote-godaddy.com",
				Description: "GoDaddy Base Url(defaults to testing).",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"godaddy_domain_record": resourceDomainRecord(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Key:     d.Get("key").(string),
		Secret:  d.Get("secret").(string),
		BaseURL: d.Get("baseurl").(string),
	}

	return config.Client()
}
