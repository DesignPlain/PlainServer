package types

type Lambda_FunctionImageConfig struct {
	// Parameters that you want to pass in with `entry_point`.
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`

	// Entry point to your application, which is typically the location of the runtime executable.
	EntryPoints []string `json:"entryPoints,omitempty" yaml:"entryPoints,omitempty"`

	// Working directory.
	WorkingDirectory string `json:"workingDirectory,omitempty" yaml:"workingDirectory,omitempty"`
}
