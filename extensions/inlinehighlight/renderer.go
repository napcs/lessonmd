package inlinehighlight

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// InlineHighlightHTMLRenderer is a renderer.NodeRenderer implementation that
// renders InlineHighlight nodes.
type InlineHighlightHTMLRenderer struct {
	html.Config
}

// NewInlineHighlightHTMLRenderer returns a new InlineHighlightHTMLRenderer.
func NewInlineHighlightHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &InlineHighlightHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *InlineHighlightHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindInlineHighlight, r.renderInlineHighlight)
}

func (r *InlineHighlightHTMLRenderer) renderInlineHighlight(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<mark class=\"inline-highlight\">")
	} else {
		_, _ = w.WriteString("</mark>")
	}
	return ast.WalkContinue, nil
}
