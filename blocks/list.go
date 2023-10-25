package blocks

import (
	"log"
)

type List struct {
	Items []string `json:"items"`
	Style string   `json:"type"`
}

func (l List) ToHtml() string {
	html := "<" + listTag(l.Style) + ">"
	for _, item := range l.Items {
		html += "<li>" + item + "</li>"
	}
	html += "</" + listTag(l.Style) + ">"
	return html
}

func listTag(style string) string {
	switch style {
	case "ordered":
		return "ol"
	default:
		return "ul"
	}
}

func MapToList(data map[string]interface{}) List {
	var list List
	style := "unordered"
	if data["style"] != nil {
		style = data["style"].(string)
	}
	if data["type"] != nil {
		style = data["type"].(string)
	}
	list.Style = style

	log.Printf("list.Style: %s", list.Style)
	items := data["items"].([]interface{})

	for _, item := range items {
		list.Items = append(list.Items, item.(string))
	}

	return list
}
