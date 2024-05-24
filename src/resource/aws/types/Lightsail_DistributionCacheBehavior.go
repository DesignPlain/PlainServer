package types

type Lightsail_DistributionCacheBehavior struct {
	// The path to a directory or file to cached, or not cache. Use an asterisk symbol to specify wildcard directories (path/to/assets/\-), and file types (\-.html, \-jpg, \-js). Directories and file paths are case-sensitive.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The cache behavior for the specified path.
	Behavior string `json:"behavior,omitempty" yaml:"behavior,omitempty"`
}
