package notices

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type admonitionExtender struct{}

var AdmonitionExtender = &admonitionExtender{}

func (e *admonitionExtender) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithBlockParsers(
		util.Prioritized(&admonitionParser{}, 100),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&admonitionHTMLRenderer{}, 0),
	))
}
