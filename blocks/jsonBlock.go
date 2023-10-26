package blocks

import "log"

type JsonBlock struct {
	Id    string                 `json:"id,omitempty"`
	Type  string                 `json:"type"`
	Data  map[string]interface{} `json:"data"`
	Tunes Tune                   `json:"tunes,omitempty"`
}

func (b *JsonBlock) ToBlock() Block {
	var data BlockData
	log.Println(b.Type)
	switch b.Type {
	case "paragraph":
		data = MapToParagraph(b.Data)
	case "header":
		data = MapToHeader(b.Data)
	case "list":
		data = MapToList(b.Data)
	case "table":
		data = MapToTable(b.Data)
	case "image":
		data = MapToImage(b.Data)
	case "attaches":
		data = MapToAttaches(b.Data)
	case "code":
		data = MapToCode(b.Data)
	case "quote":
		data = MapToQuote(b.Data)
	case "linkTool", "link":
		data = MapToLink(b.Data)
	case "embed":
		data = MapToEmbed(b.Data)
	}
	log.Println(data)
	return Block{
		Id:    b.Id,
		Type:  b.Type,
		Data:  data,
		Tunes: b.Tunes,
	}
}
