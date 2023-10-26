package blocks

type Code struct {
	Code string `json:"code"`
}

func (c Code) ToHtml() string {
	return "<pre><code>" + c.Code + "</code></pre>"
}

func (c Code) ToMarkdown() string {
	return "```\n" + c.Code + "\n```"
}

func MapToCode(data map[string]interface{}) Code {
	return Code{
		Code: data["code"].(string),
	}
}
