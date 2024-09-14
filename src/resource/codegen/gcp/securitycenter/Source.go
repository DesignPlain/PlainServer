package securitycenter

type Source struct {
	/*
	   The organization whose Cloud Security Command Center the Source
	   lives in.


	   - - -
	*/
	Organization string `json:"organization,omitempty" yaml:"organization,omitempty"`

	// The description of the source (max of 1024 characters).
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The source’s display name. A source’s display name must be unique
	   amongst its siblings, for example, two sources with the same parent
	   can't share the same display name. The display name must start and end
	   with a letter or digit, may contain letters, digits, spaces, hyphens,
	   and underscores, and can be no longer than 32 characters.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
