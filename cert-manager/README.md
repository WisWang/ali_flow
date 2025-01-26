# Cert-Manager with Aliyun DNS Setup

这个 Helm Chart 用于安装 cert-manager 并配置使用阿里云 DNS 进行 ACME DNS01 验证。

## 前置条件

1. Kubernetes 集群
2. Helm 3
3. 阿里云 AccessKey 和 Secret（需要有 DNS 解析权限）

## 安装

1. 添加 cert-manager 仓库：
```bash
helm repo add jetstack https://charts.jetstack.io
helm repo update
```


配置 alidns-secret

```
AK=$(echo -n <AccessKey ID> | base64)
SK=$(echo -n <AccessKey Secret> | base64)
cat << EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: alidns-secret
  namespace: cert-manager
data:
  access-key: $AK
  secret-key: $SK
EOF
```



2. 安装 Chart：
```bash
helm install cert-manager-setup . \
  --set alicloud.accessKeyId=your-key-id \
  --set alicloud.accessKeySecret=your-key-secret \
  --set certificate.commonName=*.your-domain.com \
  --set "certificate.dnsNames[0]=your-domain.com" \
  --set "certificate.dnsNames[1]=*.your-domain.com"
```

## 配置说明

### 必需值

| 参数 | 描述 |
|------|------|
| `alicloud.accessKeyId` | 阿里云 AccessKey ID |
| `alicloud.accessKeySecret` | 阿里云 AccessKey Secret |

### 可选值

| 参数 | 描述 | 默认值 |
|------|------|--------|
| `alicloud.regionId` | 阿里云区域 | `cn-hangzhou` |
| `certificate.name` | 证书名称 | `example-cert` |
| `certificate.namespace` | 证书命名空间 | `default` |
| `certificate.commonName` | 证书通用名称 | `*.example.com` |
| `certificate.secretName` | 证书密钥名称 | `example-tls` |

## 验证安装

1. 检查 cert-manager 是否正常运行：
```bash
kubectl get pods -n cert-manager
```

2. 检查证书状态：
```bash
kubectl get certificate -n default
kubectl get certificaterequest -n default
kubectl get order -n cert-manager
kubectl get challenge -n cert-manager
```

## 故障排查

1. 查看 cert-manager 日志：
```bash
kubectl logs -n cert-manager -l app=cert-manager
```

2. 检查证书事件：
```bash
kubectl describe certificate example-cert
```

## 注意事项

1. 确保阿里云 AccessKey 有足够的权限
2. 生产环境建议使用 Secrets 管理敏感信息
3. 首次申请证书可能需要几分钟时间
4. Let's Encrypt 生产环境有速率限制

## 卸载

```bash
helm uninstall cert-manager-setup
```

## 常见问题

1. Q: 证书一直处于 pending 状态？
   A: 检查 DNS 权限和 cert-manager 日志

2. Q: 如何更新证书配置？
   A: 修改 values.yaml 后执行 helm upgrade

3. Q: 支持多域名证书吗？
   A: 是的，可以在 certificate.dnsNames 中添加多个域名 