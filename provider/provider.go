package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/justsimplify/redis-client/modules/redis"
	"github.com/justsimplify/redis-terraform/resources/crud_resources"
	"github.com/justsimplify/redis-terraform/resources/data_sources"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("REDIS_HOST", "0.0.0.0"),
			},
			"port": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("REDIS_PORT", "6379"),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("REDIS_PASSWORD", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"redis-config": crud_resources.RedisCRUDResource(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"redis-config": data_sources.RedisResource(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	host := d.Get("host").(string)
	port := d.Get("port").(string)
	password := d.Get("password").(string)

	rc := redis.Client{
		Host:     host,
		Port:     port,
		Password: password,
	}

	return rc, diags
}

