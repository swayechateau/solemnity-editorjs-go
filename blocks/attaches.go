package blocks

type Attaches struct {
	File  File   `json:"file"`
	Title string `json:"title"`
}

func (a Attaches) ToHtml() string {
	return "<a href=\"" + a.File.URL + "\">" + a.Title + "</a>"
}

func (a Attaches) ToMarkdown() string {
	return "[" + a.Title + "](" + a.File.URL + ")"
}

func MapToAttaches(data map[string]interface{}) Attaches {
	return Attaches{
		File:  MapToFile(data["file"].(map[string]interface{})),
		Title: data["title"].(string),
	}
}
