# 简单Shell命令Flow示例

这个示例展示了如何使用云效Flow Terraform Provider创建一个最简单的流水线，该流水线只执行一个shell命令。

## 前置条件

1. 已安装Terraform (>= 0.12)
2. 已获取阿里云访问密钥（Access Key ID 和 Access Key Secret）
3. 已正确配置provider

## 配置说明

在使用之前，您需要配置以下认证信息：

- `access_key`: 阿里云访问密钥ID (Access Key ID)
- `access_secret`: 阿里云访问密钥密码 (Access Key Secret)
- `endpoint`: 云效API接入点（可选，默认为杭州区域 https://devops.cn-hangzhou.aliyuncs.com）

您可以通过以下方式获取访问密钥：
1. 登录阿里云控制台
2. 点击右上角的账号
3. 选择"AccessKey管理"
4. 创建或查看您的AccessKey

注意：请妥善保管您的AccessKey信息，不要泄露给他人。

这个示例创建了一个包含以下内容的流水线：

- 流水线名称：simple-shell-flow
- 包含一个阶段(stage)：执行Shell命令
- 阶段中包含一个shell任务，执行 `echo 'Hello, YunXiao Flow!'` 命令
- 任务超时时间设置为3600秒（1小时）

## 安装和使用

1. 编译provider:
```bash
# Linux/MacOS
go build -o terraform-provider-yunxiao_v1.0.0

# Windows
go build -o terraform-provider-yunxiao_v1.0.0.exe
```

2. 创建本地plugin目录并复制provider:

Linux/MacOS:
```bash
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/aliyun/yunxiao/1.0.0/$(go env GOOS)_$(go env GOARCH)
cp terraform-provider-yunxiao_v1.0.0 ~/.terraform.d/plugins/registry.terraform.io/aliyun/yunxiao/1.0.0/$(go env GOOS)_$(go env GOARCH)/
```

Windows:
```powershell
mkdir -p "$env:APPDATA/terraform.d/plugins/registry.terraform.io/aliyun/yunxiao/1.0.0/windows_amd64"
copy terraform-provider-yunxiao_v1.0.0.exe "$env:APPDATA/terraform.d/plugins/registry.terraform.io/aliyun/yunxiao/1.0.0/windows_amd64/"
```

3. 初始化Terraform:
```bash
terraform init
```

## 配置参数说明

- `name`: 流水线名称
- `description`: 流水线描述
- `config`: 流水线配置
  - `stages`: 流水线阶段配置
    - `name`: 阶段名称
    - `stage_name`: 阶段标识
    - `jobs`: 阶段中的任务列表
      - `name`: 任务名称
      - `job_type`: 任务类型（这里是shell）
      - `timeout`: 任务超时时间（秒）
      - `commands`: 要执行的shell命令列表

## 注意事项

1. 确保access_token具有创建流水线的权限
2. 建议在正式使用前先用plan命令查看配置效果
3. 该示例仅作为参考，实际使用时可能需要根据具体需求调整配置 