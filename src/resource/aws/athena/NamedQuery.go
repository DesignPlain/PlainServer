package athena

type NamedQuery struct {
	// Brief explanation of the query. Maximum length of 1024.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Plain language name for the query. Maximum length of 128.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Text of the query itself. In other words, all query statements. Maximum length of 262144.
	Query string `json:"query,omitempty" yaml:"query,omitempty"`

	// Workgroup to which the query belongs. Defaults to `primary`
	Workgroup string `json:"workgroup,omitempty" yaml:"workgroup,omitempty"`

	// Database to which the query belongs.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`
}
