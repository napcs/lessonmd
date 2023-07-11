package details

import (
	h "html"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type detailsHTMLRenderer struct {
}

func NewDetailsHTMLRenderer() renderer.NodeRenderer {
	return &detailsHTMLRenderer{}
}

func (r *detailsHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(DetailsKind, r.renderDetails)
}

func (r *detailsHTMLRenderer) renderDetails(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*Details)
	if entering {
		o := ""
		if n.Open {
			o = " open"
		}
		_, _ = w.WriteString("<details" + o + "><summary>" + h.EscapeString(n.Title) + "</summary>\n")
		_, _ = w.WriteString("<div class=\"details-content\">\n")
	} else {

		_, _ = w.WriteString("</div>\n")
		_, _ = w.WriteString("</details>\n")
	}
	return ast.WalkContinue, nil
}
