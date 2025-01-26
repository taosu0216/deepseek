package deepseek

const (
	Reasoner = "deepseek-reasoner"
	Chat     = "deepseek-chat"
)

const (
	ChatMessageRoleSystem    = "system"
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
	ChatMessageRoleFunction  = "function"
	ChatMessageRoleTool      = "tool"
)

const (
	TEXT        = "text"
	JSON_OBJECT = "json_object"
)

const (
	ChatCompletionToolChoiceNone     = "none"
	ChatCompletionToolChoiceAuto     = "auto"
	ChatCompletionToolChoiceRequired = "required"
)

const (
	ChatCompletionToolTypeFunc = "function"

	ToolParamTypeInt     = "int"
	ToolParamTypeFloat64 = "float64"
	ToolParamTypeStr     = "string"
	ToolParamTypeArr     = "array"
	ToolParamTypeObj     = "object"
)
