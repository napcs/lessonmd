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
	line, segment := reader.PeekLine()
	
	// If we encounter another === line, this tab is done
	if bytes.HasPrefix(line, []byte("=== \"")) {
		return parser.Close
	}
	
	// If we have an empty line, look ahead to check what follows
	if len(bytes.TrimSpace(line)) == 0 {
		// Create a temporary reader to look ahead without affecting the current position
		source := reader.Source()
		pos := segment.Start + segment.Len()
		
		// Skip additional blank lines
		for pos < len(source) {
			lineStart := pos
			// Find end of line
			for pos < len(source) && source[pos] != '\n' {
				pos++
			}
			if pos < len(source) {
				pos++ // Skip the newline
			}
			
			// Check if this line has content
			lineContent := bytes.TrimSpace(source[lineStart:pos-1])
			if len(lineContent) == 0 {
				continue // Skip blank lines
			}
			
			// If the next non-empty line is not a tab line, close this tab group
			if !bytes.HasPrefix(lineContent, []byte("=== \"")) {
				return parser.Close
			}
			
			// If it is a tab line, we can continue
			break
		}
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