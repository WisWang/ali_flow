# 云效流水线金丝雀发布 Helm Chart

在 /canary-relase 下是一个用于云效流水线金丝雀发布的 Helm Chart，支持应用的灰度发布和流量控制。

## 目录结构

```
canary-relase/
├── Chart.yaml          # Chart 元数据文件，包含版本信息
├── templates/          # Kubernetes 资源模板目录
└── values.yaml         # 默认配置值
```

## 功能特性

- 支持通过云效流水线进行金丝雀发布
- 自动更新 Chart 版本和应用版本
- 支持金丝雀版本和稳定版本的切换
- 提供自动回滚机制

## 使用方法

### 在云效流水线中使用

1. 配置流水线命令：
```yaml
step_2:
  name: 执行命令
  step: Command
  workspace: repo_1
  with:
    ifGivenShell: false
    run: |
      # 设置工作目录
      cd canary-relase 
      
      # 更新 Chart.yaml 中的 appVersion
      sed -i "s/appVersion: \".*\"/appVersion: \"${CI_COMMIT_ID_1}-prod\"/" Chart.yaml
      
      # 打印使用的镜像和更新后的 Chart.yaml 内容
      echo "Using image: ${image}"
      cat Chart.yaml
      
      # 根据是否为金丝雀发布设置应用名称
      if [ $canary == "true" ]; then
        helm_app_name=${PIPELINE_NAME}-canary
      else
        helm_app_name=${PIPELINE_NAME}
        helm --namespace=default uninstall ${PIPELINE_NAME}-canary || date
      fi
      
      # 安装或升级 Helm 发布
      helm upgrade $helm_opts \
        --namespace=default \
        --set canary=$canary \
        $helm_app_name .
```

### 配置说明

#### 环境变量

| 变量名 | 描述 | 示例值 |
|--------|------|--------|
| `CI_COMMIT_ID_1` | 代码提交ID | `abc123-prod` |
| `PIPELINE_NAME` | 流水线名称 | `my-app` |
| `canary` | 是否为金丝雀发布 | `true/false` |
| `helm_opts` | Helm 额外参数 | `--install/--install  --dry-run --debug` |

#### Chart 配置

| 参数 | 描述 | 默认值 |
|------|------|--------|
| `image.tag` | 容器镜像标签 | `latest` |

## 发布流程

1. 稳定版本发布：
这步会删除金丝雀发布的版本

```bash
# canary=false
helm_app_name=${PIPELINE_NAME}
helm --namespace=default uninstall ${PIPELINE_NAME}-canary || date
helm upgrade $helm_opts \
  --namespace=default \
  --set canary=$canary \
  $helm_app_name .

```

2. 金丝雀发布：
```bash
# canary=true
helm_app_name=${PIPELINE_NAME}-canary
helm upgrade $helm_opts \
  --namespace=default \
  --set canary=$canary \
  $helm_app_name .

```



## 注意事项

1. 确保 Chart.yaml 中的版本号格式正确
2. 金丝雀版本会使用 `-canary` 后缀
3. 稳定版本发布时会自动清理金丝雀版本
4. 使用 `helm_opts` 可以传递额外的 Helm 参数

# Terraform Provider for 云效Flow

这是一个用于管理阿里云云效Flow的Terraform Provider，允许用户通过Terraform来创建和管理云效Flow的配置。

## 代码结构

```
terraform-provider-yunxiao/
├── main.go                 # provider的入口文件
├── go.mod                  # Go模块定义文件
├── README.md              # 项目说明文档
└── yunxiao/               # provider核心代码目录
    ├── provider.go        # provider定义，包含认证配置
    ├── resource_flow.go   # Flow资源的CRUD实现
    └── client.go          # 云效API客户端实现
```

## 主要组件说明

1. **main.go**
   - provider的入口点
   - 使用terraform-plugin-sdk注册provider

2. **yunxiao/provider.go**
   - 定义provider的配置项（access_token和endpoint）
   - 注册可用的资源（目前包含yunxiao_flow）

3. **yunxiao/client.go**
   - 实现与云效API的交互
   - 定义API客户端结构和方法
   - 处理HTTP请求和响应

4. **yunxiao/resource_flow.go**
   - 定义Flow资源的schema
   - 实现资源的CRUD操作
   - 处理资源状态的转换

## 配置示例

```hcl
# 配置provider
provider "yunxiao" {
  access_token = "your_access_token"
  endpoint     = "https://devops.aliyun.com"  # 可选
}

# 创建Flow
resource "yunxiao_flow" "example" {
  name        = "my-pipeline"
  description = "My CI/CD Pipeline"
  
  config = {
    "stages" = jsonencode([
      {
        "name" = "Build"
        "steps" = [
          {
            "type" = "shell"
            "commands" = [
              "npm install",
              "npm run build"
            ]
          }
        ]
      }
    ])
  }
}
```

## 安装和使用

1. 编译provider:
```bash
go build -o terraform-provider-yunxiao
```

2. 将编译好的provider放到Terraform插件目录:
```bash
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/aliyun/yunxiao/1.0.0/linux_amd64
cp terraform-provider-yunxiao ~/.terraform.d/plugins/registry.terraform.io/aliyun/yunxiao/1.0.0/linux_amd64/
```

3. 创建Terraform配置文件并使用provider

## 待完善功能

- [ ] 实现Flow资源的Read操作
- [ ] 实现Flow资源的Update操作
- [ ] 实现Flow资源的Delete操作
- [ ] 添加资源导入功能
- [ ] 添加更多的验证逻辑
- [ ] 完善错误处理
- [ ] 添加重试机制
- [ ] 添加单元测试
- [ ] 支持更多的Flow配置选项

## 注意事项

1. 确保有正确的云效访问令牌（access_token）
2. 当前实现是基础框架，需要根据实际的云效Flow API进行调整
3. 建议在生产环境使用前完善错误处理和重试机制
4. 使用前请确保了解云效Flow的配置规范

## 贡献

欢迎提交Issue和Pull Request来帮助改进这个provider。

## 许可证

[MIT License](LICENSE)
## 开发模式

在开发过程中，可以使用以下方式直接运行源代码而不需要编译：

1. 在一个终端中运行 provider：
```bash
go run main.go
```

2. 在另一个终端中设置环境变量：

Linux/MacOS:
```bash
export TF_CLI_CONFIG_FILE=$(pwd)/../../dev.tfrc
terraform init
terraform plan
```

Windows (PowerShell):
```powershell
$env:TF_CLI_CONFIG_FILE = "$(Get-Location)\dev\dev.tfrc"
terraform init
terraform plan
```

这样设置后，修改代码后直接运行 terraform 命令即可看到效果，无需重新编译。
