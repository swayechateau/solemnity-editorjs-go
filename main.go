package main

import (
	"fmt"

	"github.com/swayechateau/solemnity-editorjs-go/editorjs"
)

func main() {
	e, err := editorjs.ConvertJson([]byte(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Parsed JSON Data: %+v\n", e.ToHtml())
}

var jsonData = `{
    "time": 1698090878313,
    "blocks": [
        {
            "type": "embed",
            "data": {
                "source": "https://www.youtube.com/embed/6OctHLex_Io?si=zKUGNwi53viSq_7P",
                "service": "youtube"
            }
        }
    ]
}
`
