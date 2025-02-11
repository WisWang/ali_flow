# Nginx Web Helm Release

这个目录包含了用于部署 Nginx Web 服务的 Helm Release 配置。




Dockerfile 如下

```
FROM busybox:latest 
ADD  dist /opt/ 
```
前端打包后，将 dist 目录下的文件放到 /opt 下，之后通过 initContainer 将 dist 目录下的文件复制到 nginx 的静态目录下。


helm 配置如下

```
set -e
chartBase="helm-releases/nginx-web"
# 如果没有额外的变量，使用下面的命令
# 替换 Chart.yaml 中的 appVersion 为 CI_COMMIT_ID_1-prod
sed -i "s/appVersion: \"[^\"]*\"/appVersion: \"${CI_COMMIT_ID_2}-prod\"/" $chartBase/Chart.yaml
# 打印镜像信息
echo "${image}"
# 查看修改后的 Chart.yaml 文件内容
cat $chartBase/Chart.yaml
# 定义 Helm 安装和升级的公共参数
HELM_COMMON_ARGS="-f web/values_files/${PIPELINE_NAME}/values.yaml --install --namespace=default"
# 定义镜像参数
IMAGE_PARAM="--set image=registry-vpc.cn-beijing.aliyuncs.com/wis/$PIPELINE_NAME:${CI_COMMIT_ID_2}-${BUILD_NUMBER}"
# 执行 Helm 升级命令
helm upgrade $HELM_COMMON_ARGS $IMAGE_PARAM $PIPELINE_NAME $chartBase


```
