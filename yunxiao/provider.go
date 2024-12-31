package yunxiao

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("YUNXIAO_ACCESS_KEY", nil),
				Description: "阿里云访问密钥ID (Access Key ID). 可以通过环境变量 YUNXIAO_ACCESS_KEY 设置",
			},
			"access_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("YUNXIAO_ACCESS_SECRET", nil),
				Description: "阿里云访问密钥密码 (Access Key Secret). 可以通过环境变量 YUNXIAO_ACCESS_SECRET 设置",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "https://devops.cn-hangzhou.aliyuncs.com",
				Description: "云效API接入点（默认为杭州区域）",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"yunxiao_flow": resourceFlow(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	accessKey := d.Get("access_key").(string)
	accessSecret := d.Get("access_secret").(string)
	endpoint := d.Get("endpoint").(string)

	// 创建客户端并处理错误
	client, err := NewClient(accessKey, accessSecret, endpoint)
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("Error creating Yunxiao client: %s", err))
	}

	return client, nil
} 