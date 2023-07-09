package lessonmd

import (
	"os"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	input := []byte("Hello World")
	expected := "<p>Hello World</p>"

	o := ConverterOptions{
		Wrap: false, WrapperClass: "item",
		AddStyleTag:      false,
		AddHighlightJS:   false,
		UseSVGforMermaid: false,
		AddMermaidJS:     false,
	}

	output, _ := Converter.Run(input, o)

	if !strings.Contains(output, expected) {
		t.Errorf("Expected the output to include %q but it was %q", expected, output)
	}
}

func TestAllowRawHTML(t *testing.T) {
	input := []byte("<p>Hello World</p>")
	expected := "<p>Hello World</p>"

	o := ConverterOptions{
		Wrap: false, WrapperClass: "item",
		AddStyleTag:      false,
		AddHighlightJS:   false,
		UseSVGforMermaid: false,
		AddMermaidJS:     false,
	}

	output, _ := Converter.Run(input, o)

	if !strings.Contains(output, expected) {
		t.Errorf("Expected the output to include %q but it was %q", expected, output)
	}
}

func TestAddWrapper(t *testing.T) {
	input := []byte("Hello World")
	expected := "<div class=\"item\">"

	o := ConverterOptions{
		Wrap:             true,
		WrapperClass:     "item",
		AddStyleTag:      false,
		AddHighlightJS:   false,
		UseSVGforMermaid: false,
		AddMermaidJS:     false,
	}

	output, _ := Converter.Run(input, o)

	if !strings.Contains(output, expected) {
		t.Errorf("Expected the output to include %q but it was %q", expected, output)
	}
}

func TestAddWrapperWithCustomClass(t *testing.T) {
	input := []byte("Hello World")
	expected := "<div class=\"lesson-item\">"

	o := ConverterOptions{
		Wrap:             true,
		WrapperClass:     "lesson-item",
		AddStyleTag:      false,
		AddHighlightJS:   false,
		UseSVGforMermaid: false,
		AddMermaidJS:     false,
	}

	output, _ := Converter.Run(input, o)

	if !strings.Contains(output, expected) {
		t.Errorf("Expected the output to include %q but it was %q", expected, output)
	}
}

func TestCustomClassInCSS(t *testing.T) {
	input := []byte("Hello World")
	expected := ".lesson-item h1"

	o := ConverterOptions{
		Wrap:             true,
		WrapperClass:     "lesson-item",
		AddStyleTag:      true,
		AddHighlightJS:   false,
		UseSVGforMermaid: false,
		AddMermaidJS:     false,
	}

	output, _ := Converter.Run(input, o)

	if !strings.Contains(output, expected) {
		t.Errorf("Expected the output to include %q but it was %q", expected, output)
	}
}

func TestCommandBlocks(t *testing.T) {
	input := []byte("```command\nls -alh\n```")
	expected := "<pre><code class=\"language-bash command\">ls -alh\n</code></pre>\n"

	o := ConverterOptions{
		Wrap:             false,
		WrapperClass:     "item",
		AddStyleTag:      false,
		AddHighlightJS:   false,
		UseSVGforMermaid: false,
		AddMermaidJS:     false,
	}

	output, _ := Converter.Run(input, o)

	if !strings.Contains(output, expected) {
		t.Errorf("Expected the output to include %q but it was %q", expected, output)
	}
}

func TestOutputBlocks(t *testing.T) {
	input := []byte("```output\nls -alh\n```")
	expected := "<div class=\"output\">\n<p>Output</p>\n<pre><code>ls -alh\n</code></pre>\n</div>"

	o := ConverterOptions{
		Wrap:             false,
		WrapperClass:     "item",
		AddStyleTag:      false,
		AddHighlightJS:   false,
		UseSVGforMermaid: false,
		AddMermaidJS:     false,
	}

	output, _ := Converter.Run(input, o)

	if !strings.Contains(output, expected) {
		t.Errorf("Expected the output to include %q but it was %q", expected, output)
	}
}

func TestIntegration(t *testing.T) {
	input, _ := os.ReadFile("examples/lesson.md")

	expected := `<div class="item">
<h1 id="lesson-item-title">Lesson item title</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
<pre><code class="language-js">let x = 2;
console.log(&quot;The answer is &quot; + x);
console.log(&quot;It is the best answer.&quot;);
</code></pre>
<p>Run the command to execute <mark class="inline-highlight">foo.js</mark>:</p>
<pre><code class="language-bash command">node foo.js
</code></pre>
<p>You'll see this output:</p>
<div class="output">
<p>Output</p>
<pre><code>The answer is 2
It is the best answer.
</code></pre>
</div><p>Notice that the answer is <code>2</code> in the output.</p>
<p>Here's a diagram:</p>
<div class="mermaid">graph TD;
    A--&gt;B;
    A--&gt;C;
    B--&gt;D;
    C--&gt;D;
</div><p>See <a href="https://example.com">https://example.com</a> for an example site.</p>
<p>Here's a table:</p>
<table>
<thead>
<tr>
<th>First name</th>
<th>Last name</th>
</tr>
</thead>
<tbody>
<tr>
<td>Homer</td>
<td>Simpson</td>
</tr>
<tr>
<td>Marge</td>
<td>Simpson</td>
</tr>
<tr>
<td>Barney</td>
<td>Gumble</td>
</tr>
</tbody>
</table>
<div class="notice note">
  <div class="notice-heading">This is a note</div>
  <div class="notice-body">
<p>Use this to let people know things.</p>
<p>This is <code>syntax</code>.</p>
  </div>
</div>
<div class="notice tip">
  <div class="notice-heading">This is a tip</div>
  <div class="notice-body">
<p>Use this to give people a tip.</p>
<p>This is <code>syntax</code>.</p>
  </div>
</div>
<div class="notice info">
  <div class="notice-heading">This is an info box</div>
  <div class="notice-body">
<p>Use this to give people additional info.</p>
<p>This is <code>syntax</code>.</p>
  </div>
</div>
<div class="notice caution">
  <div class="notice-heading">This is a caution</div>
  <div class="notice-body">
<p>Use this to give cautionary advice.</p>
<p>This is <code>syntax</code>.</p>
  </div>
</div>
<div class="notice warning">
  <div class="notice-heading">This is a warning</div>
  <div class="notice-body">
<p>Use this to warn people of something that may go wrong.</p>
<p>This is <code>syntax</code>.</p>
  </div>
</div>
<p>That's the end.</p>

</div>`

	o := ConverterOptions{
		Wrap:             true,
		WrapperClass:     "item",
		AddStyleTag:      false,
		AddHighlightJS:   false,
		UseSVGforMermaid: false,
		AddMermaidJS:     false,
	}

	output, _ := Converter.Run(input, o)

	if output != expected {
		t.Errorf("Expected the output to include %q but it was %q", expected, output)
	}
}
func TestFrontMatterHide(t *testing.T) {
	input, _ := os.ReadFile("examples/lesson2.md")
	expected := "<p>This is a test.</p>\n"

	o := ConverterOptions{
		Wrap:               false,
		WrapperClass:       "item",
		AddStyleTag:        false,
		AddHighlightJS:     false,
		UseSVGforMermaid:   false,
		AddMermaidJS:       false,
		IncludeFrontmatter: false,
	}

	output, _ := Converter.Run(input, o)

	if output != expected {
		t.Errorf("Expected the output to include %q but it was %q", expected, output)
	}
}

func TestFrontMatterShow(t *testing.T) {
	input, _ := os.ReadFile("examples/lesson2.md")
	expected := `<table>
<thead>
<tr>
<th>title</th>
<th>summary</th>
</tr>
</thead>
<tbody>
<tr>
<td>this is a title</td>
<td>this is a summary</td>
</tr>
</tbody>
</table>
<p>This is a test.</p>
`

	o := ConverterOptions{
		Wrap:               false,
		WrapperClass:       "item",
		AddStyleTag:        false,
		AddHighlightJS:     false,
		UseSVGforMermaid:   false,
		AddMermaidJS:       false,
		IncludeFrontmatter: true,
	}

	output, _ := Converter.Run(input, o)

	if output != expected {
		t.Errorf("Expected the output to include %q but it was %q", expected, output)
	}
}

func TestFences(t *testing.T) {
	input := []byte(":::note Notice\ntest\n:::\n")
	expected := `<div class="notice note">
  <div class="notice-heading">Notice</div>
  <div class="notice-body">
<p>test</p>
  </div>
</div>
`

	o := ConverterOptions{
		Wrap:             false,
		WrapperClass:     "item",
		AddStyleTag:      false,
		AddHighlightJS:   false,
		UseSVGforMermaid: false,
		AddMermaidJS:     false,
	}

	output, _ := Converter.Run(input, o)

	if !strings.Contains(output, expected) {
		t.Errorf("Expected the output to include %q but it was %q", expected, output)
	}
}

func TestFencesWithCodeBlocks(t *testing.T) {
	inputString := `:::note Notice
This is a test note.

It supports code.
`

	inputString += "```js\nlet x = 25;\n```"
	inputString += `
This is the output.
`

	input := []byte(inputString)
	expected := `<div class="notice note">
  <div class="notice-heading">Notice</div>
  <div class="notice-body">
<p>This is a test note.</p>
<p>It supports code.</p>
<pre><code class="language-js">let x = 25;
</code></pre>
<p>This is the output.</p>
  </div>
</div>
`

	o := ConverterOptions{
		Wrap:             false,
		WrapperClass:     "item",
		AddStyleTag:      false,
		AddHighlightJS:   false,
		UseSVGforMermaid: false,
		AddMermaidJS:     false,
	}

	output, _ := Converter.Run(input, o)

	if !strings.Contains(output, expected) {
		t.Errorf("Expected the output to include %q but it was %q", expected, output)
	}
}
