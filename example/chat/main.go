package main

import (
	"context"
	"deepseek"
	"deepseek/example"
	"fmt"
)

func init() { example.Init() }

func main() {
	ctx := context.Background()

	input := "say hello"
	req := &deepseek.ChatCompletionRequest{
		Model: deepseek.Reasoner,
		Messages: []deepseek.Message{
			{
				Role:    deepseek.ChatMessageRoleSystem,
				Content: deepseek.Prompt,
			},
			{
				Role:    deepseek.ChatMessageRoleUser,
				Content: input,
			},
		},
	}

	resp, err := example.DeepSeekClient.CreateChatCompletion(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Choices[0].Message.Content)
	fmt.Println()
	fmt.Println(resp)
}
