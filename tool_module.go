package deepseek

type JSONSchema struct {
	Type        string                `json:"type,omitempty"`
	Description string                `json:"description,omitempty"`
	Properties  map[string]JSONSchema `json:"properties,omitempty"`
	Required    []string              `json:"required,omitempty"`
	Items       *JSONSchema           `json:"items,omitempty"` // 用于数组类型
	Enum        []interface{}         `json:"enum,omitempty"`
}
