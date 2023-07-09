package notices

import (
	h "html"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type admonitionHTMLRenderer struct {
}

func NewAdmonitionHTMLRenderer() renderer.NodeRenderer {
	return &admonitionHTMLRenderer{}
}

func (r *admonitionHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(AdmonitionKind, r.renderAdmonition)
}

func (r *admonitionHTMLRenderer) renderAdmonition(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*Admonition)
	if entering {
		_, _ = w.WriteString("<div class=\"notice " + h.EscapeString(n.AdmonitionType) + "\">\n")
		_, _ = w.WriteString("  <div class=\"notice-heading\">" + h.EscapeString(n.Title) + "</div>\n")
		_, _ = w.WriteString("  <div class=\"notice-body\">\n")
	} else {
		_, _ = w.WriteString("  </div>\n")
		_, _ = w.WriteString("</div>\n")
	}
	return ast.WalkContinue, nil
}
