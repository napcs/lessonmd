# Lesson Markdown Processor

Minimal Markdown to HTML converter implemented using [Goldmark](https://github.com/yuin/goldmark) with support for [GitHub-Flavored Markdown](https://github.github.com/gfm/), with additional extensions for developing technical course content.

Supports
* [Mermaid](https://mermaid.js.org/) diagrams.
* Tables.
* Strikethroughs.
* Automatic linking of URLs.
* Raw "pass-through" HTML.
* Client-side highlighting of source code blocks with "copy to clipboard" functionality.

It also has a few additional Markdown extensions for rendering commands and program output, as well as highlighting text.

This tool can also emit a minimal starter stylesheet. Custom CSS is not supported with this tool. If you need it, append the output of this program after a `<style>` definition containing your CSS.

The tool is designed to be used with other CLI tools or as part of a build chain. 

Generated code blocks are not marked up for syntax highlighting. Use [Highlight.js](https://highlightjs.org/) on the client to syntax highlight code. This tool can generate the appropriate Highlight.js JavaScript code as well as buttons to copy the code snippets.

## Installation

Download the binary for your OS and place it on your `PATH`.

## Usage

This tool accepts Markdown text via Standard Input and prints the converted HTML to Standard Output:

The most basic usage would be like the following example, where `lesson.md` is the input, and `lesson.html` is the output:

```bash
lessonmd < lesson.md > lesson.html
```

If you prefer, you can pipe the output of `cat`:

```bash
cat lesson.md | lessonmd > lesson.html
```

You can also use a Bash herestring:

```bash
lessonmd <<< "Hello world" > lesson.html
```

If you need to convert multiple files, you can do this in Bash:

```bash
for f in *.md; do lessonmd < "${f}" > "${f%.md}.html"; done
```

The HTML output will be wrapped in a `<div>` tag with the class `item`, which will make styling easier. Use the `-no-wrap` flag to generate unwrapped output, or use the `-c` flag to specify a different class name.

Use the `-include-stylesheet` flag to generate a `<style>` block at the top of the document with some basic CSS styling you can build on.

```bash
lessonmd -include-stylesheet < lesson.md > lesson.html
```

You can also emit the basic CSS file using the `-print-stylesheet` flag:

```bash
lessonmd -print-stylesheet > style.css
```

You can add the `-c` flag to specify a different class.

```bash
lessonmd -print-stylesheet -c lesson-item > style.css
```

Use `-h` to see the options:

```
  -c string
        The class name for outer div (defaults to 'item'. (default "item")
  -h    Show this help message.
  -include-highlight-js
        Include script tags to include Highlight.js client-side libraries from CDN and add copy-to-clipboard functionality.
  -include-mermaid-js
        Include script tags for client-side Mermaid rendering.
  -include-stylesheet
        Include CSS in a <style> tag in the output.
  -no-wrap
        Do not wrap output with outer <div> tag.
  -print-highlight-js
        Print the JavaScript code for client-side syntax and clipboard support.
  -print-mermaid-js
        Print the JavaScript code for Mermaid support.
  -print-stylesheet -c
        Print the CSS file to standard output. Provide optional parent class. (defaults to 'item' - use -c to change.)
  -use-mermaid-svg-renderer
        Use embedded SVG for Mermaid instead of client-side JavaScript.
  -v    Prints current app version.
```

To generate a single page with CSS, Highlight.js, and Mermaid support, use the following command:

```bash
lessonmd -include-highlight-js \
         -include-mermaid-js \
         -include-stylesheet \
         -c "lesson_item" \
         < examples/lesson.md > lesson.html
```

## Features

The following features are available:

### GitHub Flavored Markdown

Tables, automatic linking, and strikethroughs are available.

### Highlighting words

Wrap words or phrases with `==` to trigger highlighting. 

Example:

    This is ==fancy==.

### Code blocks

Marking up code fences with the `command` language type will transform them to use the `bash` language, but annotate them so you can style them with CSS differently.

Example:

    ```command
    make clean
    ```

### Output blocks

Marking up code fences with the `output` language will transform them into a `<div>` with the `Output` label and the output. This will let you use CSS to differentiate them from regular code snippets, commands, or file listings.

Example:

    ```output
    This is program output
    ```


### Mermaid diagrams

Add Mermaid diagrams using the `mermaid` language type:

    ```mermaid
    graph TD;
        A-->B;
        A-->C;
        B-->D;
        C-->D;
    ```


By default, you'll only be able to process Mermaid diagrams when the page is rendered and you've added Mermaid's client-side processing. Use the `-include-mermaid-js` flag to append a `<script>` block that loads Mermaid, or add it yourself.


To process Mermaid diagrams server-side, install the Mermaid CLI:


```bash
npm install -g @mermaid-js/mermaid-cli
```

Then use the `-use-mermaid-svg-renderer` flag.

This embeds SVGs into the Markdown, so there's no need for client-side JavaScript.

## Syntax highlighting

This tool generates code blocks compatible with [Highlight.js](https://highlightjs.org/). Add HighlightJS to your system and the code blocks will highlight automatically.

Use the `-include-highlight-js` flag to add a `<script>` block to the bottom of the output that will load Highlight.js from a CDN and add "Copy" buttons so readers can copy code to the clipboard.

Alternatively, run with the `-print-highlight-js` option to emit just the script so you can add it to your LMS or platform.


## For developers

This is built using Goldmark which supports Common Mark. Goldmark is a good fit because you can add extensions to the AST or the rendering functions separately. This means adding extensions will be easier.

Here's a map of what the files are in the project:

```
.
├── Makefile                <- Builds all of the executables
├── README.md               <- this file
├── bin
│   └── lessonmd.go         <- The CLI interface
├── converter.go            <- The main Markdown to HTML converter
├── converter_test.go       <- Test cases
├── examples
│   └── lesson.md           <- An example doc 
├── extensions              <- Custom GoldMark extensions
│   ├── commandblocks       <- Parser and HTML renderer for command blocks
│   ├── inlinehighlight     <- Parser and HTML renderer for inline highlighting
│   └── outputblocks        <- Parser and HTML renderer for output blocks
├── go.mod
└── go.sum
```

## Roadmap

There's a ton to do.

* Admonisions (notes, tips, etc)
* Math support - MathJax support for Goldmark is spotty and requires client-side work as well.
* Support highlighting within code blocks
* Create code block labels

What's not going to happen:
* custom CSS: You can do this using `cat` to append your own stylesheet.
* direct file reading and writing.
* Full HTML page generation: Again, use `cat` or another tool to wrap this output with your own template.
* Conversion to other formats: Use Pandoc to convert the HTML.

