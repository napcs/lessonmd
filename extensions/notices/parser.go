package notices

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type admonitionParser struct {
}

var defaultAdmonitionParser = &admonitionParser{}

func NewAdmonitionParser() parser.BlockParser {
	return defaultAdmonitionParser
}

func (a *admonitionParser) Trigger() []byte {
	return []byte(":::")
}

func (a *admonitionParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, _ := reader.PeekLine()
	line = bytes.Trim(line, ": \t")

	splitLine := bytes.SplitN(line, []byte(" "), 2)
	if len(splitLine) < 2 {
		return nil, parser.NoChildren
	}

	noticeType := splitLine[0]
	title := string(bytes.TrimSpace(splitLine[1]))
	switch string(noticeType) {
	case "note":
		return NewAdmonition("note", string(title)), parser.NoChildren
	case "tip":
		return NewAdmonition("tip", string(title)), parser.NoChildren
	case "info":
		return NewAdmonition("info", string(title)), parser.NoChildren
	case "caution":
		return NewAdmonition("caution", string(title)), parser.NoChildren
	case "warning":
		return NewAdmonition("warning", string(title)), parser.NoChildren
	default:
		return nil, parser.NoChildren
	}
}

func (a *admonitionParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, segment := reader.PeekLine()

	// If line is ":::", it is the end of the block
	if bytes.Equal(bytes.TrimSpace(line), []byte(":::")) {
		reader.Advance(segment.Len()) // Consume the line
		return parser.Close
	}

	return parser.Continue | parser.HasChildren
}

func (a *admonitionParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
}

func (a *admonitionParser) CanInterruptParagraph() bool {
	return true
}

func (a *admonitionParser) CanAcceptIndentedLine() bool {
	return false
}
