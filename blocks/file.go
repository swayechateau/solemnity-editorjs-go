package blocks

type File struct {
	URL       string  `json:"url"`
	Size      *int    `json:"size"`
	Extension *string `json:"extension"`
	Name      *string `json:"name"`
}

func MapToFile(data map[string]interface{}) File {
	var size int
	if data["size"] != nil {
		size = int(data["size"].(float64))
	}
	var extension string
	if data["extension"] != nil {
		extension = data["extension"].(string)
	}
	var name string
	if data["name"] != nil {
		name = data["name"].(string)
	}
	return File{
		URL:       data["url"].(string),
		Size:      &size,
		Extension: &extension,
		Name:      &name,
	}
}
