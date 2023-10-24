package blocks

import (
	"fmt"
	"log"
	"strings"
)

type Block struct {
	Id    string      `json:"id,omitempty"`
	Type  string      `json:"type"`
	Data  interface{} `json:"data"`
	Tunes Tune        `json:"tunes,omitempty"`
}

type JsonBlock struct {
	Id    string                 `json:"id,omitempty"`
	Type  string                 `json:"type"`
	Data  map[string]interface{} `json:"data"`
	Tunes Tune                   `json:"tunes,omitempty"`
}

type Tune struct {
	Footnotes []string `json:"footnotes,omitempty"`
}

func (b *Block) ToHtml() string {
	switch b.Type {
	case "paragraph":
		p := b.Data.(Paragraph)
		return p.ToHtml()
	case "header":
		h := b.Data.(Header)
		return h.ToHtml()
	case "list":
		l := b.Data.(List)
		return l.ToHtml()
	case "table":
		t := b.Data.(Table)
		return t.ToHtml()
	}
	return ""
}

func (b *JsonBlock) ToHtml() string {
	switch b.Type {
	case "paragraph":
		return fmt.Sprintf("<p>%s</p>", b.Data["text"])
	case "header":
		return fmt.Sprintf("<h%d>%s</h%d>", int(b.Data["level"].(float64)), b.Data["text"], int(b.Data["level"].(float64)))
	case "list":
		items := b.Data["items"].([]interface{})
		var html strings.Builder
		html.WriteString("<ul>")
		for _, item := range items {
			html.WriteString(fmt.Sprintf("<li>%s</li>", item))
		}
		html.WriteString("</ul>")
		return html.String()
	default:
		return ""
	}
}

func (b *JsonBlock) ToBlock() Block {
	var data interface{}
	log.Println(b.Type)
	switch b.Type {
	case "paragraph":
		data = MapToParagraph(b.Data)
	case "header":
		data = MapToHeader(b.Data)
	case "list":
		data = MapToList(b.Data)
	case "table":
		data = MapToTable(b.Data)
	case "image":
		data = MapToImage(b.Data)
	case "file":
		data = MapToFile(b.Data)
	case "attaches":
		data = MapToAttaches(b.Data)
	}
	log.Println(data)
	return Block{
		Id:    b.Id,
		Type:  b.Type,
		Data:  data,
		Tunes: b.Tunes,
	}
}
