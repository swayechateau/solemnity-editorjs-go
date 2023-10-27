package blocks

type Table struct {
	WithHeadings bool       `json:"withHeadings"`
	Content      [][]string `json:"content"`
}

func (t Table) ToHtml() string {
	html := "<table>"
	header, body := t.Content[0], t.Content[1:]

	if t.WithHeadings {
		html += genHeaders(header)
		html += genBody(body)
	} else {
		html += genBody(t.Content)
	}
	html += "</table>"
	return html
}

func (t Table) ToMarkdown() string {
	markdown := ""
	for _, row := range t.Content {
		for _, cell := range row {
			markdown += cell + " | "
		}
		markdown += "\n"
	}
	return markdown
}

func MapToTable(data map[string]interface{}) Table {
	var table Table
	withHeadings := false
	if data["withHeadings"] != nil {
		withHeadings = data["withHeadings"].(bool)
	}

	table.WithHeadings = withHeadings
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

func genHeaders(header []string) string {
	html := "<thead><tr>"
	for _, heading := range header {
		html += "<th>" + heading + "</th>"
	}
	html += "</tr></thead>"

	return html
}

func genBody(body [][]string) string {
	html := "<tbody>"
	for _, row := range body {
		html += "<tr>"
		for _, cell := range row {
			html += "<td>" + cell + "</td>"
		}
		html += "</tr>"
	}
	html += "</tbody>"
	return html
}
