package blocks

type Embed struct {
	Service string `json:"service"`
	Source  string `json:"source"`
	Width   string `json:"width"`
	Height  string `json:"height"`
}

func MapToEmbed(data map[string]interface{}) Embed {
	width := "100%"
	height := "100%"
	if data["width"] != nil {
		height = data["height"].(string)

	}
	if data["height"] != nil {
		width = data["width"].(string)
	}
	return Embed{
		Service: data["service"].(string),
		Source:  data["source"].(string),
		Width:   width,
		Height:  height,
	}
}

func (e Embed) ToHtml() string {
	return "<iframe src=\"" + e.Source + "\" width=\"" + string(e.Width) + "\" height=\"" + string(e.Height) + "\"></iframe>"
}

func (e Embed) ToMarkdown() string {
	return e.ToHtml()
}
