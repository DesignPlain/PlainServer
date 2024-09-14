package types

type Appengine_FlexibleAppVersionDeploymentFile struct {
	// The identifier for this object. Format specified above.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// SHA1 checksum of the file
	Sha1Sum string `json:"sha1Sum,omitempty" yaml:"sha1Sum,omitempty"`

	// Source URL
	SourceUrl string `json:"sourceUrl,omitempty" yaml:"sourceUrl,omitempty"`
}
