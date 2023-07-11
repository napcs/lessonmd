package codeblocks

import (
	"text/template"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// CodeBlockHTMLRenderer renders code blocks.
type CodeblockHTMLRenderer struct{}

func (r *CodeblockHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(CodeKind, r.Render)
}

// Render does the actual rendering.
func (r *CodeblockHTMLRenderer) Render(w util.BufWriter, src []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*CodeBlock)
	if entering {

		if n.Filename != "" {
			w.WriteString("<p>" + n.Filename + "</p>\n")
		}

		w.WriteString("<pre><code class=\"language-" + n.Language + "\">")

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
