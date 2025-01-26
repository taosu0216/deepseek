package example

import "deepseek"

var (
	DeepSeekClient *deepseek.Client
)

func Init() {
	DeepSeekClient = deepseek.NewClient("sk-********************************")
}
