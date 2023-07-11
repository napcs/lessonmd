package details

import (
	"strconv"

	"github.com/yuin/goldmark/ast"
)

var DetailsKind = ast.NewNodeKind("Details")

type Details struct {
	ast.BaseBlock
	Title string
	Open  bool
}

func NewDetails(title string, open bool) *Details {
	return &Details{
		Title:     title,
		Open:      open,
		BaseBlock: ast.BaseBlock{},
	}
}

func (d *Details) Kind() ast.NodeKind {
	return DetailsKind
}

func (d *Details) Dump(source []byte, level int) {
	ast.DumpHelper(d, source, level, map[string]string{
		"Title": d.Title,
		"Open":  strconv.FormatBool(d.Open),
	}, nil)
}
