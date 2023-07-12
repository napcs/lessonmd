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
	inputString := `:::note Notice of wrongdoing
This is a test note.

It supports code.
`

	inputString += "```js\nlet x = 25;\n```"
	inputString += `
This is the output.
`

	input := []byte(inputString)
	expected := `<div class="notice note">
  <div class="notice-heading">Notice of wrongdoing</div>
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

func TestDetails(t *testing.T) {
	str := `[details What is the best Markdown tool?
lessonmd is the best.

You need to use it.
]`
	input := []byte(str)
	expected := `<details><summary>What is the best Markdown tool?</summary>
<div class="details-content">
<p>lessonmd is the best.</p>
<p>You need to use it.</p>
</div>
</details>
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

func TestDetailsOpen(t *testing.T) {
	str := `[details open What is the best Markdown tool?
lessonmd is the best.

You need to use it.
]`
	input := []byte(str)
	expected := `<details open><summary>What is the best Markdown tool?</summary>
<div class="details-content">
<p>lessonmd is the best.</p>
<p>You need to use it.</p>
</div>
</details>
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
