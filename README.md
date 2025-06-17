# Lesson Markdown Processor

Minimal Markdown to HTML converter implemented using [Goldmark](https://github.com/yuin/goldmark) with support for [GitHub-Flavored Markdown](https://github.github.com/gfm/), with additional extensions for developing technical course content.

Supports
* [Mermaid](https://mermaid.js.org/) diagrams.
* Tables.
* Strikethroughs.
* Automatic linking of URLs.
* Raw "pass-through" HTML.
* Client-side highlighting of source code blocks with "copy to clipboard" functionality.
* Notices (admonitions, like "tip", "warning", "note", and others.)
* Tabbed content sections for organizing related information.

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

### Configuration File

You can configure default options using a YAML configuration file. The tool looks for configuration files in the following locations (in order):

1. `.lessonmd.yaml` in the current directory
2. `.lessonmd.yml` in the current directory  
3. `$HOME/.lessonmd.yaml`
4. `$HOME/.lessonmd.yml`

#### Configuration File Format

Create a YAML file with any of the following options:

```yaml
# Basic output options
wrapper-class: "lesson"          # CSS class for outer div (default: "item")
no-wrap: false                   # Don't wrap output in div (default: false)

# Content inclusion options
include-stylesheet: true         # Include CSS in <style> tag (default: false)
include-frontmatter: false       # Include YAML frontmatter as table (default: false)

# JavaScript library options
include-highlight-js: true       # Include Highlight.js from CDN (default: false)
include-mermaid-js: false        # Include Mermaid.js from CDN (default: false)
include-tabs-js: false           # Include tabs JavaScript (default: false)

# Mermaid rendering options
use-mermaid-svg-renderer: false  # Use server-side SVG for Mermaid (default: false)
```

#### Configuration Example

Here's a complete example configuration file (`.lessonmd.yaml`):

```yaml
# Use custom CSS class for styling
wrapper-class: "course-content"

# Always include essential features
include-stylesheet: true
include-highlight-js: true
include-mermaid-js: true

# Don't include frontmatter by default
include-frontmatter: false

# Use client-side Mermaid rendering
use-mermaid-svg-renderer: false
```

**Note:** Command-line flags always override configuration file settings. This allows you to set sensible defaults in your config file while still being able to override them when needed.

If you prefer, you can pipe the output of `cat`:

```bash
cat lesson.md | lessonmd > lesson.html
```

You can also use a Bash herestring:

```bash
lessonmd <<< "Hello world" > lesson.html
```

If you're on macOS, you can send the output directly to the clipboard using the built-in `pbcopy` command:

```bash
lessonmd < lesson.md | pbcopy
```

If you need to convert multiple files, you can do this in Bash:

```bash
for f in *.md; do lessonmd < "${f}" > "${f%.md}.html"; done
```

The HTML output will be wrapped in a `<div>` tag with the class `item`, which will make styling easier. Use the `-no-wrap` flag to generate unwrapped output, or use the `-c` flag to specify a different class name.

YAML frontmatter is skipped by default. To preserve it, add the `-include-frontmatter` flag which renders the front matter as an HTML table at the top of the document.

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
  -include-frontmatter
        Include YAML frontmatter as a table. Defaults to false - frontmatter is omitted.
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

### Details (Expandable sections)


You can now define a region that expands and collapses. These are good for questions and answers:

    [details What's the best Markdown tool?
    lessonmd is the best.
    ]

You can also make them open by default:

    [details open These details are open
    And everyone can see them.
    ]

### Tabs

You can create tabbed content sections to organize related information. Each tab is defined using the `=== "Tab Title"` syntax:

    === "Installation"
    You can install this using npm:
    
    ```bash
    npm install -g lessonmd
    ```
    
    === "Configuration"
    Create a config file:
    
    ```yaml
    theme: default
    ```
    
    === "Usage"
    Run the tool like this:
    
    ```bash
    lessonmd < input.md > output.html
    ```

This creates an interactive tabbed interface where users can click between different sections. The first tab is automatically selected as active.



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

## Notices (Admonitions)

Sometimes you'll want to have notices or callouts in your documents, often called "admonitions."

    :::warning This is a warning
    Use this to warn people of something that may go wrong.
    :::

This generates a block with a title and some contents:

    <div class="notice warning">
      <div class="notice-heading">This is a warning</div>
      <div class="notice-body">
        <p>Use this to warn people of something that may go wrong.</p>
      </div>
    </div>

You can use `tip`, `info`, `note`, `caution`, or `warning`. You can place links or code blocks within this as well.

The included stylesheet has basic styling for these as well.

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
│   ├── notices             <- Parser and HTML renderer for notices
│   └── outputblocks        <- Parser and HTML renderer for output blocks
├── go.mod
└── go.sum
```

## Roadmap

There's a ton to do.

* Math support - MathJax support for Goldmark is spotty and requires client-side work as well.
* Support highlighting within code blocks
* Create code block labels

What's not going to happen:
* custom CSS: You can do this using `cat` to append your own stylesheet.
* Full HTML page generation: Again, use `cat` or another tool to wrap this output with your own template.
* Conversion to other formats: Use Pandoc to convert the HTML.

## Changelog

### 0.0.5 (upcoming)
* Add support for tabbed content sections

### 0.0.4 2023-07-11
* Add support for details (expandable sections)

### 0.0.3 2023-07-10
* Add admonitions: `tip`, `note`, `info`, `caution`, and `warning` and appropriate CSS.
* Add link color to CSS.
* Add table header CSS to make it stand out.
* Add zebra striping to tables.


### 0.0.2 - 2023-01-17
* If the document contains YAML front matter, it's ignored by default.
* The `-include-frontmatter` flag renders front matter as an HTML table.

### 0.0.1 - 2023-01-13

* Initial release
