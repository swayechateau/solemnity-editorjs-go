package editorjs

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/swayechateau/solemnity-editorjs-go/blocks"
)

// Block represents a single block of content in Editor.js JSON.
type EditorJs struct {
	Time    int64          `json:"time"`
	Blocks  []blocks.Block `json:"blocks"`
	Version string         `json:"version"`
}

type JsonEditorJs struct {
	Time    int64              `json:"time"`
	Blocks  []blocks.JsonBlock `json:"blocks"`
	Version string             `json:"version"`
}

func (e *EditorJs) ToHtml() string {
	var html strings.Builder
	for _, block := range e.Blocks {
		html.WriteString(block.ToHtml())
	}
	return html.String()
}

func (e *JsonEditorJs) ToHtml() string {
	var html strings.Builder
	for _, block := range e.Blocks {
		html.WriteString(block.ToHtml())
	}
	return html.String()
}

func (e *JsonEditorJs) ToEditorJs() EditorJs {
	var editorJs EditorJs
	var blocks []blocks.Block
	editorJs.Time = e.Time
	editorJs.Version = e.Version
	for _, block := range e.Blocks {
		blocks = append(blocks, block.ToBlock())
	}
	editorJs.Blocks = blocks
	return editorJs
}

func ConvertToEditorJs(data interface{}) (EditorJs, error) {
	editorJs, err := Convert(data)
	if err != nil {
		fmt.Println("Error:", err)
		return editorJs.ToEditorJs(), err
	}
	return editorJs.ToEditorJs(), nil
}

func Convert(data interface{}) (JsonEditorJs, error) {
	var editorJs JsonEditorJs

	var err error
	var bytes []byte
	bytes, err = toJson(data)
	if err != nil {
		fmt.Println("Error:", err)
		return editorJs, err
	}
	err = json.Unmarshal(bytes, &editorJs)
	if err != nil {
		fmt.Println("Error:", err)
		return editorJs, err
	}
	return editorJs, nil
}

func toJson(p interface{}) ([]byte, error) {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return bytes, err
	}

	return bytes, nil
}
