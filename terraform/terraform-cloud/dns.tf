
resource "alicloud_dns_domain" "example_domain" {
  domain_name = "aliyun.18600113834.shop"
}

resource "alicloud_dns_record" "example_a_record" {
  domain_name = alicloud_dns_domain.example_domain.name
  type        = "A"
  host_record        = "mfa"
  value       = "192.168.1.100"
  ttl         = 600
}