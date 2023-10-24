package blocks

type Paragraph struct {
	Text string `json:"text"`
}

func (p Paragraph) ToHtml() string {
	return "<p>" + p.Text + "</p>"
}

func (p Paragraph) ToMarkdown() string {
	return p.Text
}

func MapToParagraph(data map[string]interface{}) Paragraph {
	return Paragraph{
		Text: data["text"].(string),
	}
}
