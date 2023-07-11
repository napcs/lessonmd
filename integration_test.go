//go:build integration
// +build integration

package lessonmd

import (
	"fmt"
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {
	input, _ := os.ReadFile("examples/lesson.md")

	expected := `<div class="item">
<h1 id="lesson-item-title">Lesson item title</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
<pre><code class="language-js">let x = 2;
console.log(&#34;The answer is &#34; + x);
console.log(&#34;It is the best answer.&#34;);
</code></pre>
<p>app.js</p>
<pre><code class="language-js">let x = 2;
console.log(&#34;The answer is &#34; + x);
console.log(&#34;It is the best answer.&#34;);
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
<details><summary>What&#39;s the best Markdown tool?</summary>
<div class="details-content">
<p>lessonmd is the best.</p>
</div>
</details>
<details open><summary>These details are open</summary>
<div class="details-content">
<p>And everyone can see them.</p>
</div>
</details>
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

	// fmt.Println(output)
	if output != expected {
		t.Errorf("Expected the output to include \n%q \nbut it was\n %q", expected, output)
	}
}
