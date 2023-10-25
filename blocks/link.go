package blocks

type Link struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

func MapToLink(data map[string]interface{}) Link {
	link := "#"
	text := ""
	if data["text"] != nil {
		text = data["text"].(string)
	}
	if data["link"] != nil {
		link = data["link"].(string)
	}
	return Link{
		Text: text,
		Link: link,
	}
}

func (l Link) ToHtml() string {
	return "<a href=\"" + l.Link + "\" alt=\"" + l.Text + "\">" + l.Text + "</a>"
}

func (l Link) ToMarkdown() string {
	return "[" + l.Text + "](" + l.Link + ")"
}
