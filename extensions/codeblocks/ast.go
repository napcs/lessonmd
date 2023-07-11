package codeblocks

import "github.com/yuin/goldmark/ast"

//-----ast

// Kind is CodeBlock
var CodeKind = ast.NewNodeKind("CodeBlock")

// Its raw contents are the plain text of the code
type CodeBlock struct {
	ast.BaseBlock
	Filename string
	Language string
}

// IsRaw reports that this block should be rendered as-is.
func (*CodeBlock) IsRaw() bool { return true }

// Kind reports that this is a CodeBlock.
func (*CodeBlock) Kind() ast.NodeKind { return CodeKind }

// Dump dumps the contents of this block to stdout.
func (b *CodeBlock) Dump(src []byte, level int) {
	ast.DumpHelper(b, src, level, nil, nil)
}
