package deepseek

// 请求
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`

	Prefix           bool   `json:"prefix,omitempty"`
	ReasoningContent string `json:"reasoning_content,omitempty"`

	Name string `json:"name,omitempty"`

	ToolCalls  []*ToolCall `json:"tool_calls,omitempty"`
	ToolCallID string      `json:"tool_call_id,omitempty"`
}

type Tool struct {
	Type     string   `json:"type"`
	Function Function `json:"function"`
}
type Function struct {
	Description string      `json:"description"`
	Name        string      `json:"name"`
	Parameters  *JSONSchema `json:"parameters"`
}

type ChatCompletionRequest struct {
	Model        string    `json:"model"`
	Messages     []Message `json:"messages"`
	Stream       bool      `json:"stream"`
	ResponseType string    `json:"response_type"`

	Tools                    []*Tool `json:"tools,omitempty"`
	ChatCompletionToolChoice string  `json:"chatCompletionToolChoice,omitempty"`
	Temperature              float64 `json:"temperature,omitempty"`
	TopP                     float64 `json:"top_p,omitempty"`
	MaxTokens                int     `json:"max_tokens,omitempty"`
	N                        int     `json:"n,omitempty"`
	PresencePenalty          float64 `json:"presence_penalty,omitempty"`
	FrequencyPenalty         float64 `json:"frequency_penalty,omitempty"`
	User                     string  `json:"user,omitempty"`
}

// 响应
type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index int `json:"index"`
		// not stream
		Message Message `json:"message,omitempty"`
		// stream
		Delta        Delta  `json:"delta,omitempty"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage             Usage  `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}

type Delta struct {
	Role             string `json:"role"`
	Content          string `json:"content"`
	ReasoningContent string `json:"reasoning_content"`
 Name string `json:"name,omitempty"`

 ToolCalls  []*ToolCall `json:"tool_calls,omitempty"`
 ToolCallID string      `json:"tool_call_id,omitempty"`
}

type ToolCall struct {
	ID       string       `json:"id,omitempty"`
	Type     string       `json:"type"`
	Function FunctionCall `json:"function"`
}

type FunctionCall struct {
	Name      string `json:"name,omitempty"`
	Arguments string `json:"arguments,omitempty"`
}

type Usage struct {
	PromptTokens     float64 `json:"prompt_tokens"`
	CompletionTokens float64 `json:"completion_tokens"`
	TotalTokens      float64 `json:"total_tokens"`
}
