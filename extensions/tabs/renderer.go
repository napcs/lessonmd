package tabs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// TabGroupHTMLRenderer renders TabGroup nodes to HTML
type TabGroupHTMLRenderer struct {
	html.Config
	tabGroupCounter int
}

// NewTabGroupHTMLRenderer returns a new TabGroupHTMLRenderer
func NewTabGroupHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &TabGroupHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// normalizeTabName converts a tab title to a normalized data attribute value
func normalizeTabName(title []byte) string {
	name := strings.ToLower(string(title))
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "_", "-")
	// Remove any characters that aren't alphanumeric or hyphens
	var result strings.Builder
	for _, r := range name {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// RegisterFuncs registers the rendering functions
func (r *TabGroupHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindTabGroup, r.renderTabGroup)
	reg.Register(KindTab, r.renderTab)
}

// renderTabGroup renders a TabGroup node
func (r *TabGroupHTMLRenderer) renderTabGroup(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		r.tabGroupCounter++
		tabGroupID := "tabs-" + strconv.Itoa(r.tabGroupCounter)
		
		_, _ = w.WriteString("<div class=\"tabs\" id=\"")
		_, _ = w.WriteString(tabGroupID)
		_, _ = w.WriteString("\">\n")
		
		
		// Render tab navigation
		_, _ = w.WriteString("  <div class=\"tabs-nav\" role=\"tablist\">\n")
		
		tabIndex := 1
		for child := n.FirstChild(); child != nil; child = child.NextSibling() {
			if tab, ok := child.(*Tab); ok {
				tabID := fmt.Sprintf("tab-%d-%d", r.tabGroupCounter, tabIndex)
				panelID := fmt.Sprintf("tab-panel-%d-%d", r.tabGroupCounter, tabIndex)
				activeClass := ""
				ariaSelected := "false"
				
				if tabIndex == 1 {
					activeClass = " active"
					ariaSelected = "true"
				}
				
				_, _ = w.WriteString("    <button class=\"tab-button")
				_, _ = w.WriteString(activeClass)
				_, _ = w.WriteString("\" role=\"tab\" aria-selected=\"")
				_, _ = w.WriteString(ariaSelected)
				_, _ = w.WriteString("\" aria-controls=\"")
				_, _ = w.WriteString(panelID)
				_, _ = w.WriteString("\" id=\"")
				_, _ = w.WriteString(tabID)
				_, _ = w.WriteString("\" data-tab-name=\"")
				_, _ = w.WriteString(normalizeTabName(tab.Title))
				_, _ = w.WriteString("\">")
				_, _ = w.Write(util.EscapeHTML(tab.Title))
				_, _ = w.WriteString("</button>\n")
				
				tabIndex++
			}
		}
		
		_, _ = w.WriteString("  </div>\n")
		_, _ = w.WriteString("  <div class=\"tab-panels\">\n")
	} else {
		_, _ = w.WriteString("  </div>\n")
		_, _ = w.WriteString("</div>\n")
	}
	
	return ast.WalkContinue, nil
}

// renderTab renders a Tab node
func (r *TabGroupHTMLRenderer) renderTab(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		// Check if this is the first tab in a sequence of consecutive tabs
		if r.isFirstTabInSequence(n) {
			// Start tab group HTML
			r.tabGroupCounter++
			r.renderTabGroupStart(w, n)
		}
		
		// Render individual tab panel
		r.renderTabPanel(w, n, entering)
	} else {
		r.renderTabPanel(w, n, entering)
		
		// Check if this is the last tab in sequence
		if r.isLastTabInSequence(n) {
			// End tab group HTML
			r.renderTabGroupEnd(w)
		}
	}
	
	return ast.WalkContinue, nil
}

// isFirstTabInSequence checks if this tab is the first in a sequence of consecutive tabs
func (r *TabGroupHTMLRenderer) isFirstTabInSequence(n ast.Node) bool {
	// Check if previous sibling is also a Tab
	if prev := n.PreviousSibling(); prev != nil {
		if _, ok := prev.(*Tab); ok {
			return false // Not first, previous is also a tab
		}
	}
	return true // Either no previous sibling or previous is not a tab
}

// isLastTabInSequence checks if this tab is the last in a sequence of consecutive tabs
func (r *TabGroupHTMLRenderer) isLastTabInSequence(n ast.Node) bool {
	// Check if next sibling is also a Tab
	if next := n.NextSibling(); next != nil {
		if _, ok := next.(*Tab); ok {
			return false // Not last, next is also a tab
		}
	}
	return true // Either no next sibling or next is not a tab
}

// renderTabGroupStart renders the beginning of a tab group
func (r *TabGroupHTMLRenderer) renderTabGroupStart(w util.BufWriter, firstTab ast.Node) {
	tabGroupID := "tabs-" + strconv.Itoa(r.tabGroupCounter)
	
	_, _ = w.WriteString("<div class=\"tabs\" id=\"")
	_, _ = w.WriteString(tabGroupID)
	_, _ = w.WriteString("\">\n")
	
	// Render tab navigation by collecting all consecutive tabs
	_, _ = w.WriteString("  <div class=\"tabs-nav\" role=\"tablist\">\n")
	
	tabIndex := 1
	for tab := firstTab; tab != nil; tab = tab.NextSibling() {
		if tabNode, ok := tab.(*Tab); ok {
			tabID := fmt.Sprintf("tab-%d-%d", r.tabGroupCounter, tabIndex)
			panelID := fmt.Sprintf("tab-panel-%d-%d", r.tabGroupCounter, tabIndex)
			activeClass := ""
			ariaSelected := "false"
			
			if tabIndex == 1 {
				activeClass = " active"
				ariaSelected = "true"
			}
			
			_, _ = w.WriteString("    <button class=\"tab-button")
			_, _ = w.WriteString(activeClass)
			_, _ = w.WriteString("\" role=\"tab\" aria-selected=\"")
			_, _ = w.WriteString(ariaSelected)
			_, _ = w.WriteString("\" aria-controls=\"")
			_, _ = w.WriteString(panelID)
			_, _ = w.WriteString("\" id=\"")
			_, _ = w.WriteString(tabID)
			_, _ = w.WriteString("\" data-tab-name=\"")
			_, _ = w.WriteString(normalizeTabName(tabNode.Title))
			_, _ = w.WriteString("\">")
			_, _ = w.Write(util.EscapeHTML(tabNode.Title))
			_, _ = w.WriteString("</button>\n")
			
			tabIndex++
		} else {
			break // Stop when we hit a non-tab
		}
	}
	
	_, _ = w.WriteString("  </div>\n")
	_, _ = w.WriteString("  <div class=\"tab-panels\">\n")
}

// renderTabGroupEnd renders the end of a tab group
func (r *TabGroupHTMLRenderer) renderTabGroupEnd(w util.BufWriter) {
	_, _ = w.WriteString("  </div>\n")
	_, _ = w.WriteString("</div>\n")
}

// renderTabPanel renders an individual tab panel
func (r *TabGroupHTMLRenderer) renderTabPanel(w util.BufWriter, n ast.Node, entering bool) {
	if entering {
		// Find tab index by counting tabs from first tab in sequence
		firstTab := n
		for prev := n.PreviousSibling(); prev != nil; prev = prev.PreviousSibling() {
			if _, ok := prev.(*Tab); ok {
				firstTab = prev
			} else {
				break
			}
		}
		
		tabIndex := 1
		for tab := firstTab; tab != nil && tab != n; tab = tab.NextSibling() {
			if _, ok := tab.(*Tab); ok {
				tabIndex++
			}
		}
		
		tabID := fmt.Sprintf("tab-%d-%d", r.tabGroupCounter, tabIndex)
		panelID := fmt.Sprintf("tab-panel-%d-%d", r.tabGroupCounter, tabIndex)
		activeClass := ""
		
		if tabIndex == 1 {
			activeClass = " active"
		}
		
		tab := n.(*Tab)
		_, _ = w.WriteString("    <div class=\"tab-panel")
		_, _ = w.WriteString(activeClass)
		_, _ = w.WriteString("\" role=\"tabpanel\" aria-labelledby=\"")
		_, _ = w.WriteString(tabID)
		_, _ = w.WriteString("\" id=\"")
		_, _ = w.WriteString(panelID)
		_, _ = w.WriteString("\" data-tab-name=\"")
		_, _ = w.WriteString(normalizeTabName(tab.Title))
		_, _ = w.WriteString("\">\n")
	} else {
		_, _ = w.WriteString("    </div>\n")
	}
}