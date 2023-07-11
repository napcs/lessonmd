package lessonmd

import (
	"bytes"
	"lessonmd/extensions/codeblocks"
	"lessonmd/extensions/commandblocks"
	"lessonmd/extensions/details"
	"lessonmd/extensions/inlinehighlight"
	"lessonmd/extensions/notices"
	"lessonmd/extensions/outputblocks"
	"strings"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/mermaid"
)

// AppVersion is the version of the app itself
var AppVersion = "0.0.3"

// ConverterOptions specifies options for converting.
// wrap: wrap the results with a div
// wrapClass: class to give the outer wrapper div. Defaults to "item"
type ConverterOptions struct {
	Wrap               bool
	WrapperClass       string
	AddStyleTag        bool
	AddHighlightJS     bool
	UseSVGforMermaid   bool
	AddMermaidJS       bool
	IncludeFrontmatter bool
}

type converter struct{}

// Converter converts markdown to HTML
var Converter = &converter{}

// Run does the conversion, using ConverterOptions. Takes a byte slice (usually from a reader) and returns a string.
func (c *converter) Run(markdown []byte, o ConverterOptions) (string, error) {

	mmRenderMode := mermaid.RenderModeClient

	if o.UseSVGforMermaid {
		mmRenderMode = mermaid.RenderModeServer
	}

	var extensions = []goldmark.Extender{
		extension.GFM, // builtin
		&mermaid.Extender{NoScript: true, RenderMode: mmRenderMode}, // imported
		outputblocks.OutputExtender,                                 // custom -> outputblocks.go
		inlinehighlight.InlineHighlighter,                           // custom -> inlinehighlight.go
		commandblocks.CommandExtender,                               // custom -> commandblokcs.go
		notices.AdmonitionExtender,
		details.DetailsExtender,
		codeblocks.CodeblockExtender,
	}

	if o.IncludeFrontmatter {
		extensions = append(extensions, meta.New(meta.WithTable()))
	} else {
		extensions = append(extensions, meta.Meta)
	}

	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithAttribute(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // allow raw html
		),
		goldmark.WithExtensions(extensions...),
	)

	var html bytes.Buffer
	// Convert Markdown to HTML
	err := md.Convert(markdown, &html)

	if err != nil {
		return "", err
	}

	out := html.String()

	// add a style tag with css code at the top if reqeusted (default is no)
	if o.AddStyleTag {
		out = c.addCSS(o.WrapperClass) + out + "\n"

	}

	// add the wrapper class if requested (default is yes)
	if o.Wrap {
		out = "<div class=\"" + o.WrapperClass + "\">\n" + out + "\n</div>"
	}

	// add the highlight.js and clipboard code snippet at the bottom if requested (default is no)
	if o.AddHighlightJS {
		out = out + c.addHighlightJS(o.WrapperClass)
	}

	// add the mermaid.js code snippet at the bottom if requested (default is no)
	if o.AddMermaidJS {
		out = out + c.addMermaidJS()
	}

	// Print HTML to standard output
	return out, nil
}

func (c *converter) addCSS(class string) string {
	return "<style>" + c.GenerateCSS(class) + "</style>\n"
}

// GenerateCSS returns a string with the basic stylesheet.
func (c *converter) GenerateCSS(class string) string {
	style := `
.item {
  color: #333;
  direction: ltr;
  font-family: -apple-system,BlinkMacSystemFont,Segoe UI,Helvetica,Arial,sans-serif,Apple Color Emoji,Segoe UI Emoji;
  font-size: 16px;
  line-height: 1.5;
  text-align: left;
  word-wrap: break-word
}

.item a { color: #0969da; }

.item code,.item pre{
  font-family:Monaco, Andale Mono, Courier New, monospace;
  font-size:12px;
  -webkit-border-radius:3px;
  -moz-border-radius:3px;
  border-radius:3px;
}

.item code{
  background-color:#fafafa;
  border: 1px solid #e1e1e8;
  border-radius: 3px;
  font-weight:bolder;
  white-space: nowrap;
}

.item pre>code {
  color: #000;
  display:block;
  font-size:12px;
  line-height:18px;
  margin:0 0 18px;
  padding:8.5px;
  white-space:pre;
  white-space:pre-wrap;
  word-wrap:break-word;
}

.item blockquote, .item details,.item dl,.item ol,.item p,.item pre,.item table, .item ul {
  margin-top: 0;
  margin-bottom: 16px
}

.item blockquote{ padding: 0 1em; color: #6a737d; border-left: .25em solid #dfe2e5 }
.item blockquote>:first-child { margin-top: 0 }
.item blockquote>:last-child { margin-top: 0 }

.item h1, .item h2, .item h3, .item h4, .item h5, .item h6 {
  margin-top: 24px;
  margin-bottom: 16px;
  font-weight: 600;
  line-height: 1.25
}

.item h1 { font-size: 2em }
.item h1, .item h2 { padding-bottom: .3em; border-bottom: 1px solid #eaecef }
.item h3 { font-size: 1.25em }
.item h4 { font-size: 1em }
.item h5 { font-size: .875em }
.item h6 { font-size: .85em; color: #6a737d }
.item ol,.item ul { padding-left: 2em }

.item ol ol, .item ol ul, .item ul ol, .item ul ul { margin-top: 0; margin-bottom: 0 }

.item li { word-wrap: break-all }
.item li>p { margin-top: 16px }
.item li+li { margin-top: .25em }

.item dl { padding: 0 }
.item dl dt { padding: 0; margin-top: 16px; font-size: 1em; font-style: italic; font-weight: 600 }
.item dl dd { padding: 0 16px; margin-bottom: 16px }

.item table{width:100%;margin-bottom:18px;padding:0;border-collapse:separate;*border-collapse:collapse;font-size:13px;border:1px solid #ddd;-webkit-border-radius:4px;-moz-border-radius:4px;border-radius:4px;}table th,table td{padding:10px 10px 9px;line-height:18px;text-align:left;}
.item table th{padding-top:9px;font-weight:bold;vertical-align:middle;border-bottom:1px solid #ddd;color:ddd;background-color:#333;}
.item table td{vertical-align:top;}
.item table th+th, .item table td+td{border-left:1px solid #ddd;}
.item table tr+tr td{border-top:1px solid #ddd;}
.item table tbody tr:first-child td:first-child{-webkit-border-radius:4px 0 0 0;-moz-border-radius:4px 0 0 0;border-radius:4px 0 0 0;}
.item table tbody tr:first-child td:last-child{-webkit-border-radius:0 4px 0 0;-moz-border-radius:0 4px 0 0;border-radius:0 4px 0 0;}
.item table tbody tr:last-child td:first-child{-webkit-border-radius:0 0 0 4px;-moz-border-radius:0 0 0 4px;border-radius:0 0 0 4px;}
.item table tbody tr:last-child td:last-child{-webkit-border-radius:0 0 4px 0;-moz-border-radius:0 0 4px 0;border-radius:0 0 4px 0;}
.item table tr:nth-child(even) { background-color: #e5e5e5; }

.item img { max-width: 100%; box-sizing: initial; background-color: #fff }
.item strong { font-weight: bolder }

.item .hljs-copy {
  float: right;
  cursor: pointer;
}

.item code.command::before {
  content: "$ ";
  font-weight: bolder;
}

.item .output { background-color: #ddd; }
.item .output p { margin: 0 0 0 4px}

.item mark {
  background-color: #fff8c5;
  color: #24292f;
}

.item .notice {
  border-style: solid;
  border-width: 0 0 0 5px;
  box-shadow: 0 1px 2px 0 #ddd;
  color: #193c47;
  margin-bottom: 1em;
  padding: 0.5rem;
}

.item .notice .notice-heading {
  font-size: 0.9em;
  font-weight: bolder;
  text-transform: uppercase;
  margin-bottom: 1em;
}

.item .notice.note {
  background-color: rgb(253, 253, 254);
  border-color: rgb(212, 213, 216);
  color: rgb(71, 71, 72)
}
.item .notice.note a{color: rgb(71, 71, 72); text-decoration-color: rgb(212, 213, 216)}

.item .notice.tip {
  background-color: rgb(230, 246, 230);
  border-color: rgb(0, 148, 0);
  color: rgb(0, 49, 0);
}
.item .notice.tip code {background-color: rgba(0, 164, 0, 0.15)}
.item .notice.tip a {color: rgb(0, 49, 0); text-decoration-color: #009400}

.item .notice.info {
  background-color: rgb(238, 249, 253);
  border-color: rgb(76, 179, 212);
  color: rgb(25, 60, 71)
}
.item .notice.info code {background-color: rgba(84, 199, 236, 0.15)}
.item .notice.info a {color: rgb(25, 60, 71); text-decoration-color: #4CB3D4}

.item .notice.caution {
  background-color: #fff8e6;
  border-color: #e6a700;
  color: rgb(77, 56, 0)
}
.item .notice.caution code {background-color: rgba(255, 186, 0, 0.15)}
.item .notice.caution a {color: rgb(77, 56, 0); text-decoration-color: #e6a700}

.item .notice.warning {
  background-color: #FFEBEC;
  border-color: #E13238;
  color: #4b1113;
}
.item .notice.warning code {background-color: rgba(250, 56, 62, 0.15)}
.item .notice.warning a {  color: #4b1113; text-decoration-color: #4b1113}

.item details {
  background-color: rgb(238, 249, 253);
  border: 1px solid rgb(76, 179, 212);
  border-radius: 10px;
  box-shadow: 0 1px 2px 0 #ddd;
  color: #193c47;
  margin-bottom: 1em;
}

.item details summary ~ * {
  margin-top: 10px; /* Adjust the desired space (in pixels) */
}

.item details .details-content {
  border-top: 1px solid rgb(76, 179, 212);
  padding-top: 10px;
}

.item details summary {
  border-radius: 10px;
  color: rgb(25, 60, 71);
  cursor: pointer;
  list-style: none;
  padding: 10px;
}

.item details summary::marker {display: none; }
.item details summary::-webkit-details-marker {display: none; }
.item details summary:before { content: "\25BA"; margin-right: 5px; } /* Unicode escape sequence for ► */
.item details[open] summary:before { content: "\25BC"; } /* Unicode escape sequence for ▼ */
`
	style = strings.ReplaceAll(style, ".item", "."+class)
	return style
}

func (c *converter) addMermaidJS() string {
	return "<script>" + c.GenerateMermaidJS() + "</script>\n"
}

func (c *converter) GenerateMermaidJS() string {

	return `
function loadMermaid() {
  const m = document.createElement('script');
  m.src = 'https://cdn.jsdelivr.net/npm/mermaid/dist/mermaid.min.js';
  m.async = false;
  m.addEventListener('load', function() {
    try {
      mermaid.initialize({startOnLoad: true});
    } catch (error) {
      console.error(error);
    }
  });
  document.body.appendChild(m);
}

loadMermaid();
`

}
func (c *converter) addHighlightJS(class string) string {
	return "<script>" + c.GenerateHighlightJS(class) + "</script>\n"
}

func (c *converter) GenerateHighlightJS(class string) string {

	out := `
async function loadHighlightJS() {
  await new Promise((resolve, reject) => {
    const highlightScript = document.createElement("script");
    highlightScript.src = 'https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/highlight.min.js';
    highlightScript.onload = resolve;
    highlightScript.onerror = reject;
    document.body.appendChild(highlightScript);
  });

  await new Promise((resolve, reject) => {
    const golangScript = document.createElement("script");
    golangScript.src = 'https:////cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/languages/go.min.js';
    golangScript.onload = resolve;
    golangScript.onerror = reject;
    document.body.appendChild(golangScript);
  });

  await new Promise((resolve, reject) => {
    const css = document.createElement('link')
    css.setAttribute('rel', 'stylesheet');
    css.setAttribute('href', 'https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/styles/default.min.css');
    document.body.appendChild(css);
    css.onload = resolve;
    css.onerror = reject;
  });

  try {
    document.querySelectorAll('.item pre code').forEach(el => {
      hljs.highlightElement(el);
    })
    addButtons();
  } catch (error) {
    console.error(error);
  }
};

function addButtons() {
  var snippets = document.getElementsByClassName('hljs');
  var numberOfSnippets = snippets.length;
  for (var i = 0; i < numberOfSnippets; i++) {
    var p = snippets[i].parentElement;
    var b = document.createElement("button");
    b.classList.add('hljs-copy')
    b.innerText="Copy";

    b.addEventListener("click", function () {
      this.innerText = 'Copying..';
      code = this.nextSibling.innerText;
      navigator.clipboard.writeText(code);
      this.innerText = 'Copied!';
      var that = this;
      setTimeout(function () {
        that.innerText = 'Copy';
      }, 1000)
    });
    p.prepend(b)
  }
}

loadHighlightJS();
`
	out = strings.ReplaceAll(out, ".item", "."+class)
	return out

}
