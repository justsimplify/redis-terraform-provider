package crud_resources

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/justsimplify/redis-client/modules"
	"github.com/justsimplify/redis-client/modules/redis"
	"strconv"
	"time"
)

func RedisCRUDResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type: schema.TypeString,
				Required: true,
			},
			"value": {
				Type: schema.TypeString,
				Required: true,
			},
		},
		ReadContext: redisRead,
		UpdateContext: redisUpdate,
		DeleteContext: redisDelete,
		CreateContext: redisCreate,
	}
}

func redisRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	rc := m.(redis.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	res, err := rc.Read(ctx, d.Get("key").(string))
	if err != nil {
		d.SetId("")
		return diags
	}

	if err := d.Set("value", res); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func redisUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	rc := m.(redis.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := rc.Update(ctx, modules.RedisPayload{
		Key:   d.Get("key").(string),
		Value: d.Get("value").(string),
	})

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func redisDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	rc := m.(redis.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := rc.Delete(ctx, d.Get("key").(string))

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func redisCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	rc := m.(redis.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := rc.Create(ctx, modules.RedisPayload{
		Key:   d.Get("key").(string),
		Value: d.Get("value").(string),
	})

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
