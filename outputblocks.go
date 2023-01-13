package lessonmd

import (
	"bytes"
	"text/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

//-----ast

// Kind is OutputBlock
var Kind = ast.NewNodeKind("OutputBlock")

// Its raw contents are the plain text of the output
type OutputBlock struct {
	ast.BaseBlock
}

// IsRaw reports that this block should be rendered as-is.
func (*OutputBlock) IsRaw() bool { return true }

// Kind reports that this is an OutputBlock.
func (*OutputBlock) Kind() ast.NodeKind { return Kind }

// Dump dumps the contents of this block to stdout.
func (b *OutputBlock) Dump(src []byte, level int) {
	ast.DumpHelper(b, src, level, nil, nil)
}

// ----- OutputTransformer

// OutputTransformer transforms code fences with `output` labels. It just changes the type.
type OutputTransformer struct {
}

// Transform converts the nodes.
func (s *OutputTransformer) Transform(doc *ast.Document, reader text.Reader, pctx parser.Context) {

	// define the types
	var (
		outputBlocks []*ast.FencedCodeBlock // the type of block we're looking for
		_output      = []byte("output")     // the code fence label
	)

	// Collect all blocks to be replaced without modifying the tree.
	ast.Walk(doc, func(node ast.Node, enter bool) (ast.WalkStatus, error) {
		if !enter {
			return ast.WalkContinue, nil
		}

		cb, ok := node.(*ast.FencedCodeBlock)
		if !ok {
			return ast.WalkContinue, nil
		}

		// find the language.
		lang := cb.Language(reader.Source())

		// if not an output block, move along.
		if !bytes.Equal(lang, _output) {
			return ast.WalkContinue, nil
		}

		outputBlocks = append(outputBlocks, cb)
		return ast.WalkContinue, nil
	})

	// Nothing to do if there were not output blocks found.
	if len(outputBlocks) == 0 {
		return
	}

	// replace the old code blocks with the new ones using our type.
	for _, cb := range outputBlocks {
		b := new(OutputBlock)
		b.SetLines(cb.Lines())

		parent := cb.Parent()
		if parent != nil {
			parent.ReplaceChild(parent, cb, b)
		}
	}

}

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

type outputExtender struct {
}

var OutputExtender = &outputExtender{}

func (e *outputExtender) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(&OutputTransformer{}, 100),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&OutputHTMLRenderer{}, 0),
	))
}
