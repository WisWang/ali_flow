package yunxiao

import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type Client struct {
	AccessKey    string
	AccessSecret string
	Endpoint     string
	client       *sdk.Client
}

type Flow struct {
	Name        string                 `json:"flowName"`
	Description string                 `json:"description"`
	Config      map[string]interface{} `json:"config"`
}

func NewClient(accessKey, accessSecret, endpoint string) (*Client, error) {
	client, err := sdk.NewClientWithAccessKey(
		"cn-hangzhou",  // 区域ID
		accessKey,      // AccessKey ID
		accessSecret,   // AccessKey Secret
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		AccessKey:    accessKey,
		AccessSecret: accessSecret,
		Endpoint:     endpoint,
		client:       client,
	}, nil
}

func (c *Client) CreateFlow(flow *Flow) error {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = strings.TrimPrefix(c.Endpoint, "https://")
	request.Domain = strings.TrimPrefix(request.Domain, "http://")
	request.PathPattern = "/organization/67707f8935730943ed495595/pipelines"
	request.Headers["Content-Type"] = "application/json"
	
	// 设置 API 版本
	request.Version = "2021-06-25"
	request.ApiName = "CreatePipeline"
	request.Product = "devops"

	// 准备请求参数
	requestBody := map[string]interface{}{
		"name": flow.Name,
		"content": fmt.Sprintf(`
stages:
  %s:  # 使用 stageName 作为 key
    name: "%s"
    jobs:
      shell_job:  # 使用固定的 job key
        name: "%s"
        type: "Shell"
        runsOn: "public/cn-hangzhou"
        commands:
          - %s
`, "run_shell", "执行Shell命令", "shell任务", "echo 'Hello, YunXiao Flow!'"),
	}
	
	body, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	request.Content = body

	// 添加调试信息
	fmt.Printf("Request URL: %s\n", request.Domain+request.PathPattern)
	fmt.Printf("Request Headers: %v\n", request.Headers)
	fmt.Printf("Request Body: %s\n", string(body))

	response, err := c.client.ProcessCommonRequest(request)
	if err != nil {
		return fmt.Errorf("API request failed: %v", err)
	}

	if !response.IsSuccess() {
		return fmt.Errorf("API request failed with status code: %d, response: %s", 
			response.GetHttpStatus(), response.GetHttpContentString())
	}

	return nil
}

func (c *Client) UpdateFlow(id string, flow *Flow) error {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = strings.TrimPrefix(c.Endpoint, "https://")
	request.Domain = strings.TrimPrefix(request.Domain, "http://")
	request.PathPattern = fmt.Sprintf("/organization/%s/pipelines/update", "67707f8935730943ed495595")
	request.Headers["Content-Type"] = "application/json"
	
	// 设置 API 版本
	request.Version = "2021-06-25"
	request.ApiName = "UpdatePipeline"
	request.Product = "devops"

	// 准备请求参数
	requestBody := map[string]interface{}{
		"name": flow.Name,
		"pipelineId": id,
		"content": fmt.Sprintf(`
stages:
  %s:  # 使用 stageName 作为 key
    name: "%s"
    jobs:
      shell_job:  # 使用固定的 job key
        name: "%s"
        type: "Shell"
        runsOn: "public/cn-hangzhou"
        commands:
          - %s
`, "run_shell", "执行Shell命令", "shell任务", "echo 'Hello, YunXiao Flow!'"),
	}
	
	body, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	request.Content = body

	response, err := c.client.ProcessCommonRequest(request)
	if err != nil {
		return fmt.Errorf("API request failed: %v", err)
	}

	if !response.IsSuccess() {
		return fmt.Errorf("API request failed with status code: %d, response: %s", 
			response.GetHttpStatus(), response.GetHttpContentString())
	}

	return nil
} 