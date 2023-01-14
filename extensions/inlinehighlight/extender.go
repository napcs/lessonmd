package inlinehighlight

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type inlineHighlighter struct {
}

// InlineHighlighter is an extension that allows you to use expressions like '<^>text<^>' to highlight pieces of text.
var InlineHighlighter = &inlineHighlighter{}

func (e *inlineHighlighter) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewInlineHighlightParser(), 900),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewInlineHighlightHTMLRenderer(), 0),
	))
}
