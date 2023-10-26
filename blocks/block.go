package blocks

type BlockData interface {
	ToHtml() string
	ToMarkdown() string
}

type Tune struct {
	Footnotes []string `json:"footnotes,omitempty"`
}

type Block struct {
	Id    string    `json:"id,omitempty"`
	Type  string    `json:"type"`
	Data  BlockData `json:"data"`
	Tunes Tune      `json:"tunes,omitempty"`
}

func (b *Block) ToHtml() string {
	return b.Data.ToHtml()
}

func (b *Block) ToMarkdown() string {
	return b.Data.ToMarkdown()
}
