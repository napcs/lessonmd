package tabs

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)


type tabsParser struct{}

var defaultTabsParser = &tabsParser{}

// NewTabsParser returns a new BlockParser for tabs
func NewTabsParser() parser.BlockParser {
	return defaultTabsParser
}

// Trigger returns the characters that trigger this parser
func (p *tabsParser) Trigger() []byte {
	return []byte{'='}
}

// Open parses the beginning of a tab block
func (p *tabsParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, segment := reader.PeekLine()
	
	// Check for === "Title" pattern
	if !bytes.HasPrefix(line, []byte("=== \"")) {
		return nil, parser.NoChildren
	}
	
	// Extract title between quotes
	title := extractTabTitle(line)
	
	// Create new Tab node - renderer will handle grouping
	tab := &Tab{Title: title}
	
	reader.Advance(segment.Len())
	return tab, parser.HasChildren
}

// Continue continues parsing a tab block
func (p *tabsParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, _ := reader.PeekLine()
	
	// If we encounter another === line, this tab is done
	if bytes.HasPrefix(line, []byte("=== \"")) {
		return parser.Close
	}
	
	return parser.Continue | parser.HasChildren
}

// Close closes the tab block
func (p *tabsParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	// Nothing special needed for closing
}

// CanInterruptParagraph returns true if this parser can interrupt paragraphs
func (p *tabsParser) CanInterruptParagraph() bool {
	return true
}

// CanAcceptIndentedLine returns true if this parser can accept indented lines
func (p *tabsParser) CanAcceptIndentedLine() bool {
	return false
}


// extractTabTitle extracts the title from a === "Title" line
func extractTabTitle(line []byte) []byte {
	// Find content between === " and closing "
	start := bytes.Index(line, []byte("\""))
	if start == -1 {
		return []byte("Tab")
	}
	start++ // Skip opening quote
	
	end := bytes.LastIndex(line, []byte("\""))
	if end <= start {
		return []byte("Tab")
	}
	
	return util.TrimRightSpace(line[start:end])
}