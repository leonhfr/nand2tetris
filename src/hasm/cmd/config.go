package cmd

const hasm = "0.0.1"

const (
	versionDefault = false
)

const (
	inputUsage   = "input assembly file path"
	outputUsage  = "output binary file path"
	versionUsage = "displays current hasm version"
)

// A Config represents a hasm configuration.
type Config struct {
	version string // Current hasm version
	input   string // Input assembly file path
	output  string // Output binary file path
}

// NewConfig creates a new Config.
func NewConfig() *Config {
	return &Config{
		version: hasm,
	}
}
