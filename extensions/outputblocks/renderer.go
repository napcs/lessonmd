package outputblocks

import (
	"html/template"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// ----- Render

// OutputHTMLRenderer renders code blocks.
type OutputHTMLRenderer struct{}

func (r *OutputHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(Kind, r.Render)
}

// Render does the actual rendering.
func (r *OutputHTMLRenderer) Render(w util.BufWriter, src []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*OutputBlock)
	if entering {
		w.WriteString("<div class=\"output\">\n")
		w.WriteString("<p>Output</p>\n<pre><code>")
		lines := n.Lines()
		for i := 0; i < lines.Len(); i++ {
			line := lines.At(i)
			template.HTMLEscape(w, line.Value(src))
		}
	} else {
		w.WriteString("</code></pre>\n")
		w.WriteString("</div>")
	}
	return ast.WalkContinue, nil
}
