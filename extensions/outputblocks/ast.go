package outputblocks

import "github.com/yuin/goldmark/ast"

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
