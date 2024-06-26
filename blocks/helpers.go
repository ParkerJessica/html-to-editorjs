package blocks

import (
	"strconv"

	"github.com/ParkerJessica/html-to-editorjs/scheme"
	"github.com/PuerkitoBio/goquery"
)

func getImgFile(selection *goquery.Selection) scheme.BlockData {
	file := scheme.BlockData{}
	if attr, _ := selection.Attr("src"); attr != "" {
		file["url"] = attr
	}

	if attr, _ := selection.Attr("height"); attr != "" {
		if i, err := strconv.Atoi(attr); err == nil {
			file["height"] = i
		}
	}

	if attr, _ := selection.Attr("width"); attr != "" {
		if i, err := strconv.Atoi(attr); err == nil {
			file["width"] = i
		}
	}

	return file
}
