package details

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type detailsParser struct {
}

var defaultDetailsParser = &detailsParser{}

func NewDetailsParser() parser.BlockParser {
	return defaultDetailsParser
}

func (d *detailsParser) Trigger() []byte {
	return []byte("[details")
}

func (a *detailsParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	open := false
	line, _ := reader.PeekLine()

	if !bytes.HasPrefix(line, []byte("[details")) {
		return nil, parser.NoChildren
	}

	// Get the part after "[details "
	line = bytes.TrimPrefix(line, []byte("[details "))

	// Check if the line contains "open"
	if bytes.HasPrefix(line, []byte("open ")) {
		// Get the part after "open "
		open = true
		line = bytes.TrimPrefix(line, []byte("open "))
	}

	// Remove the trailing new line or white space
	line = bytes.TrimRight(line, " \n")

	// line now contains the rest of the line after "[details " or "[details open ", which can be used as the title.
	title := string(line)

	// Add your logic to handle different types of admonitions if needed.

	// reader.Advance(segment.Len()) // Consume the line

	return NewDetails(string(title), open), parser.NoChildren
}

func (d *detailsParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, segment := reader.PeekLine()

	// If line is "]", it is the end of the block
	if bytes.Equal(bytes.TrimSpace(line), []byte("]")) {
		reader.Advance(segment.Len()) // Consume the line
		return parser.Close
	}

	return parser.Continue | parser.HasChildren
}

func (d *detailsParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
}

func (d detailsParser) CanInterruptParagraph() bool {
	return true
}

func (d *detailsParser) CanAcceptIndentedLine() bool {
	return false
}
