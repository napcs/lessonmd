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

// Kind is CommandBlock
var CommandKind = ast.NewNodeKind("CommandBlock")

// Its raw contents are the plain text of the command
type CommandBlock struct {
	ast.BaseBlock
}

// IsRaw reports that this block should be rendered as-is.
func (*CommandBlock) IsRaw() bool { return true }

// Kind reports that this is a CommandBlock.
func (*CommandBlock) Kind() ast.NodeKind { return CommandKind }

// Dump dumps the contents of this block to stdout.
func (b *CommandBlock) Dump(src []byte, level int) {
	ast.DumpHelper(b, src, level, nil, nil)
}

// ----- CommandTransformer

// CommandTransformer transforms code fences with `output` labels. It just changes the type.
type CommandTransformer struct {
}

// Transform converts the nodes.
func (s *CommandTransformer) Transform(doc *ast.Document, reader text.Reader, pctx parser.Context) {

	// define the types
	var (
		commandBlocks []*ast.FencedCodeBlock // the type of block we're looking for
		_command      = []byte("command")    // the code fence label
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

		// if not a command block, move along.
		if !bytes.Equal(lang, _command) {
			return ast.WalkContinue, nil
		}

		commandBlocks = append(commandBlocks, cb)
		return ast.WalkContinue, nil
	})

	// Nothing to do if there were no command blocks found.
	if len(commandBlocks) == 0 {
		return
	}

	// replace the old code blocks with the new ones using our type.
	for _, cb := range commandBlocks {
		b := new(CommandBlock)
		b.SetLines(cb.Lines())

		parent := cb.Parent()
		if parent != nil {
			parent.ReplaceChild(parent, cb, b)
		}
	}

}

// ----- Render

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

type commandExtender struct{}

var CommandExtender = &commandExtender{}

func (e *commandExtender) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(&CommandTransformer{}, 100),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&CommandHTMLRenderer{}, 0),
	))
}
