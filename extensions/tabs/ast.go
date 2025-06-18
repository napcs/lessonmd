package tabs

import (
	"github.com/yuin/goldmark/ast"
)

// TabGroup represents the entire tab group container
type TabGroup struct {
	ast.BaseBlock
	tabCounter int // Counter for generating unique IDs
}

// KindTabGroup is the NodeKind for TabGroup
var KindTabGroup = ast.NewNodeKind("TabGroup")

// Kind returns the kind of this node
func (n *TabGroup) Kind() ast.NodeKind {
	return KindTabGroup
}

// Dump dumps the TabGroup node to stdout
func (n *TabGroup) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

// Tab represents a single tab with title and content
type Tab struct {
	ast.BaseBlock
	Title []byte // Tab title from === "Title"
}

// KindTab is the NodeKind for Tab
var KindTab = ast.NewNodeKind("Tab")

// Kind returns the kind of this node
func (n *Tab) Kind() ast.NodeKind {
	return KindTab
}

// Dump dumps the Tab node to stdout
func (n *Tab) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, map[string]string{
		"Title": string(n.Title),
	}, nil)
}