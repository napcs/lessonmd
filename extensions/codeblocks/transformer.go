package codeblocks

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// ----- CodeblockTransformer

// CodeblockTransformer transforms code fences with `output` labels. It just changes the type.
type CodeblockTransformer struct {
}

// Transform converts the nodes.
func (s *CodeblockTransformer) Transform(doc *ast.Document, reader text.Reader, pctx parser.Context) {

	// define the types
	var (
		codeBlocks []*ast.FencedCodeBlock
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

		codeBlocks = append(codeBlocks, cb)
		return ast.WalkContinue, nil
	})

	// Nothing to do if there were no code blocks found.
	if len(codeBlocks) == 0 {
		return
	}

	// replace the old code blocks with the new ones using our type.
	for _, cb := range codeBlocks {
		b := new(CodeBlock)
		b.SetLines(cb.Lines())

		infoLine := cb.Info.Segment.Value(reader.Source())
		parts := bytes.SplitN(infoLine, []byte(" "), 2)
		b.Language = string(parts[0])

		if len(parts) >= 2 {
			b.Filename = string(bytes.TrimSpace(parts[1]))
		} else {
			b.Filename = ""
		}

		parent := cb.Parent()
		if parent != nil {
			parent.ReplaceChild(parent, cb, b)
		}
	}

}
