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
	
	// Check if this is an empty line
	trimmedLine := bytes.TrimSpace(line)
	if len(trimmedLine) == 0 {
		// Empty lines are allowed within tabs, continue parsing
		return parser.Continue | parser.HasChildren
	}
	
	// Check indentation level - content must be indented by at least 2 spaces to stay in tab
	indentLevel := getIndentationLevel(line)
	if indentLevel < 2 {
		// Content has returned to root level (less than 2 spaces), close this tab
		return parser.Close
	}
	
	// Remove the first 2 spaces from the line to normalize indentation
	// This is done by advancing the segment start position
	spacesToRemove := 0
	for i := 0; i < len(line) && i < 2 && line[i] == ' '; i++ {
		spacesToRemove++
	}
	
	if spacesToRemove > 0 {
		// Adjust the reader to skip the leading spaces
		reader.Advance(spacesToRemove)
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


// getIndentationLevel returns the number of leading spaces in a line
func getIndentationLevel(line []byte) int {
	count := 0
	for i, b := range line {
		if b == ' ' {
			count++
		} else if b == '\t' {
			// Count tabs as 4 spaces for consistency
			count += 4
		} else {
			break
		}
		// Don't count beyond the line length
		if i >= len(line)-1 {
			break
		}
	}
	return count
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