package inlinehighlight

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
)

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
