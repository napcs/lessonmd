package inlinehighlight

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

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
