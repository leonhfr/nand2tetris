package cmd

const vm = "0.0.1"

const (
	versionDefault = false
)

const (
	inputUsage   = "input .vm file path"
	versionUsage = "displays current version"
)

// A Config represents a vm configuration.
type Config struct {
	version   string // Current vm version
	input     string // Input path
	output    string // Output .asm file path
	filename  string // File name
	directory bool   // Input is a directory
}

// NewConfig creates a new Config.
func NewConfig() *Config {
	return &Config{
		version: vm,
	}
}
