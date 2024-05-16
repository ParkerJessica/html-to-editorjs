package scheme

import (
	"github.com/PuerkitoBio/goquery"
)

type BlockHandler func(selection *goquery.Selection) *Block

type BlockData map[string]interface{}

type Response struct {
	Time    int64   `json:"time"`
	Blocks  []Block `json:"blocks"`
	Version string  `json:"version"`
}

type Block struct {
	Id   *string     `json:"id,omitempty"`
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type ListicleImage struct {
	Title   string `json:"title"`
	Alt     string `json:"alt"`
	Caption string `json:"caption"`
	Src     string `json:"src"`
	Credit  string `json:"credit"`
	Height  string `json:"height"`
	Width   string `json:"width"`
	ID      string `json:"id"`
}
