package commandblocks

import "github.com/yuin/goldmark/ast"

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
