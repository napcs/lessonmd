package codeblocks

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type codeblockExtender struct{}

var CodeblockExtender = &codeblockExtender{}

func (e *codeblockExtender) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(&CodeblockTransformer{}, 100),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&CodeblockHTMLRenderer{}, 90),
	))
}
