package types

type Firebaserules_RulesetSourceFile struct {
	// Fingerprint (e.g. github sha) associated with the `File`.
	Fingerprint string `json:"fingerprint,omitempty" yaml:"fingerprint,omitempty"`

	/*
	   File name.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Textual Content.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`
}
