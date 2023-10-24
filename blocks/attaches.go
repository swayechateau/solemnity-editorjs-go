package blocks

type Attaches struct {
	File  File   `json:"file"`
	Title string `json:"title"`
}

func MapToAttaches(data map[string]interface{}) Attaches {
	return Attaches{
		File:  MapToFile(data["file"].(map[string]interface{})),
		Title: data["title"].(string),
	}
}

func (a Attaches) ToHtml() string {
	return "<a href=\"" + a.File.URL + "\">" + a.Title + "</a>"
}
