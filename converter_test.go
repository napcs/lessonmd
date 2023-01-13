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
console.log(&quot;The answer is &quot; + x&quot;);
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
