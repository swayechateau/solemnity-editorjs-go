package blocks

import (
	"fmt"
	"strings"
)

type Image struct {
	File           *File   `json:"file"`
	URL            *string `json:"url,omitempty"`
	Caption        *string `json:"caption,omitempty"`
	WithBorder     *bool   `json:"withBorder,omitempty"`
	WithBackground *bool   `json:"withBackground,omitempty"`
	Stretched      *bool   `json:"stretched,omitempty"`
}

func MapToImage(data map[string]interface{}) Image {
	var image Image
	if data["file"] != nil {
		file := MapToFile(data["file"].(map[string]interface{}))
		image.File = &file
	}
	if data["url"] != nil {
		url := data["url"].(string)
		image.URL = &url
	}
	if data["caption"] != nil {
		caption := data["caption"].(string)
		image.Caption = &caption
	}
	if data["withBorder"] != nil {
		withBorder := data["withBorder"].(bool)
		image.WithBorder = &withBorder
	}
	if data["withBackground"] != nil {
		withBackground := data["withBackground"].(bool)
		image.WithBackground = &withBackground
	}
	if data["stretched"] != nil {
		stretched := data["stretched"].(bool)
		image.Stretched = &stretched
	}
	return image
}

func (i Image) ToHtml() string {
	var html strings.Builder
	var url string
	caption := ""
	html.WriteString("<figure>")
	if i.File != nil {
		url = i.File.URL
	}
	if i.URL != nil {
		url = *i.URL
	}
	if i.Caption != nil {
		caption = *i.Caption

	}
	html.WriteString(fmt.Sprintf("<img src=\"%s\" alt=\"%s\">", url, caption))
	html.WriteString(fmt.Sprintf("<figcaption>%s</figcaption>", caption))
	html.WriteString("</figure>")
	return html.String()
}

// func (i Image) FileToHtml() string {
// 	if i.File != nil {
// 		// check file url
// 		return "<img src=\"" + i.File.URL + "\" alt=\"\">
// }

func (i Image) ToMarkdown() string {
	var markdown strings.Builder
	if i.Caption != nil {
		markdown.WriteString(fmt.Sprintf("![%s]", *i.Caption))
	} else {
		markdown.WriteString("![")
	}
	if i.File != nil {
		markdown.WriteString(i.File.URL)
	}
	if i.URL != nil {
		markdown.WriteString(*i.URL)
	}
	markdown.WriteString("]")
	return markdown.String()
}
