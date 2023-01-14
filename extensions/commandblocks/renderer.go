package commandblocks

import (
	"text/template"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// CommandHTMLRenderer renders code blocks.
type CommandHTMLRenderer struct{}

func (r *CommandHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(CommandKind, r.Render)
}

// Render does the actual rendering.
func (r *CommandHTMLRenderer) Render(w util.BufWriter, src []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*CommandBlock)
	if entering {
		w.WriteString("<pre><code class=\"language-bash command\">")
		lines := n.Lines()
		for i := 0; i < lines.Len(); i++ {
			line := lines.At(i)
			template.HTMLEscape(w, line.Value(src))
		}
	} else {
		w.WriteString("</code></pre>\n")
	}
	return ast.WalkContinue, nil
}
