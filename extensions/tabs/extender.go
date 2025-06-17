package tabs

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type tabsExtender struct{}

// TabsExtender is the tabs extension
var TabsExtender = &tabsExtender{}

// Extend extends the Goldmark parser with tabs functionality
func (e *tabsExtender) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithBlockParsers(
		util.Prioritized(NewTabsParser(), 100),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewTabGroupHTMLRenderer(), 0),
	))
}