package inlinehighlight

import "github.com/yuin/goldmark/ast"

type InlineHighlight struct {
	ast.BaseInline
}

// Dump implements Node.Dump.
func (n *InlineHighlight) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

// KindInlineHighlight is a NodeKind of the InlineHighlight node.
var KindInlineHighlight = ast.NewNodeKind("InlineHighlight")

// Kind implements Node.Kind.
func (n *InlineHighlight) Kind() ast.NodeKind {
	return KindInlineHighlight
}

// NewInlineHighlight returns a new InlineHighlight node.
func NewInlineHighlight() *InlineHighlight {
	return &InlineHighlight{}
}
