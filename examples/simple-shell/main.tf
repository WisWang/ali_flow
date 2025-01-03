# 配置 provider
provider "yunxiao" {
  # 从环境变量读取凭证
  # YUNXIAO_ACCESS_KEY
  # YUNXIAO_ACCESS_SECRET
  endpoint = "https://devops.cn-hangzhou.aliyuncs.com"
}



# 创建一个简单的flow
resource "yunxiao_flow" "simple_shell" {
  name        = "demo5"
  description = "一个简单的shell命令执行流水线"
  
  config = {
    stages = jsonencode([
      {
        name      = "执行Shell命令"
        stageName = "run_shell"
        jobs = [
          {
            name     = "shell任务"
            jobType  = "shell"
            timeout  = 3600
            commands = [
              "echo 'Hello, YunXiao Flow!'"
            ]
          }
        ]
      }
    ])
  }
}

# 输出flow的信息
output "flow_id" {
  value = yunxiao_flow.simple_shell.id
  description = "创建的Flow ID"
} 