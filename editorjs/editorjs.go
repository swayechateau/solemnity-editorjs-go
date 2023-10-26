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

func (e *EditorJs) ToHtml() string {
	var html strings.Builder
	for _, block := range e.Blocks {
		html.WriteString(block.ToHtml())
	}
	return html.String()
}

func Convert(data interface{}) (EditorJs, error) {
	editorJs, err := fromInterface(data)
	if err != nil {
		fmt.Println("Error:", err)
		return editorJs.toEditorJs(), err
	}
	return editorJs.toEditorJs(), nil
}

func ConvertJson(bytes []byte) (EditorJs, error) {
	e, err := fromJson(bytes)
	if err != nil {
		fmt.Println("Error:", err)
		return e.toEditorJs(), err
	}
	return e.toEditorJs(), nil
}

type jsonEditorJs struct {
	Time    int64              `json:"time"`
	Blocks  []blocks.JsonBlock `json:"blocks"`
	Version string             `json:"version"`
}

func (e *jsonEditorJs) toEditorJs() EditorJs {
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

func fromInterface(data interface{}) (jsonEditorJs, error) {
	var editorJs jsonEditorJs

	var err error
	var bytes []byte
	bytes, err = toJson(data)
	if err != nil {
		fmt.Println("Error:", err)
		return editorJs, err
	}
	editorJs, err = fromJson(bytes)
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

func fromJson(bytes []byte) (jsonEditorJs, error) {
	var editorJs jsonEditorJs
	err := json.Unmarshal(bytes, &editorJs)
	if err != nil {
		fmt.Println("Error:", err)
		return editorJs, err
	}
	return editorJs, nil
}
