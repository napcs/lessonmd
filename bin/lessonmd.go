package main

import (
	"flag"
	"fmt"
	"io"
	"lessonmd"
	"os"
)

func banner() {
	fmt.Println("lessonmd v" + lessonmd.AppVersion)
}

func main() {

	version := flag.Bool("v", false, "Prints current app version.")
	nowrap := flag.Bool("no-wrap", false, "Do not wrap output with outer <div> tag.")
	wrapperClass := flag.String("c", "item", "The class name for outer div (defaults to 'item'.")
	help := flag.Bool("h", false, "Show this help message.")
	highlightjs := flag.Bool("include-highlight-js", false, "Include script tags to include Highlight.js client-side libraries from CDN and add copy-to-clipboard functionality.")
	mermaidJS := flag.Bool("include-mermaid-js", false, "Include script tags for client-side Mermaid rendering.")
	styleTag := flag.Bool("include-stylesheet", false, "Include CSS in a <style> tag in the output.")
	mermaidSVG := flag.Bool("use-mermaid-svg-renderer", false, "Use embedded SVG for Mermaid instead of client-side JavaScript.")
	printMermaid := flag.Bool("print-mermaid-js", false, "Print the JavaScript code for Mermaid support.")
	printHighlight := flag.Bool("print-highlight-js", false, "Print the JavaScript code for client-side syntax and clipboard support.")
	printCSS := flag.Bool("print-stylesheet", false, "Print the CSS file to standard output. Provide optional parent class. (defaults to 'item' - use `-c` to change.)")

	flag.Parse()

	if *version {
		banner()
		os.Exit(0)
	}

	if *help {
		helpMessage()
		os.Exit(0)
	}

	if *printHighlight {
		out := lessonmd.Converter.GenerateHighlightJS(*wrapperClass)
		io.WriteString(os.Stdout, out)
		os.Exit(0)
	}

	if *printMermaid {
		out := lessonmd.Converter.GenerateMermaidJS()
		io.WriteString(os.Stdout, out)
		os.Exit(0)
	}

	if *printCSS {
		css := lessonmd.Converter.GenerateCSS(*wrapperClass)
		io.WriteString(os.Stdout, css)
		os.Exit(0)
	}

	// Read Markdown from standard input
	markdown, err := io.ReadAll(os.Stdin)

	if err != nil {
		io.WriteString(os.Stderr, "Unable to read file.")
		os.Exit(1)
	}

	o := lessonmd.ConverterOptions{
		Wrap:             !*nowrap,
		WrapperClass:     *wrapperClass,
		AddStyleTag:      *styleTag,
		AddHighlightJS:   *highlightjs,
		UseSVGforMermaid: *mermaidSVG,
		AddMermaidJS:     *mermaidJS,
	}

	out, err := lessonmd.Converter.Run(markdown, o)

	if err != nil {
		io.WriteString(os.Stderr, "Unable to convert file: "+err.Error()+"\n")
		os.Exit(1)
	}

	io.WriteString(os.Stdout, out)
}

func helpMessage() {
	banner()
	fmt.Println("")
	fmt.Println("Minimal Markdown to HTML converter with support for MathJax, Mermaid, and GitHub-Flavored Markdown, with additional extensions for developing technical course content.")
	flag.Usage()
	fmt.Println("")
	fmt.Println("Accepts Markdown document from STDIN and prints to STDOUT.")
	fmt.Println("Use with other CLI tools to convert files.")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("")
	fmt.Println("\tcat lesson.md | lessonmd > lesson.html")
	fmt.Println("\tlessonmd < lesson.md > lesson.html")
	fmt.Println("\tlessonmd <<< \"Hello world\" > lesson.html")
	fmt.Println("")
	fmt.Println("Use other programs to minify, transform, etc.")

}
