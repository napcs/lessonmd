# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

LessonMD is a Go-based Markdown to HTML converter built on top of Goldmark, specifically designed for creating technical course content. It extends standard GitHub Flavored Markdown with custom extensions for educational materials.

## Architecture

The project follows a modular extension-based architecture:

- **Core converter** (`converter.go`): Main conversion logic using Goldmark
- **CLI interface** (`bin/lessonmd.go`): Command-line tool with flag parsing
- **Custom extensions** (`extensions/`): Goldmark extensions that add specialized functionality
  - `commandblocks/`: Transforms command code blocks for styling differentiation
  - `outputblocks/`: Handles program output display blocks
  - `inlinehighlight/`: Enables `==text==` highlighting syntax
  - `notices/`: Implements admonitions (tip, warning, note, etc.)
  - `details/`: Provides expandable/collapsible sections
  - `tabs/`: Creates tabbed content sections using `=== "Tab Title"` syntax

Each extension follows the same pattern:
- `extender.go`: Goldmark extension registration
- `ast.go`: AST node definitions
- `parser.go`: Markdown parsing logic
- `renderer.go`: HTML rendering logic
- `transformer.go`: AST transformations (when needed)

## Development Commands

### Building
```bash
# Build for current platform
go build -o lessonmd bin/lessonmd.go

# Build all release targets
make all

# Clean build artifacts
make clean
```

### Testing
```bash
# Run all tests
go test -v

# Run specific test
go test -run TestConvert

# Run integration tests
go test -v integration_test.go
```

### Running the tool locally
```bash
# Basic conversion
go run bin/lessonmd.go < examples/lesson.md > output.html

# With all features enabled
go run bin/lessonmd.go -include-highlight-js -include-mermaid-js -include-stylesheet < examples/lesson.md > output.html
```

## Key Implementation Details

- Uses Goldmark's extension system with parsers, AST transformers, and renderers
- Extensions are registered in `converter.go` and applied to the Goldmark instance
- The CLI tool in `bin/lessonmd.go` handles flag parsing and calls the converter
- Configuration file support (`config.go`): YAML files for setting default options (`.lessonmd.yaml/yml`)
- Version is managed in `converter.go` (`AppVersion` variable) and used by the Makefile for releases
- All custom extensions follow Goldmark's priority system for proper parsing order
- Output is wrapped in configurable div containers with CSS class support

## Testing Approach

Tests are in `converter_test.go` and use table-driven testing patterns. Each test converts markdown input and compares against expected HTML output. Integration tests in `integration_test.go` verify the full CLI workflow including file I/O and flag processing.