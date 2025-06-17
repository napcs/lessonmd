package lessonmd

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration file structure
type Config struct {
	NoWrap               bool   `yaml:"no-wrap"`
	WrapperClass         string `yaml:"wrapper-class"`
	IncludeHighlightJS   bool   `yaml:"include-highlight-js"`
	IncludeMermaidJS     bool   `yaml:"include-mermaid-js"`
	IncludeTabsJS        bool   `yaml:"include-tabs-js"`
	IncludeStylesheet    bool   `yaml:"include-stylesheet"`
	IncludeFrontmatter   bool   `yaml:"include-frontmatter"`
	UseMermaidSVGRenderer bool  `yaml:"use-mermaid-svg-renderer"`
}

// DefaultConfig returns a config with default values
func DefaultConfig() *Config {
	return &Config{
		NoWrap:               false,
		WrapperClass:         "item",
		IncludeHighlightJS:   false,
		IncludeMermaidJS:     false,
		IncludeTabsJS:        false,
		IncludeStylesheet:    false,
		IncludeFrontmatter:   false,
		UseMermaidSVGRenderer: false,
	}
}

// LoadConfig attempts to load configuration from standard locations
func LoadConfig() (*Config, error) {
	config := DefaultConfig()
	
	// Check for config files in order of preference
	configPaths := []string{
		".lessonmd.yaml",
		".lessonmd.yml",
		filepath.Join(os.Getenv("HOME"), ".lessonmd.yaml"),
		filepath.Join(os.Getenv("HOME"), ".lessonmd.yml"),
	}
	
	for _, path := range configPaths {
		if _, err := os.Stat(path); err == nil {
			data, err := os.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
			}
			
			if err := yaml.Unmarshal(data, config); err != nil {
				return nil, fmt.Errorf("failed to parse config file %s: %w", path, err)
			}
			
			break
		}
	}
	
	return config, nil
}