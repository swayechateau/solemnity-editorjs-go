package editorjs

import (
	"solemnity-editorjs-go/blocks"
	"strings"
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

// func Convert() string {

// }
