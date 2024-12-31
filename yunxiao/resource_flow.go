package yunxiao

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFlowCreate,
		ReadContext:   resourceFlowRead,
		UpdateContext: resourceFlowUpdate,
		DeleteContext: resourceFlowDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "流水线名称",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "流水线描述",
			},
			"config": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "流水线配置",
			},
		},
	}
}

func resourceFlowCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	configRaw := d.Get("config").(map[string]interface{})

	flow := &Flow{
		Name:        name,
		Description: description,
		Config: map[string]interface{}{
			"stages": configRaw["stages"],
		},
	}

	err := client.CreateFlow(flow)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(flow.Name)

	return resourceFlowRead(ctx, d, m)
}

func resourceFlowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// 实现读取逻辑
	return nil
}

func resourceFlowUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	if d.HasChange("name") || d.HasChange("description") || d.HasChange("config") {
		flow := &Flow{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Config:      d.Get("config").(map[string]interface{}),
		}

		err := client.UpdateFlow(d.Id(), flow)
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error updating Flow: %s", err))
		}
	}

	return resourceFlowRead(ctx, d, m)
}

func resourceFlowDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// 实现删除逻辑
	return nil
} 