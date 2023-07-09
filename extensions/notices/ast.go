package notices

import (
	"github.com/yuin/goldmark/ast"
)

var AdmonitionKind = ast.NewNodeKind("Admonition")

type Admonition struct {
	ast.BaseBlock
	AdmonitionType string
	Title          string
}

func NewAdmonition(typ, title string) *Admonition {
	return &Admonition{
		AdmonitionType: typ,
		Title:          title,
		BaseBlock:      ast.BaseBlock{},
	}
}

func (a *Admonition) Kind() ast.NodeKind {
	return AdmonitionKind
}

func (a *Admonition) Dump(source []byte, level int) {
	ast.DumpHelper(a, source, level, map[string]string{
		"AdmonitionType": a.AdmonitionType,
		"Title":          a.Title,
	}, nil)
}
