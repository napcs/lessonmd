package commandblocks

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

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
