package blocks

type Quote struct {
	Text string `json:"text"`
}

func MapToQuote(data map[string]interface{}) Quote {
	return Quote{
		Text: data["text"].(string),
	}
}

func (q Quote) ToHtml() string {
	return "<blockquote>" + q.Text + "</blockquote>"
}

func (q Quote) ToMarkdown() string {
	return "> " + q.Text
}
