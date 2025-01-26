package main

import (
	"context"
	"deepseek"
	"deepseek/example"
	"encoding/json"
	"fmt"
	"time"
)

func init() { example.Init() }

// 定义工具函数类型
type ToolFunc func(string) (string, error)

// 工具函数字典
var toolMap = map[string]ToolFunc{
	"get_weather": getWeather,
	"get_time":    getTime,
}

func main() {
	ctx := context.Background()

	// 创建天气查询工具
	weatherParams := deepseek.NewParameters().
		WithProperty("location", deepseek.ToolParamTypeStr, "城市和地区，例如：青岛市").
		WithRequired("location")

	// 创建时间查询工具（无参数）
	timeParams := deepseek.NewParameters() // 无参数的工具也需要空参数对象

	tools := []*deepseek.Tool{
		deepseek.NewTool("get_weather", "获取指定地区的天气信息", weatherParams),
		deepseek.NewTool("get_time", "获取当前系统时间", timeParams),
	}

	dialog := []deepseek.Message{
		{
			Role:    deepseek.ChatMessageRoleSystem,
			Content: "你是一个智能助手，可以根据用户需求使用工具查询信息",
		},
		{
			Role:    deepseek.ChatMessageRoleUser,
			Content: "请告诉我青岛现在的天气和当前时间",
		},
	}

	// 首次请求配置工具
	req := &deepseek.ChatCompletionRequest{
		Model:                    deepseek.Chat,
		Messages:                 dialog,
		Tools:                    tools,
		ChatCompletionToolChoice: "auto",
	}

	// 处理多轮对话
	for {
		resp, err := example.DeepSeekClient.CreateChatCompletion(ctx, req)
		if err != nil {
			panic(fmt.Errorf("API请求失败: %v", err))
		}

		msg := resp.Choices[0].Message
		dialog = append(dialog, msg)

		// 如果没有工具调用则结束
		if len(msg.ToolCalls) == 0 {
			fmt.Println("最终结果:", msg.Content)
			break
		}

		// 处理所有工具调用
		for _, toolCall := range msg.ToolCalls {
			// 从字典获取工具函数
			fn, exists := toolMap[toolCall.Function.Name]
			if !exists {
				panic(fmt.Errorf("未注册的工具函数: %s", toolCall.Function.Name))
			}

			// 执行工具调用
			result, err := fn(toolCall.Function.Arguments)
			if err != nil {
				panic(fmt.Errorf("工具执行失败: %v", err))
			}

			// 添加工具响应
			dialog = append(dialog, deepseek.Message{
				Role:       deepseek.ChatMessageRoleTool,
				Content:    result,
				ToolCallID: toolCall.ID,
				Name:       toolCall.Function.Name,
			})
		}

		// 准备后续请求（不再指定工具）
		req = &deepseek.ChatCompletionRequest{
			Model:    deepseek.Chat,
			Messages: dialog,
		}
	}
}

// getWeather 天气查询工具函数
func getWeather(arguments string) (string, error) {
	type WeatherParams struct {
		Location string `json:"location"`
	}

	var params WeatherParams
	if err := json.Unmarshal([]byte(arguments), &params); err != nil {
		return "", fmt.Errorf("参数解析错误: %v", err)
	}

	// 模拟天气查询
	fmt.Printf("[工具调用] 正在查询 %s 的天气...\n", params.Location)
	return fmt.Sprintf(`{
		"location": "%s",
		"temperature": 25,
		"condition": "晴",
		"humidity": 65
	}`, params.Location), nil
}

// getTime 时间查询工具函数
func getTime(_ string) (string, error) {
	// 参数不需要使用，但保持签名一致
	fmt.Println("[工具调用] 正在获取系统时间...")
	return fmt.Sprintf(`{
		"timestamp": "%s",
		"format": "2006-01-02 15:04:05"
	}`, time.Now().Format(time.RFC3339)), nil
}
