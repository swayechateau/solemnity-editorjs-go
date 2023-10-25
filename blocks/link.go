package blocks

type Link struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

func MapToLink(data map[string]interface{}) Link {
	return Link{
		Text: data["text"].(string),
		Link: data["link"].(string),
	}
}

func (l Link) ToHtml() string {
	return "<a href=\"" + l.Link + "\" alt=\"" + l.Text + "\">" + l.Text + "</a>"
}

func (l Link) ToMarkdown() string {
	return "[" + l.Text + "](" + l.Link + ")"
}
