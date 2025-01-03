# 云效流水线金丝雀发布 Helm Chart

这是一个用于云效流水线金丝雀发布的 Helm Chart，支持应用的灰度发布和流量控制。

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
| `canary` | 是否启用金丝雀发布 | `false` |
| `image.tag` | 容器镜像标签 | `latest` |

## 发布流程

1. 稳定版本发布：
```bash
# canary=false
helm_app_name=${PIPELINE_NAME}
helm upgrade --install --namespace=default $helm_app_name .
```

2. 金丝雀发布：
```bash
# canary=true
helm_app_name=${PIPELINE_NAME}-canary
helm upgrade --install --namespace=default --set canary=true $helm_app_name .
```

3. 金丝雀回滚：
```bash
# 删除金丝雀版本
helm --namespace=default uninstall ${PIPELINE_NAME}-canary
```

## 注意事项

1. 确保 Chart.yaml 中的版本号格式正确
2. 金丝雀版本会使用 `-canary` 后缀
3. 稳定版本发布时会自动清理金丝雀版本
4. 使用 `helm_opts` 可以传递额外的 Helm 参数

## 最佳实践

1. 总是使用明确的版本标签
2. 保持 Chart.yaml 结构的一致性
3. 使用 `--atomic` 参数确保原子性部署
4. 合理设置资源限制和请求

## 故障排查

1. 检查 Chart.yaml 版本更新是否成功
2. 验证环境变量是否正确设置
3. 查看 Helm 部署历史和状态
4. 检查 Kubernetes 集群状态

## 许可证

[MIT License](LICENSE) 