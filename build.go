package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 钉钉机器人访问令牌
	accessToken := "c82ba73f86f5fb9b24c9dc2c8ef6311a26c8ebbaa08e6a413cd8f909fae75f96"
	// 钉钉机器人发送消息的URL
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", accessToken)

	// 要发送的消息内容
	message := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": "这是一个测试消息，关键词是test。",
		},
	}

	// 将消息内容编码为JSON
	messageBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("JSON编码错误: %v\n", err)
		return
	}

	// 创建一个HTTP POST请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(messageBytes))
	if err != nil {
		fmt.Printf("创建请求错误: %v\n", err)
		return
	}

	// 设置请求头信息
	req.Header.Set("Content-Type", "application/json")

	// 发送HTTP请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求错误: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应错误: %v\n", err)
		return
	}

	// 打印响应状态码和响应体
	fmt.Printf("响应状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应体: %s\n", body)
}
