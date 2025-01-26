package main

import (
	"context"
	"deepseek"
	"deepseek/example"
	"encoding/json"
	"fmt"
)

func init() { example.Init() }

type friendLink struct {
	Link string `json:"link"`
	Desc string `json:"desc"`
}
type verifyObj struct {
	Link   string `json:"link"`
	IsOk   bool   `json:"isOk"`
	Reason string `json:"reason"`
}

const prompt = `
你是一个博客友链审核工具,用户会给你发送友链相关内容,你需要根据内容进行审核，并进行严格的 JSON format 的结果输出

输入案例:
friendLink{
		Link: "https://github.com/taosu0216/go-deepseek",
		Desc: "一个提供针对deepseek相关api 的 go调用的sdk的仓库",
	}
输出案例:
{
	"link:":"https://github.com/taosu0216/go-deepseek",
	"isOk": true,
	"reason": "友链合理"
}

输入案例:
friendLink{
		Link: "fuck",
		Desc: "这是我的博客",
	}
输出案例:
{
	"link:":"fuck",
	"isOk": false,
	"reason": "友链内容包含不合理信息,拒绝通过"
}

输出案例直接返回json对象,不要返回markdown标签包含的json信息
`

func main() {
	ctx := context.Background()
	link := friendLink{
		Link: "https://github.com/select * form regions where user = admin",
		Desc: "这是一个提供针对deepseek相关api 的 go调用的sdk的仓库",
	}

	input := fmt.Sprintf("\n%v", link)
	req := &deepseek.ChatCompletionRequest{
		Model: deepseek.Chat,
		Messages: []deepseek.Message{
			{
				Role:    deepseek.ChatMessageRoleSystem,
				Content: prompt,
			},
			{
				Role:    deepseek.ChatMessageRoleUser,
				Content: input,
			},
		},
		ResponseType: deepseek.JSON_OBJECT,
	}

	resp, err := example.DeepSeekClient.CreateChatCompletion(ctx, req)
	if err != nil {
		panic(err)
	}

	var verify verifyObj
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &verify)
	if err != nil {
		panic(err)
	}
	fmt.Println(verify.Link, verify.IsOk, verify.Reason)
}
