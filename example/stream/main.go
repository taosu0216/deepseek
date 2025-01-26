package main

import (
	"context"
	"deepseek"
	"deepseek/example"
	"errors"
	"fmt"
	"io"
)

func init() { example.Init() }

type contentObj struct {
	content string
	flag    string
}

func main() {
	input := "介绍一下你自己"
	ctx := context.Background()
	isShowReasoningContent := true

	contentCh := make(chan contentObj, 50)
	go func() {
		for {
			content := <-contentCh
			fmt.Printf("[%s]: %s\n", content.flag, content.content)
		}
	}()

	req := deepseek.ChatCompletionRequest{
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
		Stream: true,
	}

	stream, err := example.DeepSeekClient.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}
		content := ""
		flag := ""
		if response.Choices[0].Delta.ReasoningContent != "" && isShowReasoningContent {
			content = response.Choices[0].Delta.ReasoningContent
			flag = "reasoning_content"
		} else if response.Choices[0].Delta.Content != "" {
			content = response.Choices[0].Delta.Content
			flag = "content"
		} else {
			continue
		}
		contentCh <- contentObj{
			content: content,
			flag:    flag,
		}
	}
}
