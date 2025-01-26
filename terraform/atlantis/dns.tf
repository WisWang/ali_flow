



resource "alicloud_alidns_record" "m22_dns_record" {
  domain_name = "aliyun.18600113834.shop" # 替换为你的顶级域名
  rr          = "m22"                      # 子域名部分
  type        = "A"                        # DNS记录类型，这里以A记录为例
  value       = "10.8.1.2"               # A记录对应的IP地址，根据实际需求替换
  ttl         = 600                        # TTL值，可选，默认是600秒
}

