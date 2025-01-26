package deepseek

func NewParameters() *JSONSchema {
	return &JSONSchema{
		Type:       "object",
		Properties: make(map[string]JSONSchema),
	}
}

// WithProperty(age,deepseek.ToolParamTypeInt,"年龄")
// WithProperty(password,deepseek.ToolParamTypeStr,"用户名")

func (p *JSONSchema) WithProperty(name, paramType, description string) *JSONSchema {
	p.Properties[name] = JSONSchema{
		Type:        paramType,
		Description: description,
	}
	return p
}

func (p *JSONSchema) WithRequired(fields ...string) *JSONSchema {
	p.Required = append(p.Required, fields...)
	return p
}

func NewTool(name, description string, params *JSONSchema) *Tool {
	return &Tool{
		Type: ChatCompletionToolTypeFunc,
		Function: Function{
			Description: description,
			Name:        name,
			Parameters:  params,
		},
	}
}
