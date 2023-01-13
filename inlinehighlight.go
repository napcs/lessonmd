package lessonmd

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

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

type inlineHighlightDelimiterProcessor struct {
}

func (p *inlineHighlightDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == '='
}

func (p *inlineHighlightDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

func (p *inlineHighlightDelimiterProcessor) OnMatch(consumes int) ast.Node {
	return NewInlineHighlight()
}

var defaultInlineHighlightDelimiterProcessor = &inlineHighlightDelimiterProcessor{}

type inlineHighlightParser struct {
}

var defaultInlineHighlightParser = &inlineHighlightParser{}

// NewInlineHighlightParser return a new InlineParser that parses inline highlight expressions.
func NewInlineHighlightParser() parser.InlineParser {
	return defaultInlineHighlightParser
}

func (s *inlineHighlightParser) Trigger() []byte {
	return []byte{'='}
}

func (s *inlineHighlightParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := parser.ScanDelimiter(line, before, 2, defaultInlineHighlightDelimiterProcessor)
	if node == nil {
		return nil
	}
	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}

func (s *inlineHighlightParser) CloseBlock(parent ast.Node, pc parser.Context) {
	// nothing to do
}

// InlineHighlightHTMLRenderer is a renderer.NodeRenderer implementation that
// renders InlineHighlight nodes.
type InlineHighlightHTMLRenderer struct {
	html.Config
}

// NewInlineHighlightHTMLRenderer returns a new InlineHighlightHTMLRenderer.
func NewInlineHighlightHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &InlineHighlightHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *InlineHighlightHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindInlineHighlight, r.renderInlineHighlight)
}

func (r *InlineHighlightHTMLRenderer) renderInlineHighlight(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<mark class=\"inline-highlight\">")
	} else {
		_, _ = w.WriteString("</mark>")
	}
	return ast.WalkContinue, nil
}

type inlineHighlighter struct {
}

// InlineHighlighter is an extension that allows you to use expressions like '<^>text<^>' to highlight pieces of text.
var InlineHighlighter = &inlineHighlighter{}

func (e *inlineHighlighter) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewInlineHighlightParser(), 900),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewInlineHighlightHTMLRenderer(), 0),
	))
}
