package blocks

type Table struct {
	WithHeadings bool       `json:"withHeadings"`
	Content      [][]string `json:"content"`
}

func (t Table) ToHtml() string {
	html := "<table>"
	if t.WithHeadings {
		html += "<thead><tr>"
		for _, heading := range t.Content[0] {
			html += "<th>" + heading + "</th>"
		}
		html += "</tr></thead>"
	}
	html += "<tbody>"
	for _, row := range t.Content {
		html += "<tr>"
		for _, cell := range row {
			html += "<td>" + cell + "</td>"
		}
		html += "</tr>"
	}
	html += "</tbody></table>"
	return html
}

func MapToTable(data map[string]interface{}) Table {
	var table Table
	table.WithHeadings = data["withHeadings"].(bool)
	var content [][]string
	for _, row := range data["content"].([]interface{}) {
		var cells []string
		for _, cell := range row.([]interface{}) {
			cells = append(cells, cell.(string))
		}
		content = append(content, cells)
	}
	table.Content = content
	return table
}
