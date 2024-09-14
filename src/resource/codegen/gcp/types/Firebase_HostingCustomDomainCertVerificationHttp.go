package types

type Firebase_HostingCustomDomainCertVerificationHttp struct {
	// A text string to serve at the path.
	Desired string `json:"desired,omitempty" yaml:"desired,omitempty"`

	/*
	   Whether Hosting was able to find the required file contents on the
	   specified path during its last check.
	*/
	Discovered string `json:"discovered,omitempty" yaml:"discovered,omitempty"`

	/*
	   (Output)
	   The last time Hosting systems checked for the file contents.
	*/
	LastCheckTime string `json:"lastCheckTime,omitempty" yaml:"lastCheckTime,omitempty"`

	// The path to the file.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
