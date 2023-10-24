package blocks

import (
	"strconv"
	"strings"
)

type Header struct {
	Text  string `json:"text"`
	Level int    `json:"level"`
}

func (h Header) ToHtml() string {
	level := strconv.Itoa(h.Level)
	return "<h" + level + ">" + h.Text + "</h" + level + ">"
}

func (h Header) ToMarkdown() string {
	return headerHash(h.Level, h.Text)
}

func headerHash(level int, text string) string {
	return strings.Repeat("#", level) + " " + text
}

func MapToHeader(data map[string]interface{}) Header {
	return Header{
		Text:  data["text"].(string),
		Level: int(data["level"].(float64)),
	}
}
