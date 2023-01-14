package outputblocks

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

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
