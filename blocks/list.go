package blocks

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
	items := data["items"].([]interface{})
	var list List
	for _, item := range items {
		list.Items = append(list.Items, item.(string))
	}
	list.Style = data["type"].(string)
	return list
}
